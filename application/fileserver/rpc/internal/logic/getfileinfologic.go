package logic

import (
	"context"

	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文件信息
func (l *GetFileInfoLogic) GetFileInfo(in *pb.GetFileInfoRequest) (*pb.GetFileInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFileInfoResponse{}, nil
}
