package middlewares

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/zeromicro/go-zero/rest"
)

// alias
type Options = cors.Options

type corsWrapper struct {
	*cors.Cors
	optionPassthrough bool
}

func (c corsWrapper) build() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			c.HandlerFunc(w, r)
			// Process [Option] request
			if !c.optionPassthrough &&
				r.Method == http.MethodOptions &&
				r.Header.Get("Access-Control-Request-Method") != "" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next(w, r)
		}
	}
}

func NewCorsMiddleware(options Options) rest.Middleware {
	return corsWrapper{cors.New(options), options.OptionsPassthrough}.build()
}
