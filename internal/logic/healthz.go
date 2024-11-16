package logic

import (
	"JuneBlog/patch/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

var healthStatusResp = map[string]string{"status": "health"}

func HealthZHandle(ctx *gin.Context) {
	logger.Debug("handle health ok")
	ctx.JSON(http.StatusOK, healthStatusResp)
}
