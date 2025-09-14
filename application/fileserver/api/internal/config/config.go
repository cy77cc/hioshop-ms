package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	CacheRedis cache.CacheConf
	Minio      MinioConfig
}

type MinioConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	Bucket         string `yaml:"bucket"`
	AccessKey      string `yaml:"accessKey"`
	AccessPassword string `yaml:"accessPassword"`
}
