package main

import (
	// "fmt"
	// "strings"
	// "time"
	// "encoding/json"
	// "time"
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shoki-tagawa/search-freemarket/internal/interfaceadapters/handler"
	// "reflect"
	// "github.com/PuerkitoBio/goquery"
	// "github.com/gocolly/colly"
	// "github.com/sclevine/agouti"
)

func main() {


	// crawleYahooAuction()
	// crawleMercari()
	// crawlePayPayFleaMarket()
	// crawleRakutenRakuma()

	// e := echo.New()
	// e.GET("/healthz", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "ok")
	// })
	// e.Logger.Fatal(e.Start(":8000"))

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

 	// ルートを設定
  	e.GET("/", handler.CrawleItems) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する

	e.GET("/getItemImages", handler.CrawleItemImages)
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":8080"))
}

// func crawleItems(c echo.Context) error {
//     // return c.JSON(http.StatusOK, api.crawleYahooAuction())
// 	// yaItems := crawleYahooAuction()
// 	mrItems := crawleMercari()
// 	var item []Item
// 	if err := json.Unmarshal(mrItems,&item); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(item)
// 	// mrItems := crawleMercari()
// 	// rfmItems := crawleRakutenRakuma()
// 	// pfmItems := crawlePaypayFreeMarket()
// 	// crawleYahooAuction()
// 	// crawleMercari()
// 	// crawlePayPayFleaMarket()
// 	// crawleRakutenRakuma()
//     return c.JSON(http.StatusOK,item)
// }
	// "reflect"

// func P(t interface{}) {
//     fmt.Println(reflect.TypeOf(t))
// }


// func crawleYahooAuction() string{
// 	c := colly.NewCollector()
// 	items := []Item{}

// 	c.OnHTML("li[class=Product]", func(e *colly.HTMLElement) {
// 		url, _ := e.DOM.Find("a[href]").Attr("href")
// 		name := e.ChildText("h3")
// 		price := e.ChildText("dev[class=Product__price]")
// 		image, _ := e.DOM.Find("img[class=Product__imageData]").Attr("src")
// 		image = strings.Split(image,"?")[0]

// 		item := Item {
// 			Url: url,
// 			Name: name,
// 			Price: price,
// 			Image: image,
// 		}

// 		items = append(items, item)
// 		fmt.Println(item)
// 	})

// 	c.OnRequest(func(r *colly.Request) {
//        fmt.Println("Visiting", r.URL)
//    	})

// 	c.OnResponse(func(r *colly.Response) {
// 		fmt.Println("Got a response from", r.Request.URL)
// 	})

// 	c.OnError(func(r *colly.Response, e error) {
// 		fmt.Println("Got this error:", e)
// 	})

// 	c.Visit("https://auctions.yahoo.co.jp/search/search?auccat=&tab_ex=commerce&ei=utf-8&aq=-1&oq=&sc_i=&fr=auc_top&p=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89&x=0&y=0")

// 	json, _ := json.Marshal(items)
// 	fmt.Printf("%s",json)
// 	return json
// }

// func crawleMercari() string{

// 	url := "https://jp.mercari.com/search?keyword=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89&status=on_sale"

// 	items := []Item{}
// 	driver := agouti.ChromeDriver(
// 		agouti.ChromeOptions("args", []string{
// 			"--headless",
// 			"--window-size=30,120",
// 			"--disable-gpu",
// 			"no-sandbox",
// 			"disable-dev-shm-usage",
// 		}),
// 	)

// 	driver.Start()
// 	defer driver.Stop()
// 	page, _ := driver.NewPage(agouti.Browser("chrome"))
// 	page.Navigate(url)

// 	time.Sleep(2* time.Second)

// 	content, _ := page.HTML()
// 	render := strings.NewReader(content)
// 	doc, _ := goquery.NewDocumentFromReader(render)
// 	selection := doc.Find("li.ItemGrid__ItemGridCell-sc-14pfel3-1")
// 	selection.Each(func(i int, s *goquery.Selection) {
// 		url := s.Find("a.ItemGrid__StyledThumbnailLink-sc-14pfel3-2").AttrOr("href", "")
// 		e := s.Find("mer-item-thumbnail")
// 		name := e.AttrOr("alt","non-name")
// 		//NOTE: item-name属性では何故か値を取ることが出来ないため，alt属性を用いる
// 		// name := e.AttrOr("item-name","non-name")
// 		name = name[:len(name)-18]
// 		price := e.AttrOr("price","non-price")
// 		image := e.AttrOr("src", "non-image")
// 		if url != "" {
// 			item := Item {
// 				Url: url,
// 				Name: name,
// 				Price: price,
// 				Image: image,
// 			}
// 			items = append(items, item)
// 		}
// 	})
// 	json, _ := json.Marshal(items)
// 	fmt.Printf("%s",json)
// 	// fmt.Println(items)
// 	return json
// }


// func crawlePayPayFleaMarket() {

// 	url := "https://paypayfleamarket.yahoo.co.jp/search/%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89?oepn=1"

// 	c := colly.NewCollector()
// 	items := []Item{}


// 	c.OnHTML("a.sc-6dae2d2e-0", func(e *colly.HTMLElement) {
// 		url := e.Attr("href")
// 		name, _ := e.DOM.Find("img.sc-6dae2d2e-1").Attr("alt")
// 		price := e.DOM.Find("p.sc-6dae2d2e-3").Text()

// 		image, _ := e.DOM.Find("img.sc-6dae2d2e-1").Attr("src")
// 		image = strings.Split(image,"?")[0]
// 		item := Item {
// 			Url: url,
// 			Name: name,
// 			Price: price,
// 			Image: image,
// 		}
// 		items = append(items, item)
// 	})

// 	c.OnRequest(func(r *colly.Request) {
//        fmt.Println("Visiting", r.URL)
//    	})

// 	c.OnResponse(func(r *colly.Response) {
// 		fmt.Println("Got a response from", r.Request.URL)
// 	})

// 	c.OnError(func(r *colly.Response, e error) {
// 		fmt.Println("Got this error:", e)
// 	})

// 	c.OnScraped(func(r *colly.Response) {
// 		fmt.Println("Finished", r.Request.URL)
// 	})

// 	c.Visit(url)

// 	json, _ := json.Marshal(items)
// 	fmt.Printf("%s",json)

// }

// func crawleRakutenRakuma() {

// 	url := "https://fril.jp/s?query=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89"

// 	c := colly.NewCollector()
// 	items := []Item{}


// 	c.OnHTML("div.item", func(e *colly.HTMLElement) {
// 		s := e.DOM.Find("a.link_search_image")
// 		url, _ := s.Attr("href")
// 		name, _ := s.Attr("title")
// 		price := e.DOM.Find("p.item-box__item-price").Text()

// 		image, _ := s.Find("img").Attr("src")
// 		item := Item {
// 			Url: url,
// 			Name: name,
// 			Price: price,
// 			Image: image,
// 		}
// 		// fmt.Println(item)
// 		items = append(items, item)
// 	})

// 	c.OnRequest(func(r *colly.Request) {
//        fmt.Println("Visiting", r.URL)
//    	})

// 	c.OnResponse(func(r *colly.Response) {
// 		fmt.Println("Got a response from", r.Request.URL)
// 	})

// 	c.OnError(func(r *colly.Response, e error) {
// 		fmt.Println("Got this error:", e)
// 	})

// 	c.OnScraped(func(r *colly.Response) {
// 		fmt.Println("Finished", r.Request.URL)
// 	})

// 	c.Visit(url)

// 	json, _ := json.Marshal(items)
// 	fmt.Printf("%s",json)

// }
