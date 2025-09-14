package logic

import (
	"context"

	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文件元数据
func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileInfoLogic) GetFileInfo(req *types.FileIDReq) (resp *types.FileResp, err error) {
	// todo: add your logic here and delete this line

	return
}
