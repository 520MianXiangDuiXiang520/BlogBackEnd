package api

import (
	"JuneBlog/patch/ginx"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(s *gin.Engine) {
	s.Use(ginx.CorsHandler([]string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
	}))
	ginx.URLPatterns(s, "/healthz", HealthZRouter)
	ginx.URLPatterns(s, "/api/article", ArticleRouter)
	ginx.URLPatterns(s, "/api/auth", AuthRouter)
	ginx.URLPatterns(s, "/api/tag", TagRouter)
}
