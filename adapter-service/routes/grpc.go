package routes

import (
	"adapter-service/proto/proto_models"

	"google.golang.org/grpc"
)

type grpcRoute struct {
	proto_models.UnimplementedAdapterServer
	server *grpc.Server
}

func NewGrpcRoute(server *grpc.Server) *grpcRoute {
	return &grpcRoute{server: server}
}

func (g *grpcRoute) RegisterAdapterHandler(handler proto_models.AdapterServer) {
	proto_models.RegisterAdapterServer(g.server, handler)
}
