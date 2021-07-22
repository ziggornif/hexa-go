package main

import (
	"hexa-go/infra"
	"hexa-go/infra/logger"
)

func main() {
	logger := logger.GetLogger()

	server := infra.NewServer(logger)
	server.Run()

	defer server.Shutdown()
}
