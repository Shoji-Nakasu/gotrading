package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Shoji-Nakasu/gotrading/config"
	_ "github.com/mattn/go-sqlite3"
)

const tableNameSignalEvents = "signal_events"

var DbConnection *sql.DB

//BTC_JPY_1mなどの名前を返す
func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

//データベーススキーマ作成
func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	//signal_eventsというtableを作成
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY NOT NULL,
			product_code STRING,
			side STRING,
			price FLOAT,
			size FLOAT)`, tableNameSignalEvents)
	DbConnection.Exec(cmd)

	//product_code(BTC_JPY)の１秒、１分、１時間のtableを作成
	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				time DATETIME PRIMARY KEY NOT NULL,
				open FLOAT,
				close FLOAT,
				high FLOAT,
				low FLOAT,
				volume FLOAT)`, tableName)
		DbConnection.Exec(c)
	}
}
