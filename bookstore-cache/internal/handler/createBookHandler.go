package handler

import (
	"net/http"

	"bookstore-cache/internal/logic"
	"bookstore-cache/internal/svc"
	"bookstore-cache/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateBookLogic(r.Context(), svcCtx)
		resp, err := l.CreateBook(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
