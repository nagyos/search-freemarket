package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shoki-tagawa/search-freemarket/internal/interfaceadapters/handler"
)

func main() {

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

 	// ルートを設定
  	e.GET("/search-colly", handler.CrawleItemsWithColly)
  	e.GET("/search-agouch", handler.CrawleItemsWithAgouch)

	e.GET("/getItemImages/:url", handler.CrawleItemImages)
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":8080"))
}
