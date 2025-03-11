package server

import (
	"db-service/internal/config"
	"db-service/internal/grpc_api"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	grpc_api.DBServiceServer
}

func InitServer(cfg *config.Config, logger *logrus.Logger) *Server {
	server := &Server{}

	lis, err := net.Listen("tcp", ":"+cfg.Server.Port)

	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	grpc_api.RegisterDBServiceServer(s, server)
	logger.Infof("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	return server
}
