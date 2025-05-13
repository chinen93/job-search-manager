package config

import "os"

func ConfigEnv() {
	os.Setenv(ENV_MODE, TEST_MODE)
	// os.Setenv(ENV_MODE, PROD_MODE)
}
