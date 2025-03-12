package grpcclient

import (
	"chat_service/internal/config"
	"chat_service/internal/grpc/grpc_api"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClient(cfg *config.Config, log *logrus.Logger) (grpc_api.DBServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(
		cfg.Client.IP+":"+cfg.Client.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	c := grpc_api.NewDBServiceClient(conn)
	return c, conn
}
