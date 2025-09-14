package logic

import (
	"context"

	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件
func (l *UploadFileLogic) UploadFile(in *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UploadFileResponse{}, nil
}
