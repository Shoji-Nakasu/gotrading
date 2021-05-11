package controllers

import (
	"fmt"
	"html/template"
	"net/http"

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

//Handlerを登録
func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
