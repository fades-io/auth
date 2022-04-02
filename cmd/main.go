package main

import (
	"github.com/ShiryaevNikolay/auth/internal"
	"github.com/ShiryaevNikolay/auth/internal/res/logs"
	"github.com/ShiryaevNikolay/auth/pkg/logging"
	"github.com/joho/godotenv"
)

func main() {
	logger := logging.GetLogger()

	var err error
	err = godotenv.Load()
	if err != nil {
		logger.Fatalf(res.LogFailedToAccessFileEnv, err)
	} else {
		logger.Infoln(res.LogValuesFromEnvFileReceived)
	}

	internal.Run(logger)
}
