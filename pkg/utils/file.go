package utils

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(file *multipart.FileHeader, path string, context *gin.Context) (string, error) {
	ext := filepath.Ext(file.Filename)
	imageName := uuid.New().String()
	dst := path + imageName + ext

	if err := context.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}

	return imageName + ext, nil
}

func DeleteFile(filename, path string) error {
	fullPath := path + filename
	err := os.Remove(fullPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
