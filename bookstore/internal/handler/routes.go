// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"bookstore/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: createBookHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: readBookByNameHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/books"),
	)
}
