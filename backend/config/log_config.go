package config

import "log"

func configLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
