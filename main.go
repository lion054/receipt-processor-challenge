package main

import (
	"fmt"
	"log"
	"os"

	"receipt-processor-challenge/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func main() {
	fmt.Println("Hello World!")
	fmt.Println()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	server.Initialize()
	server.Router.Run("127.0.0.1:" + os.Getenv("PORT")) // add 127.0.0.1 to prevent Windows Defender Firewall from appearing on every launch
}
