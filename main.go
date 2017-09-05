package main

import (
	"github.com/noobclear/nuclear-bot/app"
	"github.com/noobclear/nuclear-bot/app/config"
	"github.com/go-resty/resty"
)

func main() {
	c := config.GetConfig()

	resty.R().SetAuthToken(c.BotConfigs[0].LIFXAccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{
			"color": "red",
			"period": 1.5,
			"cycles": 3,
			"persist": false,
			"power_on": true
		}`)).
		Post("https://api.lifx.com/v1/lights/all/effects/breathe")

	bm := app.NewBotManager(c)
	bm.StartAll()
}
