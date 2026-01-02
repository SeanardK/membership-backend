package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authMiddleware gin.HandlerFunc) {
	var RouterGroup = router.Group("/")
	{
		setupPortfolioRoutes(RouterGroup, authMiddleware)
	}
}
