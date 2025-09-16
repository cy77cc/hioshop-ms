package logic

import (
	"context"

	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFilesLogic {
	return &ListFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 列出用户文件列表
func (l *ListFilesLogic) ListFiles(in *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListFilesResponse{}, nil
}
