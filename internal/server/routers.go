package server

import (
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/middlewares"
)

const (
	userLoginURL = "/user/login"
)

func (server *Server) initRouters() {
	server.Router.HandlerFunc(http.MethodPost, userLoginURL, middlewares.SetHeadersMiddleware(server.Login))
	server.Router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})
}
