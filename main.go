package main

import (
	"github.com/noobclear/nuclear-bot/app"
	"github.com/noobclear/nuclear-bot/app/config"
)

func main() {
	c := config.GetConfig()
	bm := app.NewBotManager(c)
	bm.StartAll()
}
