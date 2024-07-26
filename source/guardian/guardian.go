package guardian

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cheggaaa/pb"
	"github.com/reujab/wallpaper"

	"github.com/tschf/unphoto/config"
	"github.com/tschf/unphoto/httpclient"
)

type GuardianSource struct{}

const sourceName string = "guardian"
const serverUrl string = "https://www.theguardian.com"

func (s GuardianSource) GetPhoto(applyWallpaper bool) {

	imageIndexresp := httpclient.GetHttpResponse("https://www.theguardian.com/news/series/ten-best-photographs-of-the-day")
	picOfDayDoc, _ := goquery.NewDocumentFromReader(imageIndexresp.Body)

	picDivs := picOfDayDoc.Find("a[data-link-name*=\"media\"][data-link-name*=\"group-0\"]")
	todaysPicPage, _ := picDivs.First().Attr("href")

	if len(todaysPicPage) == 0 {
		log.Fatal("Couldn't find pic of the day card. Check selector is still correct.")
	}

	todaysPicPageResp := httpclient.GetHttpResponse(fmt.Sprintf("%s%s", serverUrl, todaysPicPage))
	picOfDayDoc, _ = goquery.NewDocumentFromReader(todaysPicPageResp.Body)

	sources := picOfDayDoc.Find("div.content picture source")
	picSrcSet, _ := sources.First().Attr("srcset")

	srcSetArr := strings.Fields(picSrcSet)
	bestResPhoto := srcSetArr[len(srcSetArr)-2]
	fmt.Printf("Photo URL: %s\n", bestResPhoto)

	todaysPhotoResp := httpclient.GetHttpResponse(bestResPhoto)
	defer todaysPhotoResp.Body.Close()

	requestTotalSize := todaysPhotoResp.ContentLength
	progressBar := pb.New64(requestTotalSize)
	progressBar.SetUnits(pb.U_BYTES)

	imageURL, _ := url.Parse(bestResPhoto)

	progressBar.Prefix(fmt.Sprintf("%s:", path.Base(imageURL.Path)))
	progressBar.Start()
	readerWithProgress := progressBar.NewProxyReader(todaysPhotoResp.Body)
	var savedFile *os.File
	var destWallpaperfile string
	var err error
	sourceFilename := path.Base(imageURL.Path)
	if applyWallpaper {
		dataDir := config.GetDataDir()
		_ = os.MkdirAll(dataDir, os.ModePerm)
		destWallpaperfile = path.Join(dataDir, "photo-of-the-day"+path.Ext(sourceFilename))
		savedFile, err = os.Create(destWallpaperfile)
	} else {
		savedFile, err = os.Create(sourceFilename)
	}
	defer savedFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(savedFile, readerWithProgress)
	if err != nil {
		log.Fatal(err)
	}
	progressBar.Finish()

	fileInfo, _ := savedFile.Stat()
	// Double check the saved file has a positive length before applying the update
	if applyWallpaper && fileInfo.Size() > 0 {
		wallpaper.SetFromFile(destWallpaperfile)
	}
}

func (s GuardianSource) PrintSourceInfo() {
	fmt.Println("Photo source " + sourceName)
}
