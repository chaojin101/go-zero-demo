package svc

import (
	"bookstore-cache/internal/config"
	"bookstore-cache/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	BookModel model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		BookModel: model.NewBookModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
