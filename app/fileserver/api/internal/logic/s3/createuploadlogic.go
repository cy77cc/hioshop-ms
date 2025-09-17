package s3

import (
	"context"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新建分段上传任务
func NewCreateUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUploadLogic {
	return &CreateUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUploadLogic) CreateUpload(req *types.CreateUploadReq) (resp *types.CreateUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
