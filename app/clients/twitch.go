package clients

import (
	"github.com/go-resty/resty"
	"github.com/sirupsen/logrus"
	"time"
	"fmt"
)

type TwitchClient interface {
	GetUserByLogin(login string) (*User, error)
	GetChannelFollows(id string) (*GetChannelFollowsResponse, error)
}

type TwitchV5Client struct {
	ClientID string
}

type GetUserByLoginResponse struct {
	Users []User `json:"users"`
}

type User struct {
	ID string `json:"_id"`
	Name string `json:"name"`
}

type GetChannelFollowsResponse struct {
	Follows []Follow `json:"follows"`
}

type Follow struct {
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
	User          User      `json:"user"`
}

func (c *TwitchV5Client) GetUserByLogin(login string) (*User, error) {
	usersResp := GetUserByLoginResponse{}

	_, err := resty.R().
		SetHeader("Accept", "application/vnd.twitchtv.v5+json").
		SetQueryParam("client_id", c.ClientID).
		SetQueryParam("login", login).
		SetResult(&usersResp).
		Get("https://api.twitch.tv/kraken/users")

	if err != nil {
		logrus.WithError(err).Errorf("Failed to retrieve user by login: %s", login)
		return nil, err
	}

	for _, u := range usersResp.Users {
		if u.Name == login {
			logrus.Infof("Found user: %+v", usersResp)
			return &u, nil
		}
	}
	logrus.Errorf("Failed to find user for login: %s", login)
	return nil, err
}

// Get the latest 100 followers
func (c *TwitchV5Client) GetChannelFollows(id string) (*GetChannelFollowsResponse, error) {
	followsResp := GetChannelFollowsResponse{}

	_, err := resty.R().
		SetHeader("Accept", "application/vnd.twitchtv.v5+json").
		SetQueryParams(map[string]string {
			"client_id": c.ClientID,
			"limit": "100",
		}).
		SetResult(&followsResp).
		Get(fmt.Sprintf("https://api.twitch.tv/kraken/channels/%s/follows", id))

	if err != nil {
		logrus.WithError(err).Errorf("Failed to get channel follows for id: %s", id)
		return nil, err
	}

	logrus.Infof("Returning channel followers: %+v", followsResp)
	return &followsResp, nil
}

func NewTwitchV5Client(clientID string) TwitchClient {
	return &TwitchV5Client{
		ClientID: clientID,
	}
}