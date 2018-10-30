package main

import (
	"fmt"

	"github.com/Stolarskis/flood-alert-app/src/server/server-handler"
)

func main() {
	fmt.Println("Starting Server")
	server := &server.Server{}
	server.Init()
	server.Run(":3000")
}
