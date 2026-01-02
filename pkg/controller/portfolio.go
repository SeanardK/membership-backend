package controller

import (
	"net/http"

	connection "github.com/SeanardK/web-profile/pkg/config"
	"github.com/SeanardK/web-profile/pkg/model"
	"github.com/gin-gonic/gin"
)

type PortfolioController struct {
}

func NewPortfolioController() *PortfolioController {
	return &PortfolioController{}
}

func (uc *PortfolioController) Create(context *gin.Context) {
	var result model.Portfolio

	err := context.ShouldBindJSON(&result)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid request payload"})
		return
	}

	db := connection.GetDB()
	errCreate := db.Create(&result).Error
	if errCreate != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Failed to create portfolio", "error": errCreate.Error()})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"messsage": "Create portfolio", "data": result})
}

func (uc *PortfolioController) GetAll(context *gin.Context) {
	var data []model.Portfolio

	db := connection.GetDB()

	err := db.Find(&data).Error
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Failed to get portfolio"})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": "Get all portfolio", "data": data})
}

func (uc *PortfolioController) GetById(context *gin.Context) {
	id := context.Param("id")

	var result model.Portfolio
	db := connection.GetDB()
	err := db.First(&result, id).Error
	if err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": "Portfolio not found"})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": "Get portfolio by ID", "data": result})
}

func (uc *PortfolioController) DeleteById(context *gin.Context) {
	id := context.Param("id")

	var result model.Portfolio
	db := connection.GetDB()
	err := db.Delete(&result, id).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete portfolio", "error": err.Error()})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": "Delete portfolio by ID", "data": result})
}

func (uc *PortfolioController) UpdateById(context *gin.Context) {
	id := context.Param("id")
	var result model.Portfolio

	db := connection.GetDB()
	err := db.First(&result, id).Error
	if err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": "Portfolio not found"})
		return
	}

	err = context.ShouldBindJSON(&result)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid request payload"})
		return
	}

	err = db.Save(&result).Error
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Failed to update portfolio", "error": err.Error()})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": "Update portfolio by ID", "data": result})
}
