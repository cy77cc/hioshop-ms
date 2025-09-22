package s3

import (
	"context"
	"encoding/json"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/pb"

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
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	uploadResp, err := l.svcCtx.FileRpc.Upload(l.ctx, &pb.UploadReq{
		Bucket:      req.Bucket,
		Uid:         uint64(uid),
		Hash:        req.Hash,
		IsLast:      false,
		PartNumber:  0,
		FileSize:    req.Size,
		FileName:    req.Name,
		ContentType: req.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateUploadResp{
		UploadID: uploadResp.UploadId,
		ETag:     uploadResp.Etag,
	}, nil
}
