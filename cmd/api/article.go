package api

import (
	"JuneBlog/internal/logic"
	"JuneBlog/patch/ginx"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(g *gin.RouterGroup) {
	g.GET("/list", ginx.HandlerWithJsonWithPage(logic.ArticleListReq))
	g.GET("/detail/:id", ginx.HandlerWithUrlInt("id", logic.ArticleDetailReq))

	g.POST("/new",
		ginx.Permitted(logic.CheckPermitted),
		ginx.HandlerWithJson(logic.NewArtifactReq))
	g.POST("/update/:id",
		ginx.Permitted(logic.CheckPermitted),
		ginx.HandlerWithUrlInt("id", logic.ArticleUpdateReq))
	g.POST("/delete/:id",
		ginx.Permitted(logic.CheckPermitted),
		ginx.HandlerWithUrlInt("id", logic.ArticleDeleteReq))
}
