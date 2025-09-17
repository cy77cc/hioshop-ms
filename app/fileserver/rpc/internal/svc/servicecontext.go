package svc

import (
	"github.com/cy77cc/hioshop_ms/app/fileserver/model"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	FileModel model.FileInfoModel
	Minio     *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	minioClient, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:      credentials.NewStaticV4(c.Minio.AccessKeyID, c.Minio.SecretAccessKey, ""),
		Secure:     c.Minio.UseSSL,
		MaxRetries: 5,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		FileModel: model.NewFileInfoModel(conn, c.CacheConf),
		Minio:     minioClient,
	}
}
