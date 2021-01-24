package middleware

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"goframe_learn/app/model"
	"goframe_learn/app/service"
	"goframe_learn/library"
	"goframe_learn/library/response"
	"time"
)

var (
	// The underlying JWT middleware.
	Auth *jwt.GfJWTMiddleware
)

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
		Realm:           "goframe_learn",
		Key:             []byte(key),
		Timeout:         time.Hour * 24 * 7,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
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

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
// @summary 用户登录
// @tags    user
// @accept json
// @produce json
// @Param b body model.UserApiLoginReq true "UserApiLoginReq"
// @router  /api/user/login [POST]
// @success 200 {object} response.JsonResponse
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var (
		apiReq     *model.UserApiLoginReq
		serviceReq *model.UserServiceLoginReq
	)
	if err := r.Parse(&apiReq); err != nil {
		return "", err
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		return "", err
	}

	if user := service.User.GetUserByUsernamePassword(serviceReq); user != nil {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
