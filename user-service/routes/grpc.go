package routes

import (
	"user-service/proto/proto_models"

	"google.golang.org/grpc"
)

type grpcRoute struct {
	server *grpc.Server
}

func NewGrpcRoute(server *grpc.Server) *grpcRoute {
	return &grpcRoute{
		server: server,
	}
}

func (g *grpcRoute) RegisterUserHandler(handler proto_models.UserServer) {
	proto_models.RegisterUserServer(g.server, handler)
}
