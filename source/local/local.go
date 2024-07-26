package local

import (
	"fmt"
	"log"
)

type LocalSource struct{}

const sourceName string = "local"

func (s LocalSource) GetPhoto(applyWallpaper bool) {
	log.Fatalf("Photo source %s is not implemented", sourceName)
}

func (s LocalSource) PrintSourceInfo() {
	fmt.Println("Photo source " + sourceName)
}
