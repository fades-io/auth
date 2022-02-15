package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/ShiryaevNikolay/auth/internal/database/postgresql"
	"github.com/ShiryaevNikolay/auth/internal/server"
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
func Run() {
	dbConfig := GetDbConfig()
	storage := GetDB(dbConfig)
	srv.Init(storage)
	srv.Run()
}

// Получение конфигурации для БД
func GetDbConfig() *DbConfig {
	return &DbConfig{
		driver:   os.Getenv("DB_DRIVER"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("POSTGRES_DB"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

// Получение БД
func GetDB(dbConfig *DbConfig) server.Storage {
	if dbConfig.driver == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.name, dbConfig.user, dbConfig.password)
		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("Не могу подсоединиться к базе данных, используя драйвер %s", dbConfig.driver)
			log.Fatal("Ошибка:", err)
		} else {
			fmt.Printf("База данных %s подключена\n", dbConfig.driver)
		}
		return postgresql.New(gormDB)
	}
	return nil
}
