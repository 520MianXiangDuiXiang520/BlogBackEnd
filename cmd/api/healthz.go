package api

import (
	"JuneBlog/internal/logic"
	"github.com/gin-gonic/gin"
)

func HealthZRouter(g *gin.RouterGroup) {
	g.GET("", logic.HealthZHandle)
}
