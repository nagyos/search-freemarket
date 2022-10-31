package api

import (
	// "fmt"
	"encoding/json"
	"time"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)

// ハンドラーを定義
// func crawlemMercari(url string) error {
	// url = formatmMercariUrl(url)
func CrawlemMercari() string {
	url := "https://jp.mercari.com/search?keyword=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89&status=on_sale"

	items := []entity.Item{}
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--window-size=30,120",
			"--disable-gpu",
			"no-sandbox",
			"disable-dev-shm-usage",
		}),
	)

	driver.Start()
	defer driver.Stop()
	page, _ := driver.NewPage(agouti.Browser("chrome"))
	page.Navigate(url)

	time.Sleep(2* time.Second)

	content, _ := page.HTML()
	render := strings.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(render)
	selection := doc.Find("li.ItemGrid__ItemGridCell-sc-14pfel3-1")
	selection.Each(func(i int, s *goquery.Selection) {
		url := s.Find("a.ItemGrid__StyledThumbnailLink-sc-14pfel3-2").AttrOr("href", "")
		e := s.Find("mer-item-thumbnail")
		name := e.AttrOr("alt","non-name")
		//NOTE: item-name属性では何故か値を取ることが出来ないため，alt属性を用いる
		// name := e.AttrOr("item-name","non-name")
		name = name[:len(name)-18]
		price := e.AttrOr("price","non-price")
		image := e.AttrOr("src", "non-image")
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
	json, _ := json.Marshal(items)
	// fmt.Printf("%s",json)
	// fmt.Println(items)


	return string(json)
}

func formatmMercariUrl(url string) (string, error) {
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
