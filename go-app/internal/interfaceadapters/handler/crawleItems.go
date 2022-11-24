package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/shoki-tagawa/search-freemarket/internal/entity"
	"github.com/shoki-tagawa/search-freemarket/internal/externalinterfaces/api"
	"golang.org/x/sync/errgroup"
)
const (
	mercari = "mercari"
	paypayfleamarket = "paypayfleamarket"
	rakutenrakuma = "rakutenrakuma"
	yahooauction = "yahooauction"
)

const (
	createdTime = "created_time"
	score = "score"
	price = "price"
	likeCount = "like_count"
)

const (
	onSale = "on_sale"
	soldOut = "sold_out"
)

type searchQueryParam struct {
	keyword string
	priceMin string
	priceMax string
	sort string
	order string
	itemStatus string
}

// ハンドラーを定義
func CrawleItemsWithColly(c echo.Context) error {
	sqp := searchQueryParam{
		keyword: c.QueryParam("keyword"),
		priceMin: c.QueryParam("price-min"),
		priceMax: c.QueryParam("price-max"),
		sort: c.QueryParam("sort"),
		order: c.QueryParam("order"),
		itemStatus: c.QueryParam("status"),
	}

	queryParam := adjustUrlWithQueryParamForColly(sqp)

	// 並列処理を開始
	// ctx := context.Background()
	eg := errgroup.Group{}
	// mutex := sync.Mutex{}
	// sem := semaphore.NewWeighted(4)

	allItems := map[string][]entity.Item{}
	eg.Go(func() error {
		pfmItems := api.CrawlePayPayFleaMarket("https://paypayfleamarket.yahoo.co.jp/search/" + queryParam[paypayfleamarket])
		allItems[paypayfleamarket] = pfmItems
		return nil
	})
	eg.Go(func() error {
		rfmItems := api.CrawleRakutenRakuma("https://fril.jp/s?" + queryParam[rakutenrakuma])
		allItems[rakutenrakuma] = rfmItems
		return nil
	})
	eg.Go(func() error {
		yaItems := api.CrawleYahooAuction("https://auctions.yahoo.co.jp/search/search?" + queryParam[yahooauction])
		allItems[yahooauction] = yaItems
		return nil
	})

	if err := eg.Wait(); err != nil { // 実行が終わるまで待つ
		// エラーハンドリング
		fmt.Println(err)
		return err
	}

    // var items []entity.Item
	// if err := json.Unmarshal(pfmItems,&items); err != nil {
	// 	panic(err)
	// }

    return c.JSON(http.StatusOK, allItems)
}
// ハンドラーを定義
func CrawleItemsWithAgouch(c echo.Context) error {
	sqp := searchQueryParam{
		keyword: c.QueryParam("keyword"),
		priceMin: c.QueryParam("price-min"),
		priceMax: c.QueryParam("price-max"),
		sort: c.QueryParam("sort"),
		order: c.QueryParam("order"),
		itemStatus: c.QueryParam("status"),
	}

	queryParam := adjustUrlWithQueryParamForAgouch(sqp)

	// 並列処理を開始
	// ctx := context.Background()
	eg := errgroup.Group{}
	// mutex := sync.Mutex{}
	// sem := semaphore.NewWeighted(4)

	allItems := map[string][]entity.Item{}
	eg.Go(func() error {
		mrItems := api.CrawleMercari("https://jp.mercari.com/search?" + queryParam[mercari])
		allItems[mercari] = mrItems
		return nil
	})

	if err := eg.Wait(); err != nil { // 実行が終わるまで待つ
		// エラーハンドリング
		fmt.Println(err)
		return err
	}

    // var items []entity.Item
	// if err := json.Unmarshal(pfmItems,&items); err != nil {
	// 	panic(err)
	// }

    return c.JSON(http.StatusOK, allItems)
}

func adjustUrlWithQueryParamForColly (sqp searchQueryParam) map[string]string {
	urls := map[string]string{}

	urls[paypayfleamarket] = adjustPayPayFleaMarketQueryParam(sqp)
	urls[rakutenrakuma] = adjustRakutenRakumaQueryParam(sqp)
	urls[yahooauction] = adjustYahooAuctionQueryParam(sqp)

	return urls
}
func adjustUrlWithQueryParamForAgouch (sqp searchQueryParam) map[string]string {
	urls := map[string]string{}

	urls[mercari] = adjustMercariQueryParam(sqp)

	return urls
}

func adjustMercariQueryParam(sqp searchQueryParam) string {
	var queryParam []string
	queryParam = append(queryParam,fmt.Sprintf("keyword=%s",sqp.keyword))
	queryParam = append(queryParam, fmt.Sprintf("sort=%s",adjustMercariSortParam(sqp.sort)))
	queryParam = append(queryParam, fmt.Sprintf("order=%s",sqp.order))

	if sqp.priceMin != "" {queryParam = append(queryParam, fmt.Sprintf("price_min=%s", sqp.priceMin))}
	if sqp.priceMax != "" {queryParam = append(queryParam, fmt.Sprintf("price_max=%s", sqp.priceMax))}

	queryParam = append(queryParam, fmt.Sprintf("status=%s",adjustMercariItemStatusParam(sqp.itemStatus)))

	return strings.Join(queryParam, "&")
}

func adjustMercariSortParam(sort string) string {
	sortType := ""
	switch sort {
		case createdTime:
			sortType = "created_time"
		case score:
			sortType = "score"
		case price:
			sortType = "price"
		case likeCount:
			sortType = "num_likes"
	}

	return sortType
}

func adjustMercariItemStatusParam(itemStatus string) string {
	status := ""
	switch itemStatus {
		case onSale:
			status = "on_sale"
		case soldOut:
			status = "sold_out"
	}
	return status
}

func adjustPayPayFleaMarketQueryParam(sqp searchQueryParam) string {
	var queryParam []string

	queryParam = append(queryParam, fmt.Sprintf("sort=%s",adjustPayPayFleaMarketSortParam(sqp.sort)))
	queryParam = append(queryParam, fmt.Sprintf("order=%s",sqp.order))

	if sqp.priceMin != "" {queryParam = append(queryParam, fmt.Sprintf("minPrice=%s", sqp.priceMin))}
	if sqp.priceMax != "" {queryParam = append(queryParam, fmt.Sprintf("maxPrice=%s", sqp.priceMax))}
	// open=1:販売中 sold=1:売り切れ
	queryParam = append(queryParam, adjustPayPayFleaMarketItemStatus(sqp.itemStatus))


	return fmt.Sprintf("%s?%s",sqp.keyword,strings.Join(queryParam, "&"))
}

func adjustPayPayFleaMarketSortParam(sort string) string {
	sortType := ""
	switch sort {
		case createdTime:
			sortType = "openTime"
		case score:
			sortType = "ranking"
		case price:
			sortType = "price"
		case likeCount:
			sortType = "likeCounts"
	}

	return sortType
}

func adjustPayPayFleaMarketItemStatus(itemStatus string) string {
	status := ""
	switch itemStatus {
		case onSale:
			status = "open=1"
		case soldOut:
			status = "sold=1"
	}
	return status
}


func adjustRakutenRakumaQueryParam(sqp searchQueryParam) string {
	var queryParam []string
	queryParam = append(queryParam,fmt.Sprintf("query=%s",url.QueryEscape(sqp.keyword)))
	// queryParam = append(queryParam,fmt.Sprintf("query=%s",sqp.keyword))
	queryParam = append(queryParam, fmt.Sprintf("sort=%s",adjustRakutenRakumaSortParam(sqp.sort)))
	queryParam = append(queryParam, fmt.Sprintf("order=%s",sqp.order))

	if sqp.priceMin != "" {queryParam = append(queryParam, fmt.Sprintf("min=%s", sqp.priceMin))}
	if sqp.priceMax != "" {queryParam = append(queryParam, fmt.Sprintf("max=%s", sqp.priceMax))}

	queryParam = append(queryParam, fmt.Sprintf("transaction=%s",adjustRakutenRakumaItemStatus(sqp.itemStatus)))

	return strings.Join(queryParam, "&")
}

func adjustRakutenRakumaSortParam(sort string) string {
	sortType := ""
	switch sort {
		case createdTime:
			sortType = "created_at"
		case score:
			sortType = "relevance"
		case price:
			sortType = "sell_price"
		case likeCount:
			sortType = "like_count"
	}

	return sortType
}

func adjustRakutenRakumaItemStatus(itemStatus string) string {
	status := ""
	switch itemStatus {
		case onSale:
			status = "selling"
		case soldOut:
			status = "soldout"
	}
	return status
}

func adjustYahooAuctionQueryParam(sqp searchQueryParam) string {
	var queryParam []string
	queryParam = append(queryParam,fmt.Sprintf("p=%s",url.QueryEscape(sqp.keyword)))

	queryParam = append(queryParam, fmt.Sprintf("s1=%s",ajustYahooAuctionSortParam(sqp.sort)))
	queryParam = append(queryParam, fmt.Sprintf("o1=%s",ajustYahooAuctionOrderParam(sqp.order)))

	if sqp.priceMin != "" {queryParam = append(queryParam, fmt.Sprintf("aucminprice=%s", sqp.priceMin))}
	if sqp.priceMax != "" {queryParam = append(queryParam, fmt.Sprintf("aucmaxprice=%s", sqp.priceMax))}

	queryParam = append(queryParam, fmt.Sprintf("n=%d",100))

	return strings.Join(queryParam, "&")
}

func ajustYahooAuctionSortParam(sort string) string {
	sortType := ""

	switch sort {
		case createdTime:
			sortType = "new"
		case score:
			sortType = "score2"
		case price:
			sortType = "cbids"
		case likeCount:
			sortType = "popular"
	}
	return sortType
}

func ajustYahooAuctionOrderParam(order string) string {
	orderType := ""

	switch order {
		case "asc":
			orderType = "a"
		case "desc":
			orderType = "d"
	}

	return orderType
}