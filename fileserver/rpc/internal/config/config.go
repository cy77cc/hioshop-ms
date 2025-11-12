package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	cache.CacheConf
	Mysql struct {
		DataSource string
	}
	MinioConfig MinioConfig
}

type MinioConfig struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
	SSL          bool
}
