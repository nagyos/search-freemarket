package handler

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/shoki-tagawa/search-freemarket/internal/externalinterfaces/api"
)

// ハンドラーを定義
func CrawleItemImages(c echo.Context) error {
	url := c.Param("url")

	var itemImages []uint8
	var images []string
	if strings.Contains(url,"mercari") {
		itemImages = api.CrawleMercariItemImages(url)
	}else if strings.Contains(url,"paypayfleamarket") {
		itemImages = api.CrawlePayPayFleaMarketItemImages(url)
	}else if strings.Contains(url,"fril") {
		itemImages =  api.CrawleRakutenRakumaItemImages(url)
	}else if strings.Contains(url,"auctions.yahoo"){
		itemImages = api.CrawleYahooAuctionItemImages(url)
	}else { return c.JSON(http.StatusOK, images) }

	if err := json.Unmarshal(itemImages,&images); err != nil {
		panic(err)
	}

    return c.JSON(http.StatusOK, images)
}
