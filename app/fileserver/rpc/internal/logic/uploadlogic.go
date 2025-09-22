package logic

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/cy77cc/hioshop_ms/app/fileserver/model"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/pb"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	FileStatusInit = iota
	FileStatusUploading
	FileStatusUploaded
	FileStatusError
)

// 流式上传分片
func (l *UploadLogic) Upload(in *pb.UploadReq) (*pb.UploadResp, error) {
	objectName := getObjectName(in.Hash)

	// 0 表示是开始请求创建分片上传，在这里面应该判断
	if in.PartNumber == 0 {
		// 判断这个文件对象是否存在
		stat, err := l.svcCtx.MinioClient.StatObject(l.ctx, in.Bucket, objectName, minio.StatObjectOptions{})
		resp := minio.ToErrorResponse(err)
		if err != nil {
			// 有报错，而且错误不是对象不存在
			if resp.Code != minio.NoSuchKey {
				return nil, err
			}
		}

		if resp.Code == minio.NoSuchKey {
			uploadID, err := l.svcCtx.MinioCore.NewMultipartUpload(l.ctx, in.Bucket, objectName, minio.PutObjectOptions{})
			if err != nil {
				return nil, err
			}
			s := l.svcCtx.UUID.String()

			_, err = l.svcCtx.FileModel.Insert(l.ctx, &model.FileInfo{
				Uploader:    in.Uid,
				FileName:    in.FileName,
				Size:        in.FileSize,
				Hash:        in.Hash,
				Bucket:      in.Bucket,
				ObjectName:  objectName,
				FileId:      s,
				UploadTime:  time.Now().Unix(),
				Status:      FileStatusInit,
				ContentType: in.ContentType,
			})
			return &pb.UploadResp{
				UploadId: uploadID,
				Etag:     "",
			}, nil
		}
		// 表示已经存在这个文件对象了
		return &pb.UploadResp{
			UploadId: "",
			Etag:     stat.ETag,
		}, nil
	}

	// 不是第一上传
	reader := bytes.NewReader(in.Data)
	part, err := l.svcCtx.MinioCore.PutObjectPart(l.ctx, in.Bucket, objectName, in.UploadId, int(in.PartNumber), reader, in.FileSize, minio.PutObjectPartOptions{})
	if err != nil {
		return nil, err
	}

	if in.IsLast {
		listObjectParts, err := l.svcCtx.MinioCore.ListObjectParts(l.ctx, in.Bucket, objectName, in.UploadId, 0, int(in.PartNumber))
		if err != nil {
			return nil, err
		}

		completePart := make([]minio.CompletePart, len(listObjectParts.ObjectParts))

		for i, v := range listObjectParts.ObjectParts {
			completePart[i].PartNumber = v.PartNumber
			completePart[i].ETag = v.ETag
			completePart[i].ChecksumCRC32 = v.ChecksumCRC32
			completePart[i].ChecksumSHA256 = v.ChecksumSHA256
		}

		if err != nil {
			return nil, fmt.Errorf("insert file info error: %v", err)
		}

		completeMultipartUpload, err := l.svcCtx.MinioCore.CompleteMultipartUpload(l.ctx, in.Bucket, objectName, in.UploadId, completePart, minio.PutObjectOptions{})
		if err != nil {
			return nil, err
		}

		return &pb.UploadResp{
			UploadId: in.UploadId,
			Etag:     completeMultipartUpload.ETag,
		}, nil
	}

	return &pb.UploadResp{
		UploadId: in.UploadId,
		Etag:     part.ETag,
	}, nil
}

func getObjectName(hash string) string {
	return fmt.Sprintf("%s/%s", string(hash[:5]), hash)
}
