package controller

import (
	"net/http"

	connection "github.com/SeanardK/web-profile/pkg/config"
	"github.com/SeanardK/web-profile/pkg/model"
	"github.com/gin-gonic/gin"
)

type StampController struct {
}

func NewStampController() *StampController {
	return &StampController{}
}

func (sc *StampController) Create(context *gin.Context) {
	var req model.StampRequest
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	stamp := model.Stamp{
		UserID:     req.UserID,
	}

	db := connection.GetDB()
	if err := db.Create(&stamp).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create stamp", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Create stamp", "data": stamp})
}

func (sc *StampController) GetAll(context *gin.Context) {
	var data []model.Stamp

	db := connection.GetDB()
	if err := db.Find(&data).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get stamps"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Get all stamps", "data": data})
}

func (sc *StampController) GetById(context *gin.Context) {
	id := context.Param("id")

	var stamp model.Stamp
	db := connection.GetDB()
	if err := db.First(&stamp, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Stamp not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Get stamp by ID", "data": stamp})
}

func (sc *StampController) DeleteById(context *gin.Context) {
	id := context.Param("id")

	var stamp model.Stamp
	db := connection.GetDB()
	if err := db.First(&stamp, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Stamp not found"})
		return
	}

	if err := db.Delete(&stamp).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete stamp", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Delete stamp by ID", "data": stamp})
}

func (sc *StampController) UpdateById(context *gin.Context) {
	id := context.Param("id")

	var stamp model.Stamp
	db := connection.GetDB()
	if err := db.First(&stamp, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Stamp not found"})
		return
	}

	var req model.StampRequest
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	stamp.UserID = req.UserID

	if err := db.Save(&stamp).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update stamp", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Update stamp by ID", "data": stamp})
}
