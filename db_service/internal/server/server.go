package server

import (
	"db-service/internal/config"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/webjaba/messanger/grpc_api"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	grpc_api.DBServiceServer
	DB     *gorm.DB
	Logger *logrus.Logger
}

// TODO: need to refactor
func InitServer(cfg *config.Config, logger *logrus.Logger, db *gorm.DB) *Server {
	server := &Server{DB: db, Logger: logger}

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

// func (s *Server) Register(ctx context.Context, in *grpc_api.AuthRequest) (*grpc_api.AuthResponse, error) {
// 	response := &grpc_api.AuthResponse{}
// 	var user storage.User
// 	s.DB.First(&user, "username = ?", in.Username)
// 	if user.Username != "" {
// 		response.Status = InvalidRequest
// 		s.Logger.Errorf("Invalid request status: %v", response.Status)
// 		return response, errors.Error("Invalid request")
// 	}

// }

// func (s *Server) Authorize(ctx *context.Context, in *grpc_api.AuthRequest) (*grpc_api.AuthResponse, error) {
// }

// func (s *Server) FindMessages(ctx *context.Context, in *grpc_api.FindMessagesRequest) (*grpc_api.FindMessagesResponse, error) {
// }

// func (s *Server) FindUser(ctx *context.Context, in *grpc_api.FindUserRequest) (*grpc_api.FindUserResponse, error) {
// }

// func (s *Server) CreateMessage(ctx *context.Context, in *grpc_api.MessageCreationRequest) (*grpc_api.MessageCreationResponse, error) {
// }

// func (s *Server) CreateMessagesPool(ctx *context.Context, in *grpc_api.MessagePoolCreationRequest) (*grpc_api.MessagePoolCreationResponse, error) {
// }
