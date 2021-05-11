package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Shoji-Nakasu/gotrading/app/models"
	"github.com/Shoji-Nakasu/gotrading/config"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	//データフレームの指定（作成）
	limit := 100
	duration := "1s"
	durationTime := config.Config.Durations[duration]
	df, _ := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)

	err := templates.ExecuteTemplate(w, "google.html", df.Candles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

//jsonで何かあったときにはjson型でエラーを返す
func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

//apicandleのurl指定
var apiValidPath = regexp.MustCompile("/api/candle/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

//データーフレームのCandle Stick情報を返すAPIを作成
//ブラウザで入力された値を取ってきてcandle情報を取得
func apiCandleHandler(w http.ResponseWriter, r *http.Request) {
	//product codeをブラウザからAjaxで送る
	productCode := r.URL.Query().Get("product_code")
	if productCode == "" {
		APIError(w, "No product_code param", http.StatusBadRequest) //指定しないとエラーを返す
		return
	}

	//limit(candleの数)をブラウザからAjaxで送る
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit) //ASCIIからIntに変換
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 1000 //デフォルトのcandleの数は1000に設定
	}

	//limit(candleの数)をブラウザからAjaxで送る
	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1m" //デフォルト（空欄の場合）は１分に設定
	}
	durationTime := config.Config.Durations[duration]

	df, _ := models.GetAllCandle(productCode, durationTime, limit)

	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Handlerを登録
func StartWebServer() error {
	http.HandleFunc("/api/candle/", apiMakeHandler(apiCandleHandler))
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
