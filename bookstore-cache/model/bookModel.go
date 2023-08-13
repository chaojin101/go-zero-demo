package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BookModel = (*customBookModel)(nil)

type (
	// BookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBookModel.
	BookModel interface {
		bookModel
		FindOneByName(ctx context.Context, name string) (*Book, error)
	}

	customBookModel struct {
		*defaultBookModel
	}
)

// NewBookModel returns a model for the database table.
func NewBookModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BookModel {
	return &customBookModel{
		defaultBookModel: newBookModel(conn, c, opts...),
	}
}

func (m *defaultBookModel) FindOneByName(ctx context.Context, name string) (*Book, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", bookRows, m.table)
	var resp Book
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
