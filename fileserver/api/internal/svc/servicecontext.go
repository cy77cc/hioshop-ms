package svc

import (
	"github.com/cy77cc/hioshop/fileserver/api/internal/config"
	"github.com/cy77cc/hioshop/fileserver/rpc/file"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc file.File
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FileRpc: file.NewFile(zrpc.MustNewClient(c.FileRpcConf)),
	}
}
