package api

import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/gocolly/colly"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)

// ハンドラーを定義
// func crawlePayPayFleaMarket(url string) string {
// 	url = formatPayPayFleaMarketUrl(url)
// func CrawlePayPayFleaMarket() string {
func CrawlePayPayFleaMarket() []uint8 {

	url :="https://paypayfleamarket.yahoo.co.jp/search/%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89?oepn=1"

	c := colly.NewCollector()
	items := []entity.Item{}


	c.OnHTML("a.sc-6dae2d2e-0", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		name, _ := e.DOM.Find("img.sc-6dae2d2e-1").Attr("alt")
		price := e.DOM.Find("p.sc-6dae2d2e-3").Text()

		image, _ := e.DOM.Find("img.sc-6dae2d2e-1").Attr("src")
		image = strings.Split(image,"?")[0]
		item := entity.Item {
			Url: url,
			Name: name,
			Price: price,
			Image: image,
		}
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

// func P(t interface{}) {
//     fmt.Println(reflect.TypeOf(t))
// }
	return json
	// return string(json)
}

func formatPayPayFleaMarketUrl(url string) (string, error) {
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
