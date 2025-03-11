package server

import (
	"db-service/internal/config"
	ga "db-service/internal/grpc_api"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	ga.DBServiceServer
}

// TODO: need to refactor
func InitServer(cfg *config.Config, logger *logrus.Logger) *Server {
	server := &Server{}

	lis, err := net.Listen("tcp", ":"+cfg.Server.Port)

	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	ga.RegisterDBServiceServer(s, server)
	logger.Infof("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	return server
}

// func (s *Server) Register(ctx *context.Context, in *ga.AuthRequest) (*ga.AuthResponse, error) {
// }

// func (s *Server) Authorize(ctx *context.Context, in *ga.AuthRequest) (*ga.AuthResponse, error) {
// }

// func (s *Server) FindMessages(ctx *context.Context, in *ga.FindMessagesRequest) (*ga.FindMessagesResponse, error) {
// }

// func (s *Server) FindUser(ctx *context.Context, in *ga.FindUserRequest) (*ga.FindUserResponse, error) {
// }

// func (s *Server) CreateMessage(ctx *context.Context, in *ga.MessageCreationRequest) (*ga.MessageCreationResponse, error) {
// }

// func (s *Server) CreateMessagesPool(ctx *context.Context, in *ga.MessagePoolCreationRequest) (*ga.MessagePoolCreationResponse, error) {
// }
