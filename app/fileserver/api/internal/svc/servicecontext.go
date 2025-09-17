package svc

import (
	"github.com/cy77cc/hioshop_ms/app/fileserver/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
