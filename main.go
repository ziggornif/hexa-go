package main

import (
	"hexa-go/infra"
)

func main() {
	logger := infra.GetLogger()

	server := infra.NewServer(logger)
	server.Run()

	defer server.Shutdown()
}
