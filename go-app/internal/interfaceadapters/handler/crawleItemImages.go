package handler
import (
	// "fmt"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/shoki-tagawa/search-freemarket/internal/externalinterfaces/api"
)

// ハンドラーを定義
func CrawleItemImages(c echo.Context) error {
    // return c.JSON(http.StatusOK, api.crawleYahooAuction())
	//NOTE:設定別にどうするか．並び替え，キーワード，販売状況などをどう受け取る？
	//TODO:urlは正しい名前？
	//TODO: url別で分岐 case

	itemImages := api.CrawlePayPayFleaMarketItemImages()

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
