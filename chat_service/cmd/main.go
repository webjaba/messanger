package main

import (
	"chat_service/internal/config"
	"chat_service/internal/logger"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%+v\n", cfg)

	log, file := logger.SetUpLogger(cfg)
	defer logger.Close(file)

	log.Info("info")

	// TODO: init server

	// TODO: handle routes
}
