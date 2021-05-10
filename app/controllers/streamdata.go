package controllers

import (
	"log"

	"github.com/Shoji-Nakasu/gotrading/app/models"
	"github.com/Shoji-Nakasu/gotrading/bitflyer"
	"github.com/Shoji-Nakasu/gotrading/config"
)

//リアルタイムで取ってきたticker情報をデータベースに書き込む
func StreamIngectionData() {
	var tickerChannel = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	//webserver.goのStartWebServer()がブロックしないように並列処理
	go func() {
		for ticker := range tickerChannel {
			log.Printf("action=StreamIngectionData, %v", ticker)
			//取得したtickerを1s,1m,1hのtableにそれぞれ書き込む
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				//新しく作成（CREATE）の場合
				if isCreated == true && duration == config.Config.TradeDuration {
					//TODO
				}
			}
		}
	}()
}
