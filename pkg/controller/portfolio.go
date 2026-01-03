package controller

import (
	"net/http"

	connection "github.com/SeanardK/web-profile/pkg/config"
	"github.com/SeanardK/web-profile/pkg/model"
	"github.com/SeanardK/web-profile/pkg/utils"
	"github.com/gin-gonic/gin"
)

type PortfolioController struct {
}

func NewPortfolioController() *PortfolioController {
	return &PortfolioController{}
}

var IMAGE_UPLOAD_PATH = "./public/portfolio/images/"

func (uc *PortfolioController) Create(context *gin.Context) {
	var result model.Portfolio

	file, err := context.FormFile("image")
	if err == nil && file != nil {
		imageName, err := utils.UploadFile(file, IMAGE_UPLOAD_PATH, context)

		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Failed to upload image", "error": err.Error()})
			return
		}
		result.Image = imageName
	}

	result.Title = context.PostForm("title")
	result.Description = context.PostForm("description")
	result.Detail = context.PostForm("detail")
	result.Framework = context.PostForm("framework")
	result.Libraries = context.PostForm("libraries")
	result.Repository = context.PostForm("repository")
	result.URL = context.PostForm("url")

	if result.Title == "" {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Title is required"})
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
		gin.H{"message": "Create portfolio", "data": result})
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

	err := db.First(&result, id).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Portfolio not found"})
		return
	}

	if result.Image != "" {
		if err := utils.DeleteFile(result.Image, IMAGE_UPLOAD_PATH); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete image file", "error": err.Error()})
			return
		}
	}

	err = db.Delete(&result, id).Error
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

	file, err := context.FormFile("image")
	if err == nil && file != nil {
		if result.Image != "" {
			utils.DeleteFile(result.Image, IMAGE_UPLOAD_PATH)
		}

		imageName, err := utils.UploadFile(file, IMAGE_UPLOAD_PATH, context)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Failed to upload image", "error": err.Error()})
			return
		}
		result.Image = imageName
	}

	if title := context.PostForm("title"); title != "" {
		result.Title = title
	}
	if description := context.PostForm("description"); description != "" {
		result.Description = description
	}
	if detail := context.PostForm("detail"); detail != "" {
		result.Detail = detail
	}
	if framework := context.PostForm("framework"); framework != "" {
		result.Framework = framework
	}
	if libraries := context.PostForm("libraries"); libraries != "" {
		result.Libraries = libraries
	}
	if repository := context.PostForm("repository"); repository != "" {
		result.Repository = repository
	}
	if url := context.PostForm("url"); url != "" {
		result.URL = url
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
