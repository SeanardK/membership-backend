package routes

import (
	"github.com/SeanardK/web-profile/pkg/controller"
	"github.com/gin-gonic/gin"
)

func setupPortfolioRoutes(RouterGroup *gin.RouterGroup) {
	portfolioController := &controller.PortfolioController{}

	portfolioRoutes := RouterGroup.Group("/portfolio")
	{
		portfolioRoutes.POST("/", portfolioController.Create)
		portfolioRoutes.GET("/", portfolioController.GetAll)
		portfolioRoutes.GET("/:id", portfolioController.GetById)
		portfolioRoutes.DELETE("/:id", portfolioController.DeleteById)
		portfolioRoutes.PATCH("/:id", portfolioController.UpdateById)
	}
}
