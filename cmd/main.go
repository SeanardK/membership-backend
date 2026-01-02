package main

import (
	"github.com/SeanardK/web-profile/pkg/config"
	"github.com/SeanardK/web-profile/pkg/database"
	"github.com/SeanardK/web-profile/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.Info("Application is starting...")

	config.ConnectDB()
	database.AutoMigrate()

	router := gin.New()
	routes.SetupRoutes(router)

	logrus.Info("Server running on port 3001")
	router.Run(":3001")
}
