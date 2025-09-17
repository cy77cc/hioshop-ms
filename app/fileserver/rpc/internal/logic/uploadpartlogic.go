package logic

import (
	"context"

	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartLogic {
	return &UploadPartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流式上传分片
func (l *UploadPartLogic) UploadPart(stream pb.File_UploadPartServer) error {
	// todo: add your logic here and delete this line

	return nil
}
