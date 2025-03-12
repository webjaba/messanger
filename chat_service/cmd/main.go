package main

import (
	"chat_service/internal/config"
	"fmt"
)

func main() {
	config := config.MustLoad()
	fmt.Printf("%+v\n", config)

	// TODO: init logger

	// TODO: init server

	// TODO: handle routes
}
