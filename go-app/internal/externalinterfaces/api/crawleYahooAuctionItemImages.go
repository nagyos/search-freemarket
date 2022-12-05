package api

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
)

// ハンドラーを定義

func CrawleYahooAuctionItemImages(url string) []uint8 {

	c := colly.NewCollector()
	var images []string

	c.OnHTML("li.ProductImage__image", func(e *colly.HTMLElement) {
		image, _ := e.DOM.Find("img").Attr("src")
		images = append(images, image)
	})

	c.OnRequest(func(r *colly.Request) {
       fmt.Println("Visiting", r.URL)
   	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:(yahooAuction + images)", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)

	json, _ := json.Marshal(images)

	return json
}
