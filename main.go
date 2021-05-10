package main

import (
	"github.com/Shoji-Nakasu/gotrading/app/controllers"
	"github.com/Shoji-Nakasu/gotrading/config"
	"github.com/Shoji-Nakasu/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngectionData()
	controllers.StartWebServer()

	// apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	// //新規注文
	// order := &bitflyer.Order{
	// 	ProductCode:     config.Config.ProductCode,
	// 	ChildOrderType:  "MARKET",
	// 	Side:            "BUY",
	// 	Size:            0.001, //どれくらいのbitcoinを買うか(最低取引数量は 0.001 BTC)
	// 	MinuteToExpires: 1,
	// 	TimeInForce:     "GTC", //GTC→キャンセルするまで有効という意味
	// }
	// res, _ := apiClient.SendOrder(order)
	// fmt.Println(res.ChildOrderAcceptanceID)

	//注文した情報のリストを取得
	// i := "ChildOrderAcceptanceIDが入る"
	// params := map[string]string{
	// 	"product_code": config.Config.ProductCode,
	// 	"child_order_acceptance_id": i,
	// }
	// r, _ := apiClient.ListOrder(params)
	// fmt.Println(r)

	//データをリアルタイムで取得
	// tickerChannel := make(chan bitflyer.Ticker)
	// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	// for ticker := range tickerChannel {
	// 	fmt.Println(ticker)
	// 	fmt.Println(ticker.GetMidPrice())
	// 	fmt.Println(ticker.DateTime())
	// 	fmt.Println(ticker.TruncateDateTime(time.Second))
	// 	fmt.Println(ticker.TruncateDateTime(time.Minute))
	// 	fmt.Println(ticker.TruncateDateTime(time.Hour))
	// }
}
