package models

import (
	"fmt"
	"time"

	"github.com/Shoji-Nakasu/gotrading/bitflyer"
)

type Candle struct {
	ProductCode string
	Duration    time.Duration
	Time        time.Time
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
}

//コンストラクタ
func NewCandle(productCode string, duration time.Duration, timeDate time.Time, open, close, high float64, low, volume float64) *Candle {
	return &Candle{
		productCode,
		duration,
		timeDate,
		open,
		close,
		high,
		low,
		volume,
	}
}

//Candleが保存されるべきTable名を取ってくるメソッド
func (c *Candle) TableName() string {
	return GetCandleTableName(c.ProductCode, c.Duration)
}

//キャンドルスティックの情報をデーターベースに書き込む（最初の作成）
func (c *Candle) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (time, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?)", c.TableName())
	_, err := DbConnection.Exec(cmd, c.Time.Format(time.RFC3339), c.Open, c.Close, c.High, c.Low, c.Volume)
	if err != nil {
		return err
	}
	return err
}

//リアルタイムで取ってきたキャンドルスティック（同じ時間）の情報をデーターベースに書き込む（アップデート）
func (c *Candle) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET open = ?, close = ?, high = ?, low = ?, volume = ? WHERE time = ?", c.TableName())
	_, err := DbConnection.Exec(cmd, c.Open, c.Close, c.High, c.Low, c.Volume, c.Time.Format(time.RFC3339))
	if err != nil {
		return err
	}
	return err
}

//データベースからCandle情報を取得するファンクション
func GetCandle(productCode string, duration time.Duration, dateTime time.Time) *Candle {
	tableName := GetCandleTableName(productCode, duration)
	cmd := fmt.Sprintf("SELECT time, open, close, high, low, volume FROM %s WHERE time = ?", tableName)
	row := DbConnection.QueryRow(cmd, dateTime.Format(time.RFC3339))
	var candle Candle
	err := row.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume)
	if err != nil {
		return nil
	}
	return NewCandle(productCode, duration, candle.Time, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)
}

//ticker情報をデータベースに書き込む（ticker情報がくるたびに呼び出すファンクション）
func CreateCandleWithDuration(ticker bitflyer.Ticker, productCode string, duration time.Duration)bool {
	currentCandle := GetCandle(productCode, duration, ticker.TruncateDateTime(duration))
	price := ticker.GetMidPrice()
	//Candle情報がなかった場合、データベースに書き込む（CREATE）
	if currentCandle == nil {
		candle := NewCandle(productCode, duration, ticker.TruncateDateTime(duration), price, price, price, price, ticker.Volume)
		candle.Create()
		return true
	}
	//取ってきたhighやlowを更新する
	if currentCandle.High <= price {
		currentCandle.High = price
	} else if currentCandle.Low >= price {
		currentCandle.Low = price
	}
	currentCandle.Volume += ticker.Volume
	currentCandle.Close = price
	currentCandle.Save()
	return false
}