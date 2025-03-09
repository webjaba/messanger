package main

import (
	"db-service/internal/config"
	"db-service/internal/storage"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	db := storage.ConnectDB(cfg)

	storage.Migrate(db)

	fmt.Println(db)
}
