package handler

import (
	"net/http"

	"bookstore-cache/internal/logic"
	"bookstore-cache/internal/svc"
	"bookstore-cache/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func readBookByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReadBookByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewReadBookByIdLogic(r.Context(), svcCtx)
		resp, err := l.ReadBookById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
