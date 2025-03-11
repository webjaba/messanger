package main

import (
	"db-service/internal/config"
	"db-service/internal/logger"
	"db-service/internal/server"
	"db-service/internal/storage"
)

func main() {
	cfg := config.MustLoad()

	log, file := logger.SetUpLogger(cfg)
	defer logger.Close(file)

	db := storage.ConnectDB(cfg)

	log.Info("DB connection was successful")

	storage.Migrate(db)

	log.Info("Migrations were successful")

	grpcServer := server.InitServer(cfg, log)

	_ = grpcServer

	log.Info("grpc server initialization was succesful")
}
