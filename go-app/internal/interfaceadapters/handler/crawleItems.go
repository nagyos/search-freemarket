package handler
import (
	// "fmt"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/shoki-tagawa/search-freemarket/internal/externalinterfaces/api"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"

)

// ハンドラーを定義
func CrawleItems(c echo.Context) error {
    // return c.JSON(http.StatusOK, api.crawleYahooAuction())
	//NOTE:設定別にどうするか．並び替え，キーワード，販売状況などをどう受け取る？
	//TODO:urlは正しい名前？
	// yaItems := api.crawleYahooAuction(url)
	// mrItems := api.crawleMercari(url)
	// rfmItems := api.crawleRakutenRakuma(url)
	// pfmItems := api.crawlePayPayFleaMarket(url)
	// yaItems := api.CrawleYahooAuction()
	// mrItems := api.CrawleMercari()
	// rtfItems := api.CrawleRakutenRakuma()
	pfmItems := api.CrawlePayPayFleaMarket()

	// fmt.Println(pfmItems)
    var items []entity.Item
	if err := json.Unmarshal(pfmItems,&items); err != nil {
		panic(err)
	}
	// fmt.Println(items)

	// return c.JSON(http.StatusOK, [yaItems, mrItems, rtfItems, pfmItems])
    return c.JSON(http.StatusOK, items)
    // return c.JSON(http.StatusOK, {"yahooAuction": yaItems, "mercari": mrItems, "RakutenRakuma":rtfItems,"PayPayFleaMarket": pfmItems})
}
