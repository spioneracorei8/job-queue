package main

import (
	"user-service/config"
	"user-service/server"
)

func getMainServer() *server.Server {
	return &server.Server{
		APP_PORT:        config.APP_PORT,
		GRPC_PORT:       config.GRPC_PORT,
		PSQL_CONNECTION: config.PSQL_CONNECTION,
	}
}

func main() {
	s := getMainServer()
	s.Start()
}
