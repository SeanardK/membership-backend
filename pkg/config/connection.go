package config

import (
	"fmt"
	"time"

	"github.com/SeanardK/web-profile/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func LoadDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("No .env file found, using system environment variables")
	}

	return &DBConfig{
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnv("DB_PORT", "5432"),
		User:     utils.GetEnv("DB_USER", "postgres"),
		Password: utils.GetEnv("DB_PASSWORD", ""),
		DBName:   utils.GetEnv("DB_NAME", "membership_db"),
		SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
	}
}

func (config *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
}

func ConnectDB() {
	config := LoadDBConfig()

	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		logrus.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatal("Failed to get underlying sql.DB:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	logrus.Info("Successfully connected to PostgreSQL with GORM")
}

func GetDB() *gorm.DB {
	return DB
}
