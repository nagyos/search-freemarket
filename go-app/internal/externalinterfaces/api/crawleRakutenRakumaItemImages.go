package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// ハンドラーを定義

func CrawleRakutenRakumaItemImages(url string) []uint8 {

	c := colly.NewCollector()
	var images []string

	c.OnHTML("div.sp-slide", func(e *colly.HTMLElement) {
		image, _ := e.DOM.Find("img.sp-image").Attr("src")
		image = strings.Split(image,"?")[0]
		images = append(images, image)
	})

	c.OnRequest(func(r *colly.Request) {
       fmt.Println("Visiting", r.URL)
   	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:(RakutenRakuma + iamges)", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)

	json, _ := json.Marshal(images)

	return json
}
