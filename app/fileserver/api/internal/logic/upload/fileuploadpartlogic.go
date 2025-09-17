package upload

import (
	"context"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传分片
func NewFileUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPartLogic {
	return &FileUploadPartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPartLogic) FileUploadPart(req *types.UploadPartReq) (resp *types.UploadPartResp, err error) {
	// todo: add your logic here and delete this line

	return
}
