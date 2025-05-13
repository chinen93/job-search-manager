package config

import "log"

func ConfigLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
