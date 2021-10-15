package middleware

import (
	"goframe_template/app/api"
	"goframe_template/app/model"
	"goframe_template/library"
	"goframe_template/library/response"
	"time"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
)

var (
	// Auth The underlying JWT middleware.
	Auth *jwt.GfJWTMiddleware
)

func JWTLogin(r *ghttp.Request) {
	Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {

	pathJWT := "./tmp/password/jwt.txt"
	var key string
	if gfile.Exists(pathJWT) {
		key = gfile.GetContents(pathJWT)
	} else {
		g.Log().Line().Info("create jwt key")
		key = library.RandStringRunes(12)
		if err := gfile.PutContents(pathJWT, key); err != nil {
			g.Log().Line().Error(err)
		}
	}

	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "goframe_template",
		Key:             []byte(key),
		Timeout:         time.Hour * 24 * 7,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "user",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   api.Login,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		LogoutResponse:  LogoutResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})
	if err != nil {
		g.Log().Line().Error(err)
	}
	Auth = authMiddleware
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the web token.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	user := data.(*model.User)

	claims["user"] = user

	return claims
}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	r.SetCtxVar("user", claims["user"])
	return claims["user"]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	response.Json(r, response.Unauthorized, message, nil)
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.Json(r, response.Success, "", g.Map{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.Json(r, response.Success, "", g.Map{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
}

// LogoutResponse is used to set token blacklist.
func LogoutResponse(r *ghttp.Request, code int) {
	response.Json(r, response.Success, "", nil)
}
