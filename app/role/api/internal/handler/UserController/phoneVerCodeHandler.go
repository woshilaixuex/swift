package UserController

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xihu/app/role/api/internal/logic/UserController"
	"xihu/app/role/api/internal/svc"
	"xihu/app/role/api/internal/types"
)

func PhoneVerCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhoneVerCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := UserController.NewPhoneVerCodeLogic(r.Context(), svcCtx)
		resp, err := l.PhoneVerCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
