package svc

import (
	"bookstore/internal/config"
	"bookstore/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	BookModel model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		BookModel: model.NewBookModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
