package main

import (
	"log"
	"simple-go-notification-system/server"
)

func main() {
	err := server.StartServer()

	if err != nil {
		log.Fatal(err)
	}
}
