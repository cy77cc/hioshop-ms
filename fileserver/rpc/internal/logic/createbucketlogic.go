package logic

import (
	"context"

	"github.com/cy77cc/hioshop/fileserver/rpc/internal/svc"
	"github.com/cy77cc/hioshop/fileserver/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBucketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBucketLogic {
	return &CreateBucketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBucketLogic) CreateBucket(in *pb.CreateBucketReq) (*pb.CreateBucketResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateBucketResp{}, nil
}
