package api

import (
	"JuneBlog/internal/logic"
	"JuneBlog/patch/ginx"
	"github.com/gin-gonic/gin"
)

func TagRouter(g *gin.RouterGroup) {
	g.GET("/list", ginx.HandlerWithJson(logic.TagListReq))
	g.POST("/new", ginx.Permitted(logic.CheckPermitted),
		ginx.HandlerWithJson(logic.NewTagReq))
	g.POST("/delete/:id",
		ginx.Permitted(logic.CheckPermitted),
		ginx.HandlerWithUrlInt("id", logic.TagDeleteReq))
}
