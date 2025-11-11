package config

import (
	"os"
)

var Version = "0.0"


var MusicDirectory string

func init() {
	dir, found := os.LookupEnv("MUSIC_DIR")
	if !found {
		dir = os.Getenv("HOME") + "/Music"
	}

	dir = dir + "/tuitune"

	MusicDirectory = dir
}
