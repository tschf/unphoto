package local

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/reujab/wallpaper"
)

type LocalSource struct{}

const sourceName string = "local"

// This gets set via the SetPicturePath method. Usually by specifying the cli
// flag `local-path`
var picturePath string

func filterPictures(dirList []fs.DirEntry) (filteredList []fs.DirEntry) {
	for _, dirEntry := range dirList {

		baseName := dirEntry.Name()
		isPicture := strings.HasSuffix(baseName, ".jpg") || strings.HasSuffix(baseName, ".png")
		if !dirEntry.IsDir() && isPicture {
			filteredList = append(filteredList, dirEntry)
		}
	}

	return filteredList
}

func (s LocalSource) GetPhoto(applyWallpaper bool) {
	// log.Fatalf("Photo source %s is not implemented", sourceName)

	folderList, err := os.ReadDir(picturePath)

	if err != nil {
		log.Fatalf("Error reading path %s. Check it exists", picturePath)
	}

	folderPictures := filterPictures(folderList)
	if len(folderPictures) == 0 {
		log.Fatal("No pictures found in the specified path")
	}

	numberOfPictures := len(folderPictures)
	randomIndex := rand.Intn(numberOfPictures)
	randomPicture := folderPictures[randomIndex]
	randomPictureAbsPath := filepath.Join(picturePath, randomPicture.Name())
	fmt.Printf("Random picture from your specified path of %d total pictures is \"%s\"\n", numberOfPictures, randomPictureAbsPath)

	if applyWallpaper {
		wallpaper.SetFromFile(randomPictureAbsPath)
	}

}

func (s LocalSource) PrintSourceInfo() {
	fmt.Println("Photo source " + sourceName)
}

func (s LocalSource) SetPicturePath(path string) {
	if len(path) == 0 {
		log.Fatal("Picture path must be specified")
	}

	picturePath = path
	fmt.Printf("Set local data source folder as %s\n", picturePath)
}
