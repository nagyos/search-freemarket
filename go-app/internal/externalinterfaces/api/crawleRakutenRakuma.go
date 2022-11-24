package api

import (
	"fmt"

	// "encoding/json"
	"github.com/gocolly/colly"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)

// ハンドラーを定義
func CrawleRakutenRakuma(url string) []entity.Item {
	c := colly.NewCollector()
	items := []entity.Item{}


	c.OnHTML("div.item", func(e *colly.HTMLElement) {
		s := e.DOM.Find("a.link_search_image")
		url, _ := s.Attr("href")
		name, _ := s.Attr("title")
		price := e.DOM.Find("p.item-box__item-price").Text()

		image, _ := s.Find("meta").Attr("content")
		// image, _ := s.Find("img").Attr("src")
		if url != "" {
			item := entity.Item {
				Url: url,
				Name: name,
				Price: price,
				Image: image,
			}
			// fmt.Println(item)
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
		fmt.Println("Got this error:(RakutenRakuma)", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)

	// json, _ := json.Marshal(items)
	// fmt.Printf("%s",json)

	// return json
	return items
}