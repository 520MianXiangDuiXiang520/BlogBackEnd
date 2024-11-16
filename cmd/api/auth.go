package api

import (
	"JuneBlog/internal/logic"
	"JuneBlog/patch/ginx"
	"github.com/gin-gonic/gin"
)

func AuthRouter(g *gin.RouterGroup) {
	g.POST("/login", ginx.HandlerWithJson(logic.LoginReq))
	g.POST("/logout",
		ginx.Permitted(logic.CheckPermitted),
		ginx.HandlerWithJson(logic.LogoutReq))
}
