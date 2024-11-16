package ginx

import (
	"JuneBlog/patch/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MiddleWare gin.HandlerFunc

var accessHeaders = []string{
	"Accept",
	"Accept-Encoding",
	"Accept-Language",
	"Authorization",
	"Cache-Control",
	"Connection",
	"Content-Length",
	"Content-Type",
	"cookies",
	"Cookies",
	"cookie",
	"Cookies",
	"DNT",
	"Host",
	"If-Modified-Since",
	"Keep-Alive",
	"Origin",
	"openid",
	"opentoken",
	"Pragma",
	"Referer",
	"Sec-Fetch-Dest",
	"Sec-Fetch-Mode",
	"Sec-Fetch-Site",
	"session",
	"User-Agent",
	"sec-ch-ua",
	"sec-ch-ua-mobile",
	"sec-ch-ua-platform",
	"Token",
	"token",
	"X-CSRF-Token",
	"X_Requested_With",
	"X-CustomHeader",
	"X-Requested-With",
}

var assessHeadersStr = strings.Join(accessHeaders, ",")

func CorsHandler(accessList []string) gin.HandlerFunc {
	return func(context *gin.Context) {

		origin := context.GetHeader("Origin")
		method := context.Request.Method
		logger.Error("ssssssss", origin, method)
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		for _, allow := range accessList {
			if allow == origin {
				context.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Allow-Headers", assessHeadersStr)
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Set("content-type", "application/json")
		// 设置返回格式是json
		if method == "OPTIONS" {
			context.Abort()
			context.JSON(http.StatusOK, map[string]string{"code": "ok"})

			//return
		}
		context.Next()
	}
}

type PermittedFunc func(context *gin.Context) bool

// Permitted 鉴权中间件， 注册该中间件后， 如果 license 返回 true 则表示
// 有权限访问该组（个）接口， 否则响应 403
func Permitted(license PermittedFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		if !license(context) {
			context.Abort()
			context.JSON(http.StatusForbidden, gin.H{"msg": "forbidden"})
		}
	}
}

type UserBase interface {
	GetID() int
}

// AuthFunc 检查授权，如果授权通过，返回授权用户，否则第二个参数返回 false
type AuthFunc func(context *gin.Context) (UserBase, bool)

// Auth 授权中间件，注册使用中间件后，如果授权未通过（af() return nil, false）
// 请求会被在此拦截并响应 301，反之， 如果授权通过，会在请求上下文对象 context
// 中添加一个 user 字段，保存授权的用户信息，授权后可以使用 ctx.Get("user")
// 经过类型转换后获取到该 User 对象
func Auth(af AuthFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, ok := af(context)
		if !ok {
			context.Abort()
			context.JSON(http.StatusUnauthorized,
				gin.H{"msg": "unauthorized"})
		}
		context.Set("user", user)
	}
}
