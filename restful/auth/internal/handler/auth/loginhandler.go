package auth

import (
	"net/http"

	"kapibara-apigateway-gozero/restful/auth/internal/logic/auth"
	"kapibara-apigateway-gozero/restful/auth/internal/svc"
	"kapibara-apigateway-gozero/restful/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("[LoginHandler][Param-parse]Error: %v, req[%v]", err, req)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			logx.Errorf("[LoginHandler][Login-logic]Error: %v", err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// Cookie mode:
			http.SetCookie(w, &http.Cookie{
				Name:     "_kapibara_access_token",
				Value:    resp.AccessToken,
				MaxAge:   int(svcCtx.Config.JwtExpired),
				Path:     "/",
				Domain:   svcCtx.Config.CookieServerDomain,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
			})

			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
