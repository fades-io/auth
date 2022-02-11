package server

import (
	"encoding/json"
	"net/http"
)

// Обработка запрос для входа пользователя
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	messageJson, err := json.Marshal(struct{ Message string }{
		Message: "some message",
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(messageJson)
}
