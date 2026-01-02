package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/SeanardK/web-profile/pkg/config"
	"github.com/SeanardK/web-profile/pkg/database"
	"github.com/SeanardK/web-profile/pkg/middleware"
	"github.com/SeanardK/web-profile/pkg/routes"
	"github.com/SeanardK/web-profile/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	baseURL := strings.TrimRight(utils.GetEnv("KEYCLOAK_BASE_URL", ""), "/")
	realm := utils.GetEnv("KEYCLOAK_REALM", "")
	clientID := utils.GetEnv("CLIENT_ID", "")
	port := utils.GetEnv("PORT", "3001")

	if baseURL == "" || realm == "" || clientID == "" {
		print(baseURL, realm, clientID, port)
		logrus.Fatal("Please set KEYCLOAK_BASE_URL, KEYCLOAK_REALM and CLIENT_ID environment variables")
	}

	issuer := fmt.Sprintf("%s/realms/%s", baseURL, realm)

	ctx := context.Background()
	auth, err := middleware.New(ctx, issuer, clientID)
	if err != nil {
		logrus.Fatalf("failed to initialize OIDC middleware: %v", err)
	}

	logrus.Info("Application is starting...")

	config.ConnectDB()
	database.AutoMigrate()

	router := gin.Default()

	routes.SetupRoutes(router, auth.Middleware())

	logrus.Infof("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}

}
