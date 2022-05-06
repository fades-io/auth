package server

import (
	"log"
	"net/http"
	"os"

	"github.com/ShiryaevNikolay/auth/internal/res"
	"github.com/ShiryaevNikolay/auth/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

// Сущность сервера
type Server struct {
	service Service
	Router *httprouter.Router
	logger *logging.Logger
}

// Инициализация сервера
func (server *Server) Init(storage Storage, logger *logging.Logger) {
	server.logger = logger
	server.service = NewService(storage)

	server.Router = httprouter.New()

	logger.Infoln(res.LogRoutersInit)
	server.initRouters()
}

// Запускаем сервер, слушаем порт
func (server *Server) Run() {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	server.logger.Infof(res.LogServerStartedOnHost, host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, server.Router))
}
