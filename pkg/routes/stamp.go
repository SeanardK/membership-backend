package routes

import (
	"github.com/SeanardK/web-profile/pkg/controller"
	"github.com/gin-gonic/gin"
)

func setupStampRoutes(RouterGroup *gin.RouterGroup, auth gin.HandlerFunc) {
	stampController := &controller.StampController{}

	stampRoutes := RouterGroup.Group("/stamp")
	{
		stampRoutes.POST("", auth, stampController.Create)
		stampRoutes.GET("", stampController.GetAll)
		stampRoutes.GET("/:id", stampController.GetById)
		stampRoutes.DELETE("/:id", auth, stampController.DeleteById)
		stampRoutes.PATCH("/:id", auth, stampController.UpdateById)
	}
}
