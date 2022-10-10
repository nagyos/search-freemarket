package main

import (
	"fmt"
	"reflect"
	// "time"
	"github.com/gocolly/colly"
)

type Item struct {
	Url string `json:"url`
	Name string `json:"name`
	Price string `json:price`
	Image string `json:image`
}

func main() {


	// scrapeYahooAuctions()
	scrapeMercari()

	// P("fadf")
	// P(c)
	// i := 0

	// c.OnHTML("article", func(e *colly.HTMLElement) {

	// 	i++
	// 	book := e.DOM.Find("a > h3").Text()
	// 	author := e.DOM.Find("div > a").Text()
	// 	if author == "" {
	// 		author = e.ChildText(".BookLink_userName__avtjq")
	// 	}

	// 	fmt.Printf("%d 著者: %s / タイトル: %s\n", i, author, book)
	// })
	//ここでタイトルを取得



}

func P(t interface{}) {
    fmt.Println(reflect.TypeOf(t))
}


func scrapeYahooAuctions() {
	c := colly.NewCollector()
	i := 0
	items := []Item{}

	c.OnHTML("li[class=Product]", func(e *colly.HTMLElement) {
		// fmt.Println(e)
		i++
		fmt.Print(i)
		url, _ := e.DOM.Find("a[href]").Attr("href")
		// fmt.Println(url)
		name := e.ChildText("h3")
		// fmt.Println(name)
		price := e.ChildText("dev[class=Product__price]")
		// fmt.Println(price)

		image, _ := e.DOM.Find("img[class=Product__imageData]").Attr("src")
		// fmt.Println(image)
		item := Item {
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

	c.Visit("https://auctions.yahoo.co.jp/search/search?auccat=&tab_ex=commerce&ei=utf-8&aq=-1&oq=&sc_i=&fr=auc_top&p=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89&x=0&y=0")
}

func scrapeMercari() {
	c := colly.NewCollector()
	i := 0
	// items := []Item{}


	c.OnHTML("html", func(e *colly.HTMLElement) {
	// c.OnHTML("li", func(e *colly.HTMLElement) {
		// fmt.Println(e)
		i++
		fmt.Print(i)
		// fmt.Println(e.Attr("href"))
		fmt.Println(e)
		url, _ := e.DOM.Find("a[href]").Attr("href")
		fmt.Println(url)
		// name := e.ChildText("span[class=item-name]")
		// // fmt.Println(name)
		// price := e.ChildText("mer-price")
		// fmt.Println(price)

		// image, _ := e.DOM.Find("img[src]").Attr("src")
		// // fmt.Println(image)
		// item := Item {
		// 	Url: url,
		// 	Name: name,
		// 	Price: price,
		// 	Image: image,
		// }

		// items = append(items, item)
	})

	// c.OnHTML("main[id=main]", func(e *colly.HTMLElement) {
	// 	fmt.Println("kitayo")
	// })

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

	c.Visit("https://jp.mercari.com/search?keyword=%E3%82%AD%E3%83%BC%E3%83%9C%E3%83%BC%E3%83%89&status=on_sale")
}