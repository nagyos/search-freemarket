package api

import (
	"fmt"
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)

// ハンドラーを定義
// func crawleRakutenRakuma(url string) string {
// 	url = formatRakutenRakumaUrl(url)
func CrawleRakutenRakuma() string {

	url := "https://fril.jp/s?query=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89"

	c := colly.NewCollector()
	items := []entity.Item{}


	c.OnHTML("div.item", func(e *colly.HTMLElement) {
		s := e.DOM.Find("a.link_search_image")
		url, _ := s.Attr("href")
		name, _ := s.Attr("title")
		price := e.DOM.Find("p.item-box__item-price").Text()

		image, _ := s.Find("img").Attr("src")
		item := entity.Item {
			Url: url,
			Name: name,
			Price: price,
			Image: image,
		}
		// fmt.Println(item)
		items = append(items, item)
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

	json, _ := json.Marshal(items)
	// fmt.Printf("%s",json)

	return string(json)
}

func formatRakutenRakumaUrl(url string) (string, error) {
	// rep := regexp.MustCompile(`^(050)([0-9]{4})([0-9]{4})$`)
	// result := rep.FindAllStringSubmatch(forwardingNumber, -1)
	// if len(result) == 0 {
	// 	return "", fmt.Errorf("unknown format: %s", forwardingNumber)
	// }

	// if len(result[0]) != 4 {
	// 	return "", fmt.Errorf("unknown format: %s", forwardingNumber)
	// }

	// return fmt.Sprintf("%s-%s-%s", result[0][1], result[0][2], result[0][3]), nil
		return "",nil

}
