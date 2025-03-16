package server

import (
	"context"
	"db-service/internal/config"
	"db-service/internal/storage"
	"fmt"
	"net"
	"time"

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
			codes.AlreadyExists,
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

func (s *Server) Authorize(ctx context.Context, in *grpc_api.AuthRequest) (*grpc_api.AuthResponse, error) {
	method := "server.Authorize"
	user := storage.User{}

	s.DB.First(&user, "username = ?", in.Username)

	if user.Username != in.Username {
		err := status.Error(
			codes.InvalidArgument,
			"user with this username does not exists",
		)
		s.Logger.Errorf("Error %v, %v", method, err)
		return nil, err
	}

	if user.Password != in.Password {
		err := status.Error(
			codes.InvalidArgument,
			"Incorrect password",
		)
		s.Logger.Errorf("Error %v, %v", method, err)
		return nil, err
	}

	return &grpc_api.AuthResponse{Id: user.ID}, nil
}

func (s *Server) FindUser(ctx context.Context, in *grpc_api.FindUserRequest) (*grpc_api.FindUserResponse, error) {
	users := []storage.User{}
	s.DB.Where("username LIKE ? LIMIT 10", fmt.Sprintf("%%%v%%", in.Username)).Find(&users)
	usernames := []string{}
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}
	return &grpc_api.FindUserResponse{Usernames: usernames}, nil
}

func (s *Server) CreateMessage(ctx context.Context, in *grpc_api.MessageCreationRequest) (*grpc_api.MessageCreationResponse, error) {
	creatingTime, err := time.Parse("2006-01-02 15:04:05", in.GetDatetime())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid time format")
	}
	msg := storage.Message{
		Text:             in.Text,
		CreatingDateTime: creatingTime,
		FromUser:         in.FromUser,
		ToUser:           in.ToUser,
	}
	s.DB.Create(&msg)
	return &grpc_api.MessageCreationResponse{Id: msg.ID}, nil

}

func (s *Server) CreateMessagesPool(ctx context.Context, in *grpc_api.MessagePoolCreationRequest) (*grpc_api.MessagePoolCreationResponse, error) {
	method := "server.CreateMessagePool"
	messages := make([]storage.Message, len(in.Messages))

	for i, msg := range in.Messages {
		creatingTime, err := time.Parse("2006-01-02 15:04:05", msg.GetDatetime())

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid time format")
		}

		messages[i] = storage.Message{
			Text:             msg.GetText(),
			CreatingDateTime: creatingTime,
			FromUser:         msg.GetFromUser(),
			ToUser:           msg.GetToUser(),
		}
	}

	res := s.DB.Create(&messages)

	if res.Error != nil {
		s.Logger.Errorf("Error in %v: %v", method, res.Error)
	}

	s.Logger.Infof("Messages created: %v", res.RowsAffected)

	ids := make([]uint32, len(messages))

	for i, msg := range messages {
		ids[i] = msg.ID
	}

	return &grpc_api.MessagePoolCreationResponse{Ids: ids}, nil
}

func (s *Server) FindMessages(ctx context.Context, in *grpc_api.FindMessagesRequest) (*grpc_api.FindMessagesResponse, error) {
	messages := []storage.Message{}

	s.DB.Where("from_user = ? AND creating_date_time > ?", in.Id, in.Datetime).Find(&messages)

	msgPointers := make([]*grpc_api.MessageForUser, len(messages))

	for i, msg := range messages {
		msgPointers[i] = &grpc_api.MessageForUser{
			Text:     msg.Text,
			Datetime: msg.CreatingDateTime.String(),
			ToUser:   msg.ToUser,
		}
	}

	return &grpc_api.FindMessagesResponse{Messages: msgPointers}, nil
}
