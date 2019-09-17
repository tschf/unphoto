package config

import (
	"os"
	"path"
)

const APP_NAME = "unphoto"
const APP_VERSION = "0.0.1"

// Get the data directory to store the photo of the day. This is so we can
// the latest photo of the day before applying it as the wallpaper.
//
// The specs, https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html,
// say to use $XDG_DATA_HOME and if it is not set, default it to $HOME/.local/share
func GetDataDir() string {
	var dataDirectory string
	defaultDataHome := os.Getenv("XDG_DATA_HOME")

	if len(defaultDataHome) == 0 {
		dataDirectory = path.Join(os.Getenv("HOME"), ".local", "share")
	} else {
		dataDirectory = path.Join(defaultDataHome)
	}

	dataDirectory = path.Join(dataDirectory, APP_NAME)
	return dataDirectory
}
