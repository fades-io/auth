package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ShiryaevNikolay/auth/internal"
	"github.com/ShiryaevNikolay/auth/internal/res"
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


	// Канал для сигналов
	sig := make(chan bool)
	// Основной канал
	loop := make(chan error)

	// Начинается мониторинг сигналов
	go signalHandler(sig, logger)

	for quit := false; !quit; {
		go func ()  {
			internal.Run(logger)
			loop <- nil
		}()

		// Блокируем программу при получении сигнала
		select {
		// Прерываем выполнение программы
		case quit = <-sig:
		// Продолжаем выполлнение программы
		case <-loop:
		}
	}
}

// Обработчик сигланов
func signalHandler(q chan bool, logger *logging.Logger) {
	var quit bool

	c := make(chan os.Signal, 1)
	signal.Notify(
		c, 
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	// Для каждого полученного сигнала
	for signal := range c {
		logger.Infof("Сигнал получен!", signal.String())

		switch signal {
		case syscall.SIGINT, syscall.SIGTERM:
			quit = true
		case syscall.SIGHUP:
			quit = false
		}

		if quit {
			quit = false
			// TODO: closeDB(), closeLog()
		}

		// Оповещаем о прекращении работы
		q <- quit
	}
}
