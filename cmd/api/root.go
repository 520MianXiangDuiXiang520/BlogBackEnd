package api

import (
	"JuneBlog/internal/config"
	"JuneBlog/patch/ginx"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(s *gin.Engine) {
	s.Use(ginx.CorsHandler(config.G.CORSList))
	ginx.URLPatterns(s, "/healthz", HealthZRouter)
	ginx.URLPatterns(s, "/api/article", ArticleRouter)
	ginx.URLPatterns(s, "/api/auth", AuthRouter)
	ginx.URLPatterns(s, "/api/tag", TagRouter)
}
