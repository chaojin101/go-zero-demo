package logic

import (
	"context"

	"bookstore-cache/internal/svc"
	"bookstore-cache/internal/types"
	"bookstore-cache/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBookLogic {
	return &CreateBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBookLogic) CreateBook(req *types.CreateReq) (resp *types.CreateResp, err error) {
	_, err = l.svcCtx.BookModel.Insert(l.ctx, &model.Book{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResp{
		Message: "create " + req.Name,
	}, nil
}
