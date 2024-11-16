package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParamDefault(c *gin.Context, key string, def string) string {
	v := c.Param(key)
	if v == "" {
		return def
	}
	return v
}

func ParamDefaultInt(c *gin.Context, key string, def int) int {
	v := c.Param(key)
	if v == "" {
		return def
	}
	vi, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return vi
}

func QueryDefault(c *gin.Context, key string, def string) string {
	return c.DefaultQuery(key, def)
}

func QueryDefaultInt(c *gin.Context, key string, def int) int {
	v := QueryDefault(c, key, "")
	if v == "" {
		return def
	}
	vi, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return vi
}

func WithGinCtx(ctx context.Context, ginCtx *gin.Context) context.Context {
	return context.WithValue(ctx, "ctx", ginCtx)
}

func GinCtx(ctx context.Context) *gin.Context {
	return ctx.Value("ctx").(*gin.Context)
}
