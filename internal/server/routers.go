package server

import (
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/middlewares"
)

const (
	userURL = "/user"
	userIdURL = userURL + "/id"
	userLoginURL = userURL + "/login"
)

func (server *Server) initRouters() {
	server.Router.HandlerFunc(http.MethodGet, userIdURL, middlewares.SetHeadersMiddleware(server.GetUserId))
	server.Router.HandlerFunc(http.MethodPost, userLoginURL, middlewares.SetHeadersMiddleware(server.Login))
}
