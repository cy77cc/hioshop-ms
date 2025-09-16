package svc

import (
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/config"
	"github.com/cy77cc/hioshop_ms/application/fileserver/rpc/file"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc file.File
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FileRpc: file.NewFile(zrpc.MustNewClient(c.FileRpc)),
	}
}
