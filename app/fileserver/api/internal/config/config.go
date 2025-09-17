package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	FileRpc zrpc.RpcClientConf
	Cache   cache.CacheConf
	Mysql   struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
}
