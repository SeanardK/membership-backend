package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	var RouterGroup = router.Group("/")
	{
		setupPortfolioRoutes(RouterGroup)
	}
}
