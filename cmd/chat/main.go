// Package main ...
package main

import (
	_ "alfachatback/docs"
	"alfachatback/internal/chat/transport/httpserver"
	"alfachatback/internal/pkg/config"
)

// @title 	Chat Service API
// @version	1.0
// @description Chat service API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host 	localhost:8443
// @BasePath /
func main() {
	config.MustLoad()

	server := httpserver.NewRouter()

	server.Run()
}
