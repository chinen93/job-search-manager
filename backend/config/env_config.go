package config

import "os"

const (
	ENV_MODE  = "mode"
	TEST_MODE = "test"
	PROD_MODE = "prod"
)

func configEnv() {
	os.Setenv(ENV_MODE, TEST_MODE)
	// os.Setenv(ENV_MODE, PROD_MODE)
}
