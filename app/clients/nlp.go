package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	WitBaseUrl = "https://api.wit.ai"
)

type NLPClient interface {
	GetIntent(text string) (*GetIntentResponse, error)
}

type WitClient struct {
	AccessToken string
	HTTPClient  *http.Client
}

type GetIntentResponse struct {
	Text     string    `json:"_text"`
	Entities Entities  `json:"entities"`
}

type Entities struct {
	Intents []Intent `json:"intent"`
}

type Intent struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}

func (wc *WitClient) GetIntent(text string) (*GetIntentResponse, error) {
	logrus.Infof("Retrieving intent for text: [%s]", text)
	if len(text) < 2 {
		msg := fmt.Sprintf("GetIntent text length < 2: [%s]", text)
		err := errors.New(msg)
		logrus.WithError(err).Error(msg)
		return nil, err
	}

	// Create basic request with auth headers
	req, err := http.NewRequest("GET", WitBaseUrl+"/message", nil)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to create GetIntent request: %s", text)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+wc.AccessToken)

	// Build query params
	q := req.URL.Query()
	q.Add("v", "20170901")
	q.Add("q", text)
	req.URL.RawQuery = q.Encode()

	resp, err := wc.HTTPClient.Do(req)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to execute GetIntent request: %s", text)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to read GetIntent response body: %s", text)
		return nil, err
	}
	logrus.Infof("Raw intent body resp: %s", string(body))

	i := GetIntentResponse{}
	err = json.Unmarshal(body, &i)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to parse JSON GetIntent response body: %s", text)
		return nil, err
	}

	logrus.Infof("Returning parsed intent: %+v", i)
	return &i, nil
}

func NewWitClient(accessToken string) NLPClient {
	if accessToken == "" {
		panic(errors.New("Missing access token to initialize Wit client"))
	}

	client := &http.Client{}
	client.Timeout = time.Second * 5

	return &WitClient{
		AccessToken: accessToken,
		HTTPClient:  client,
	}
}
