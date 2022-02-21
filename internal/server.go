package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/ShiryaevNikolay/auth/internal/database/postgresql"
	"github.com/ShiryaevNikolay/auth/internal/server"
	"github.com/ShiryaevNikolay/auth/pkg/logging"
	"github.com/tinrab/retry"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var srv = server.Server{}

// Конфиг БД
type DbConfig struct {
	driver   string
	host     string
	port     string
	name     string
	user     string
	password string
}

// Создание и запуск сервера
func Run(logger *logging.Logger) {
	logger.Infoln("Получение конфигурации для БД")
	dbConfig := GetDbConfig()

	logger.Infoln("Получение БД")
	storage := GetDB(dbConfig, logger)

	logger.Infoln("Инициализация сервера")
	srv.Init(storage, logger)

	logger.Infoln("Запуск сервера")
	srv.Run()
}

// Получение конфигурации для БД
func GetDbConfig() *DbConfig {
	return &DbConfig{
		driver:   os.Getenv("DB_DRIVER"),
		host:     os.Getenv("ADD_DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("POSTGRES_DB"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

// Получение БД
func GetDB(dbConfig *DbConfig, logger *logging.Logger) server.Storage {
	var storage server.Storage

	retry.ForeverSleep(5*time.Second, func(attempt int) error {
		logger.Infof("Попытка подключения к БД: %d\n", attempt)
		if dbConfig.driver == "postgres" {
			dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.name, dbConfig.user, dbConfig.password)

			logger.Infoln("Подключение к БД")
			gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				logger.Fatalf("Не получилось присоединиться к БД, используя драйвер: %s; Ошибка: %s\n", dbConfig.driver, err)
				return err
			} else {
				logger.Infof("База данных %s подключена\n", dbConfig.driver)
			}
			storage = postgresql.New(gormDB)
			return nil
		}
		return nil
	})

	return storage
}
