package svc

import (
	"github.com/cy77cc/hioshop_ms/app/fileserver/model"
	"github.com/cy77cc/hioshop_ms/app/fileserver/rpc/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	FileModel   model.FileInfoModel
	MinioCore   *minio.Core
	MinioClient *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioCore, _ := minio.NewCore(c.Minio.Endpoint,
		&minio.Options{
			Creds:      credentials.NewStaticV4(c.Minio.AccessKeyID, c.Minio.SecretAccessKey, ""),
			MaxRetries: 5,
			Secure:     c.Minio.UseSSL,
		},
	)
	minioClient, _ := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:      credentials.NewStaticV4(c.Minio.AccessKeyID, c.Minio.SecretAccessKey, ""),
		Secure:     c.Minio.UseSSL,
		MaxRetries: 5,
	})

	redis.NewFailoverClient(&redis.FailoverOptions{})

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FileModel:   model.NewFileInfoModel(conn, c.CacheRedis),
		MinioClient: minioClient,
		MinioCore:   minioCore,
	}
}
