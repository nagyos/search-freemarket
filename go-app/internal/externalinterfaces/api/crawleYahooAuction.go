package api
import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/gocolly/colly"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)
// ハンドラーを定義
// func crawleYahooAuction(url string) string {
// 	url = formatYahooAuctionUrl(url)
func CrawleYahooAuction() string {

	url := "https://auctions.yahoo.co.jp/search/search?auccat=&tab_ex=commerce&ei=utf-8&aq=-1&oq=&sc_i=&fr=auc_top&p=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89&x=0&y=0"

	c := colly.NewCollector()
	items := []entity.Item{}

	c.OnHTML("li[class=Product]", func(e *colly.HTMLElement) {
		url, _ := e.DOM.Find("a[href]").Attr("href")
		name := e.ChildText("h3")
		price := e.ChildText("dev[class=Product__price]")
		image, _ := e.DOM.Find("img[class=Product__imageData]").Attr("src")
		image = strings.Split(image,"?")[0]

		item := entity.Item {
			Url: url,
			Name: name,
			Price: price,
			Image: image,
		}

		items = append(items, item)
		fmt.Println(item)
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

	json, _ := json.Marshal(items)
	fmt.Printf("%s",json)


	return string(json)
}

func formatYahooAuctionUrl(url string) (string, error) {
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
