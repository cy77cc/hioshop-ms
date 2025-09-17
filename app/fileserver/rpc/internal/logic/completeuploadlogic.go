package logic

import (
	"context"

	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompleteUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteUploadLogic {
	return &CompleteUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompleteUploadLogic) CompleteUpload(in *pb.CompleteUploadReq) (*pb.CompleteUploadResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CompleteUploadResp{}, nil
}
