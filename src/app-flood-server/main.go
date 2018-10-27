package main

import (
	"fmt"

	"github.com/Stolarskis/flood-alert-app/src/app-flood-server/server"
)

func main() {
	fmt.Println("Starting Server")
	server := &server.Server{}
	server.Init()
	server.Run(":3000")
}
