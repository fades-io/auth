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
}
