package main

import (
	"chat_service/internal/config"
	grpcclient "chat_service/internal/grpc/client"
	"chat_service/internal/logger"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%+v\n", cfg)

	log, file := logger.SetUpLogger(cfg)
	defer logger.Close(file)

	log.Info("info")

	grpcClient, conn := grpcclient.InitClient(cfg, log)
	defer conn.Close()

	fmt.Println(grpcClient)

	// TODO: init server

	// TODO: handle routes
}
