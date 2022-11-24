package api

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)

// ハンドラーを定義
func CrawleYahooAuction(url string)[]entity.Item {
	c := colly.NewCollector()
	items := []entity.Item{}

	c.OnHTML("li[class=Product]", func(e *colly.HTMLElement) {
		url, _ := e.DOM.Find("a[href]").Attr("href")
		name := e.ChildText("h3")
		//現在価格を即決価格よりも優先する
		price := e.DOM.Find("div.Product__priceInfo").Find("span.u-textRed").Text()
		if price == "" {
			price = e.DOM.Find("div.Product__priceInfo").Find("span.Product__priceValue").Text()
		}

		image, _ := e.DOM.Find("img[class=Product__imageData]").Attr("src")
		image = strings.Split(image,"?")[0]
		if url != "" {
			item := entity.Item {
				Url: url,
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

	c.Visit(url)

	// json, _ := json.Marshal(items)
	// fmt.Printf("%s",json)


	// return json
	return items
}
