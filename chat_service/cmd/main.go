package main

import (
	"chat_service/internal/config"
	grpcclient "chat_service/internal/grpc/client"
	"chat_service/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log, file := logger.SetUpLogger(cfg)
	defer logger.Close(file)

	log.Debugf("Config: %+v\n", cfg)

	log.Info("logger setup successful")

	grpcClient, conn := grpcclient.InitClient(cfg, log)
	defer conn.Close()

	_ = grpcClient
	log.Infof("grpc client connected to: %v", conn.CanonicalTarget())

	// TODO: init server

	// TODO: handle routes
}
