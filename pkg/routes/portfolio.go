package routes

import (
	"github.com/SeanardK/web-profile/pkg/controller"
	"github.com/gin-gonic/gin"
)

func setupPortfolioRoutes(RouterGroup *gin.RouterGroup, auth gin.HandlerFunc) {
	portfolioController := &controller.PortfolioController{}

	portfolioRoutes := RouterGroup.Group("/portfolio")
	{
		portfolioRoutes.POST("", auth, portfolioController.Create)
		portfolioRoutes.GET("", portfolioController.GetAll)
		portfolioRoutes.GET("/:id", portfolioController.GetById)
		portfolioRoutes.DELETE("/:id", auth, portfolioController.DeleteById)
		portfolioRoutes.PATCH("/:id", auth, portfolioController.UpdateById)
	}
}
