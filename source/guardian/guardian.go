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
	"github.com/tschf/unphoto/httpclient"
)

func GetPhoto() {
	fmt.Println("Download guardian pic")

	imageIndexresp := httpclient.GetHttpResponse("https://www.theguardian.com/news/series/ten-best-photographs-of-the-day")
	picOfDayDoc, err := goquery.NewDocumentFromResponse(imageIndexresp)

	picDivs := picOfDayDoc.Find("div.fc-container__inner a.u-faux-block-link__overlay")
	todaysPicPage, _ := picDivs.First().Attr("href")

	todaysPicPageResp := httpclient.GetHttpResponse(todaysPicPage)
	picOfDayDoc, err = goquery.NewDocumentFromResponse(todaysPicPageResp)

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
	savedFile, err := os.Create(path.Base(imageURL.Path))
	defer savedFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(savedFile, readerWithProgress)
	if err != nil {
		log.Fatal(err)
	}
	progressBar.Finish()
	fmt.Println("Done")
}
