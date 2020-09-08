package main

import (
	"Sharykhin/go-election/api/http"
	"os"
)

func main() {
	serverPort, mongoUrl := os.Getenv("SERVER_PORT"), os.Getenv("MONGODB_URL")
	if serverPort == "" {
		serverPort = "3000"
	}
	if mongoUrl == "" {
		mongoUrl = "mongodb://localhost:27017/"
	}
	http.ListenAndServe(serverPort, mongoUrl)
}
