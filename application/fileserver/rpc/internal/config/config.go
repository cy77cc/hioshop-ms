package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Minio MinioConfig
	Mysql struct {
		DataSource string
	}
	cache.CacheConf
}

type MinioConfig struct {
	AccessKeyID     string
	Endpoint        string
	SecretAccessKey string
	UseSSL          bool
	Bucket          string
}
