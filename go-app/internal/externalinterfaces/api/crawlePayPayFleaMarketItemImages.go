package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// ハンドラーを定義

func CrawlePayPayFleaMarketItemImages(url string) []uint8 {

	c := colly.NewCollector()
	var images []string

	notMatchClassName := "slick-cloned"
	c.OnHTML("div.slick-slide", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Attr("class"),notMatchClassName) {
			image, _ := e.DOM.Find("img.sc-da08f04a-2").Attr("src")
			images = append(images, image)
		}
	})

	c.OnRequest(func(r *colly.Request) {
       fmt.Println("Visiting", r.URL)
   	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:(paypay + images)", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)

	json, _ := json.Marshal(images)

	return json
}
