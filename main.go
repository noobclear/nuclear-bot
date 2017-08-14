package main

import (
	"github.com/noobclear/nuclear-bot/app"
	"github.com/noobclear/nuclear-bot/app/config"
)

func main() {
	conf := config.GetConfig()
	bm := app.NewBotManager(conf)
	bm.StartAll()
}
