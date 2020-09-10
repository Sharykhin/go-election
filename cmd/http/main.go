package main

import (
	"Sharykhin/go-election/api/http"
	"os"
)

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "3000"
	}

	http.ListenAndServe(serverPort)
}
