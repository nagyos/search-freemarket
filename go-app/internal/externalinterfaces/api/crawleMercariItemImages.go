package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
)


func CrawleMercariItemImages(itemUrl string) []uint8 {
	baseImageUrl := "https://static.mercdn.net/item/detail/orig/photos/"
	splitUrl := strings.Split(itemUrl,"/")
	itemId := strings.Split(itemUrl,"/")[len(splitUrl) -1 ]
	imageUrl := fmt.Sprintf("%s%s",baseImageUrl, itemId)

	var images []string
	//Mercariの写真のUrlは.jpg前の数字で商品の写真が変わる
	//二分探索で存在する写真の枚数確認

	ok,ng := 0, 11

	for (math.Abs(float64(ng - ok)) > 1) {
		mid := (ok + ng) / 2
		if existUrl(fmt.Sprintf("%s_%d.jpg",imageUrl, mid)) {
			ok = mid
		}else {ng = mid}
	}

	for i := 1; i < ng; i++ {
		images = append(images, fmt.Sprintf("%s_%d.jpg",imageUrl, i))
	}

	json, _ := json.Marshal(images)

	return json
}

func existUrl(url string) bool {
	res, err := http.Get(url)

	return ( res.StatusCode==200 && err == nil)
}