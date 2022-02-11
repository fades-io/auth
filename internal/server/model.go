package server

import (
	"fmt"
	"log"
	"net/http"

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
	log.Fatal(http.ListenAndServe("localhost:8080", server.Router))
}
