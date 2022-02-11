package main

import (
	"fmt"
	"log"

	"github.com/ShiryaevNikolay/auth/internal"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Не удалось получить доступ к файлу '.env': %v", err)
	} else {
		fmt.Println("Значения из файла '.env' получены.")
	}

	internal.Run()
}
