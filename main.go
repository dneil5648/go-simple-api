package main

import (
	"github.com/dneil5648/go-simple-api/apiserver"
)

func main() {
	server := apiserver.CreateNewServer(":8080")

	server.AddRoute("/health", apiserver.HandleHealthRoute)

	server.Run()
}
