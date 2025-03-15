package server

import (
	"context"
	"db-service/internal/config"
	"db-service/internal/storage"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/webjaba/messanger/grpc_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) Register(ctx context.Context, in *grpc_api.AuthRequest) (*grpc_api.AuthResponse, error) {
	method := "server.Register"

	if in.Username == "" || in.Password == "" {
		err := status.Error(
			codes.InvalidArgument,
			"password and username must be specidied",
		)
		s.Logger.Errorf("Error %v, %v", method, err)
		return nil, err
	}

	for _, str := range [2]string{in.Username, in.Password} {
		for _, r := range str {
			if r > 127 {
				err := status.Error(
					codes.InvalidArgument,
					"password and username must contain only ASCII symbols",
				)
				s.Logger.Errorf("Error %v, %v", method, err)
				return nil, err
			}
		}
	}

	if len(in.Username) > 10 || len(in.Password) > 10 {
		err := status.Error(
			codes.InvalidArgument,
			"len of password or username must be less-equal than 10",
		)

		s.Logger.Errorf("Error %v, %v", method, err)

		return nil, err
	}

	user := storage.User{}

	s.DB.First(&user, "username = ?", in.Username)

	if user.Username != "" {
		err := status.Error(
			codes.InvalidArgument,
			"User with this username already exists",
		)
		s.Logger.Errorf("Error %v, %v", method, err)
		return nil, err
	}

	user = storage.User{Username: in.Username, Password: in.Password}

	s.DB.Create(&user)

	s.Logger.Infof("New user %v registered successfuly", user.Username)
	return &grpc_api.AuthResponse{Id: user.ID}, nil
}

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
