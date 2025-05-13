package main

import (
	"job-search-manager/config"
	"job-search-manager/internal/server"
)

func main() {

	config.Init()

	server.Start()
}
