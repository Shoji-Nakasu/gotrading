package main

import (
	"fmt"

	"github.com/Shoji-Nakasu/gotrading/bitflyer"
	"github.com/Shoji-Nakasu/gotrading/config"
	"github.com/Shoji-Nakasu/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
	
}
