package server

import (
	"net/http"
)

// Обработка запрос для входа пользователя
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yep"))
}
