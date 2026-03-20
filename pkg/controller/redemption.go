package controller

import (
	"net/http"
	"time"

	connection "github.com/SeanardK/web-profile/pkg/config"
	"github.com/SeanardK/web-profile/pkg/model"
	"github.com/gin-gonic/gin"
)

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

type RedemptionController struct {
}

func NewRedemptionController() *RedemptionController {
	return &RedemptionController{}
}

func (sc *RedemptionController) Create(context *gin.Context) {
	var req model.RedemptionRequest
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	redemption := model.Redemption{
		UserID:     req.UserID,
	}

	db := connection.GetDB()
	if err := db.Create(&redemption).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create redemption", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Create redemption", "data": redemption})
}

func (sc *RedemptionController) GetAll(context *gin.Context) {
	var data []model.Redemption

	db := connection.GetDB()
	if err := db.Find(&data).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get redemptions"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Get all redemptions", "data": data})
}

func (sc *RedemptionController) GetById(context *gin.Context) {
	id := context.Param("id")

	var redemption model.Redemption
	db := connection.GetDB()
	if err := db.First(&redemption, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Redemption not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Get redemption by ID", "data": redemption})
}

func (sc *RedemptionController) DeleteById(context *gin.Context) {
	id := context.Param("id")

	var redemption model.Redemption
	db := connection.GetDB()
	if err := db.First(&redemption, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Redemption not found"})
		return
	}

	if err := db.Delete(&redemption).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete redemption", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Delete redemption by ID", "data": redemption})
}

func (sc *RedemptionController) UpdateById(context *gin.Context) {
	id := context.Param("id")

	var redemption model.Redemption
	db := connection.GetDB()
	if err := db.First(&redemption, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Redemption not found"})
		return
	}

	var req model.RedemptionRequest
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	redemption.UserID = req.UserID

	if err := db.Save(&redemption).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update redemption", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Update redemption by ID", "data": redemption})
}
