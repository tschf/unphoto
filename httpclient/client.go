package httpclient

import (
	"log"
	"net/http"
)

func GetHttpResponse(url string) *http.Response {

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}
