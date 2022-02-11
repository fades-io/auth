package internal

import "github.com/ShiryaevNikolay/auth/internal/server"

var srv = server.Server{}

func Run() {
	srv.Init()
	srv.Run()
}
