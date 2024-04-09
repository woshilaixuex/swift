package UserController

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xihu/app/role/api/internal/logic/UserController"
	"xihu/app/role/api/internal/svc"
)

func VerCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := UserController.NewVerCodeLogic(r.Context(), svcCtx)
		resp, err := l.VerCode()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
