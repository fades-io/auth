package main

import (
	"github.com/ShiryaevNikolay/auth/internal"
	"github.com/ShiryaevNikolay/auth/pkg/logging"
	"github.com/joho/godotenv"
)

func main() {
	logger := logging.GetLogger()

	var err error
	err = godotenv.Load()
	if err != nil {
		logger.Fatalf("Не удалось получить доступ к файлу '.env': %v", err)
	} else {
		logger.Infoln("Значения из файла '.env' получены.")
	}

	internal.Run(logger)
}
