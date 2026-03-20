package database

import (
	"github.com/SeanardK/web-profile/pkg/config"
	"github.com/SeanardK/web-profile/pkg/model"
	"github.com/sirupsen/logrus"
)

func AutoMigrate() {
	db := config.GetDB()

	err := db.AutoMigrate(
		&model.Stamp{},
	)

	if err != nil {
		logrus.Fatal("Failed to migrate database:", err)
	}

	logrus.Info("Database migration completed successfully")
}
