package api

import (
	// "encoding/json"
	// "encoding/json"
	"fmt"
	"strings"

	// "strings"

	"github.com/gocolly/colly"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)

// ハンドラーを定義
func CrawlePayPayFleaMarket(url string) []entity.Item {
	const baseUrl = "https://paypayfleamarket.yahoo.co.jp"

	c := colly.NewCollector()
	items := []entity.Item{}


	c.OnHTML("a.sc-6dae2d2e-0", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		name, _ := e.DOM.Find("img.sc-6dae2d2e-1").Attr("alt")
		price := e.DOM.Find("p.sc-6dae2d2e-3").Text()

		image, _ := e.DOM.Find("img.sc-6dae2d2e-1").Attr("src")
		image = strings.Split(image,"?")[0]
		if url != "" {
			item := entity.Item {
				Url: baseUrl + url,
				Name: name,
				Price: price,
				Image: image,
			}
			items = append(items, item)
		}
	})

	c.OnRequest(func(r *colly.Request) {
       fmt.Println("Visiting", r.URL)
   	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)

	return items
}