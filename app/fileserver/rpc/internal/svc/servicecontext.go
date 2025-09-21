package svc

import (
	"github.com/cy77cc/hioshop_ms/app/fileserver/model"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	FileModel   model.FileInfoModel
	MinioCore   *minio.Core
	MinioClient *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	core, _ := minio.NewCore(c.MinioConfig.Endpoint, &minio.Options{
		Secure: c.MinioConfig.SSL,
		Creds:  credentials.NewStaticV4(c.MinioConfig.AccessKey, c.MinioConfig.AccessSecret, ""),
	})
	client, _ := minio.New(c.MinioConfig.Endpoint, &minio.Options{
		Secure: c.MinioConfig.SSL,
		Creds:  credentials.NewStaticV4(c.MinioConfig.AccessKey, c.MinioConfig.AccessSecret, ""),
	})
	return &ServiceContext{
		Config:      c,
		FileModel:   model.NewFileInfoModel(conn, c.CacheConf),
		MinioCore:   core,
		MinioClient: client,
	}
}
