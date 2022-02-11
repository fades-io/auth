package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Router *httprouter.Router
}

func (server *Server) Init() {
	server.Router = httprouter.New()
	server.initRouters()
}

// Запускаем сервер, слушаем порт
func (server *Server) Run() {
	fmt.Println("Запуск сервера на хосте")
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(host + ":" + port, server.Router))
}
