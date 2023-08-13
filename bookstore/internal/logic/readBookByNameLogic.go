package logic

import (
	"context"

	"bookstore/internal/svc"
	"bookstore/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadBookByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadBookByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadBookByNameLogic {
	return &ReadBookByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadBookByNameLogic) ReadBookByName(req *types.ReadBookByNameReq) (resp *types.ReadBookByNameResp, err error) {
	book, err := l.svcCtx.BookModel.FindOneByName(l.ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &types.ReadBookByNameResp{
		Id:   book.Id,
		Name: book.Name,
	}, nil
}
