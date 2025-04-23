package main

import (
	"job-search-manager/config"
	"job-search-manager/internal/server"
	"os"
)

func main() {

	os.Setenv(config.ENV_MODE, config.TEST_MODE)
	// os.Setenv(config.ENV_MODE, config.PROD_MODE)

	server.Start()
}
