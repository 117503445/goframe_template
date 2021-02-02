package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"localhost"}
	r.Response.CORS(corsOptions)
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	r.Middleware.Next()
}
