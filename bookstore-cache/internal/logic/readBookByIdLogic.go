package logic

import (
	"context"

	"bookstore-cache/internal/svc"
	"bookstore-cache/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadBookByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadBookByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadBookByIdLogic {
	return &ReadBookByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadBookByIdLogic) ReadBookById(req *types.ReadBookByIdReq) (resp *types.ReadBookByIdResp, err error) {
	book, err := l.svcCtx.BookModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.ReadBookByIdResp{
		Book: types.Book{
			Id:   book.Id,
			Name: book.Name,
		},
	}, nil
}
