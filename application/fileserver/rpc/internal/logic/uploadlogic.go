package logic

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/pb"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

// 客户端流：客户端连续发送 UploadChunk，结束后服务端返回 UploadResponse
func (l *UploadLogic) Upload(stream pb.File_UploadServer) error {
	// 创建临时文件
	tmpDir := "."
	tmpFile, err := os.CreateTemp(tmpDir, "upload-*")
	if err != nil {
		return status.Errorf(codes.Internal, "create temp file: %v", err)
	}
	tmpPath := tmpFile.Name()
	defer func() {
		tmpFile.Close()
		os.Remove(tmpPath)
	}()

	hasher := sha256.New()
	var origFilename string
	var contentType string
	var total int64

	// 接收流
	for {
		chunk, recvErr := stream.Recv()
		if recvErr == io.EOF {
			break
		}
		if recvErr != nil {
			return status.Errorf(codes.Internal, "recv chunk: %v", recvErr)
		}
		if origFilename == "" && chunk.GetFilename() != "" {
			origFilename = chunk.GetFilename()
		}
		if contentType == "" && chunk.GetContentType() != "" {
			contentType = chunk.GetContentType()
		}
		if len(chunk.GetData()) > 0 {
			n, werr := tmpFile.Write(chunk.GetData())
			if werr != nil {
				return status.Errorf(codes.Internal, "write temp file: %v", werr)
			}
			if _, herr := hasher.Write(chunk.GetData()); herr != nil {
				return status.Errorf(codes.Internal, "hash write: %v", herr)
			}
			total += int64(n)
		}
	}

	// 计算 hash
	sum := hasher.Sum(nil)
	hexHash := hex.EncodeToString(sum) // 64 chars for sha256
	group := hexHash[:5]
	objectName := filepath.Join(group, hexHash)

	// 准备上传到 minio
	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		return status.Errorf(codes.Internal, "seek temp file: %v", err)
	}

	// 进行上传
	ctx := context.Background()
	info, err := l.svcCtx.Minio.PutObject(ctx, l.svcCtx.Config.Minio.Bucket, objectName, tmpFile, total, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "minio put object: %v", err)
	}

	// 插入 MySQL 记录
	//id := uuid.New().String()
	//now := time.Now()
	//insertSQL := `INSERT INTO files (id, hash, orig_name, content_type, size, url, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	//if _, err := s.DB.Exec(insertSQL, id, hexHash, origFilename, contentType, total, url, now); err != nil {
	//	// 注意：如果入库失败，可能需要回滚 minio 文件（删除 object）以保持一致性
	//	_ = s.MinioClient.RemoveObject(ctx, s.MinioBucket, objectName, minio.RemoveObjectOptions{})
	//	return status.Errorf(codes.Internal, "insert db: %v", err)
	//}

	// 返回结果
	resp := &pb.UploadResponse{
		Id:       id,
		Url:      url,
		Filename: origFilename,
		Size:     total,
		Hash:     hexHash,
	}
	if _, err := stream.SendAndClose(resp); err != nil {
		return status.Errorf(codes.Internal, "send and close: %v", err)
	}

	_ = info // info 可用于日志
	return nil
}
