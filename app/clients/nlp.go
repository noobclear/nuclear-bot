package clients

import (
	"net/http"
	"time"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
)

const (
	WitBaseUrl = "https://api.wit.ai"
)

type NLPClient interface {
	GetIntent(text string) (*Intent, error)
}

type WitClient struct {
	AccessToken string
	HTTPClient  *http.Client
}

type Intent struct {
	MsgId string `json:"msg_id"`
	Text  string `json:"_text"`
	Outcomes []Outcome `json:"outcomes"`
}

type Outcome struct {
	Intent     string `json:"intent"`
	Confidence float32 `json:"confidence"`
}

func (wc *WitClient) GetIntent(text string) (*Intent, error) {
	logrus.Infof("Retrieving intent for: [%s]", text)
	if len(text) < 2 {
		msg := fmt.Sprintf("GetIntent text length < 2: [%s]", text)
		err := errors.New(msg)
		logrus.WithError(err).Error(msg)
		return nil, err
	}

	// Create basic request with auth headers
	req, err := http.NewRequest("GET", WitBaseUrl + "/message", nil)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to create GetIntent request: %s", text)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + wc.AccessToken)

	// Build query params
	q := req.URL.Query()
	q.Add("v", "20141022")
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

	i := Intent{}
	err = json.Unmarshal(body, &i)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to parse JSON GetIntent response body: %s", text)
		return nil, err
	}

	logrus.Infof("Received intent: %+v", i)
	return &i, nil
}

func NewWitClient(accessToken string) NLPClient {
	if accessToken == "" {
		panic(errors.New("Missing access token to initialize Wit client"))
	}

	client := &http.Client{}
	client.Timeout = time.Second * 5

	return &WitClient {
		AccessToken: accessToken,
		HTTPClient: client,
	}
}
