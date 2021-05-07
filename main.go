package main

import (
	"log"

	"github.com/Shoji-Nakasu/gotrading/config"
	"github.com/Shoji-Nakasu/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
