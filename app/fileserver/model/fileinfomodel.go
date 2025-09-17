package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FileInfoModel = (*customFileInfoModel)(nil)

type (
	// FileInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileInfoModel.
	FileInfoModel interface {
		fileInfoModel
	}

	customFileInfoModel struct {
		*defaultFileInfoModel
	}
)

// NewFileInfoModel returns a model for the database table.
func NewFileInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FileInfoModel {
	return &customFileInfoModel{
		defaultFileInfoModel: newFileInfoModel(conn, c, opts...),
	}
}
