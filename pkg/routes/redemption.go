package routes

import (
	"github.com/SeanardK/web-profile/pkg/controller"
	"github.com/gin-gonic/gin"
)

func setupRedemptionRoutes(RouterGroup *gin.RouterGroup, auth gin.HandlerFunc) {
	redemptionController := &controller.RedemptionController{}

	redemptionRoutes := RouterGroup.Group("/redemption")
	{
		redemptionRoutes.POST("", auth, redemptionController.Create)
		redemptionRoutes.GET("", redemptionController.GetAll)
		redemptionRoutes.GET("/:id", redemptionController.GetById)
		redemptionRoutes.DELETE("/:id", auth, redemptionController.DeleteById)
		redemptionRoutes.PATCH("/:id", auth, redemptionController.UpdateById)
	}
}
