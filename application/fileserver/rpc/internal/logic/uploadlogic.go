package logic

import (
	"context"

	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/pb"

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
	// todo: add your logic here and delete this line

	return nil
}
