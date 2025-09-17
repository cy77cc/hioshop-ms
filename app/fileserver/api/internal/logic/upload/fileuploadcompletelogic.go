package upload

import (
	"context"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 完成上传，合并分片
func NewFileUploadCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadCompleteLogic {
	return &FileUploadCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadCompleteLogic) FileUploadComplete(req *types.CompleteUploadReq) (resp *types.CompleteUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
