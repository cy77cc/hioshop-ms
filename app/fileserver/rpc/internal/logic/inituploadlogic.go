package logic

import (
	"context"
	"fmt"

	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/pb"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitUploadLogic {
	return &InitUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitUploadLogic) InitUpload(in *pb.InitUploadReq) (*pb.InitUploadResp, error) {
	objName := fmt.Sprintf("%s/%s", in.Hash[:5], in.Hash)
	uploadID, err := l.svcCtx.MinioCore.NewMultipartUpload(l.ctx, l.svcCtx.Config.Minio.Bucket, objName, minio.PutObjectOptions{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.InitUploadResp{
		Bucket:   l.svcCtx.Config.Minio.Bucket,
		UploadId: uploadID,
	}, nil
}
