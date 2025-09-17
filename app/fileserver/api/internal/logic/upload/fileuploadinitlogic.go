package upload

import (
	"context"
	"encoding/json"

	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/types"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 初始化上传，生成 uploadId
func NewFileUploadInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadInitLogic {
	return &FileUploadInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadInitLogic) FileUploadInit(req *types.InitUploadReq) (resp *types.InitUploadResp, err error) {
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()
	uploadResp, err := l.svcCtx.FileRpc.InitUpload(l.ctx, &pb.InitUploadReq{Hash: req.Hash, FileName: req.FileName, Uploader: userId})
	if err != nil {
		return nil, err
	}

	return &types.InitUploadResp{
		Bucket:   uploadResp.Bucket,
		UploadId: uploadResp.UploadId,
	}, nil
}
