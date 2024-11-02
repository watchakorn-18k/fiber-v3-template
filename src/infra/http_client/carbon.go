package httpclient

import (
	"errors"
	"fiber-v3-template/src/domain/entities"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3/client"
)

type carbonProvider struct {
	url    string
	client *client.Client
}

type ICarbonProvider interface {
	GetCarbon(url string) (*entities.CarbonResponse, error)
}

func NewCarbonProvider() carbonProvider {
	client := client.New()
	client.SetTimeout(10 * time.Second)
	return carbonProvider{
		url:    "https://api.websitecarbon.com/site?url=",
		client: client,
	}
}

func (c carbonProvider) GetCarbon(url string) (*entities.CarbonResponse, error) {
	client := c.client

	// Get request
	resp, err := client.Get(c.url + url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New("Invalid response status code")
	}
	bodyRaw := resp.Body()
	var carbonResp entities.CarbonResponse
	if err := json.Unmarshal(bodyRaw, &carbonResp); err != nil {
		return nil, err
	}
	return &carbonResp, nil
}
