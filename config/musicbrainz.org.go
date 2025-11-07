package config

import (
	mb "go.uploadedlobster.com/musicbrainzws2"
)


var DefaultAppInfo = mb.AppInfo{
		Name: "tuitune",
		Version: Version,
		URL: "https://github.com/ajoyfuljester/tuitune",
}
