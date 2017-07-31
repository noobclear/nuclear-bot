package main

import "github.com/noobclear/nuclear-bot/app"

func main() {
	conf := app.GetConfig()
	bm := app.NewBotManager(conf)
	bm.StartAll()
}
