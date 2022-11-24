package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
)

// ハンドラーを定義
func CrawleMercari(url string) []entity.Item {
	const baseUrl = "https://jp.mercari.com/"
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

	time.Sleep(4* time.Second)

	content, _ := page.HTML()
	render := strings.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(render)
	selection := doc.Find("li.ItemGrid__ItemGridCell-sc-14pfel3-1")
	selection.Each(func(i int, s *goquery.Selection) {
		url := s.Find("a.ItemGrid__StyledThumbnailLink-sc-14pfel3-2").AttrOr("href", "")
		e := s.Find("mer-item-thumbnail")
		name := e.AttrOr("alt","non-name")
		//NOTE: item-name属性では値を取ることが出来ないため，alt属性を用いる
		// name := e.AttrOr("item-name","non-name")
		name = strings.Split(name,"のサムネイル")[0]
		price := fmt.Sprintf("¥%s",e.AttrOr("price","non-price"))
		image := e.AttrOr("src", "non-image")
		fmt.Println(url)
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
	// json, _ := json.Marshal(items)

	// return json
	return items
}

