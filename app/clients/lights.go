package clients

import (
	"github.com/go-resty/resty"
	"github.com/sirupsen/logrus"
)

type LightsClient interface {
	TriggerPulse() error
}

type LIFXClient struct {
	AccessToken string
}

func (c *LIFXClient) TriggerPulse() error {
	resp, err := resty.R().SetAuthToken(c.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{
			"color": "red",
			"period": 1.5,
			"cycles": 3,
			"persist": false,
			"power_on": true
		}`)).
		Post("https://api.lifx.com/v1/lights/all/effects/breathe")
	if err != nil {
		logrus.WithError(err).Errorf("Failed to trigger lights: %+v", resp)
	}
	return err
}

func NewLIFXClient(accessToken string) LightsClient {
	return &LIFXClient{accessToken}
}