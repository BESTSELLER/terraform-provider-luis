package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client holds config params
type Client struct {
	Endpoint        string
	AppID           string
	Version         string
	SubscriptionKey string
	httpClient      *http.Client
}

// New creates common settings
func New(client Client) *Client {

	return &Client{
		Endpoint:        client.Endpoint,
		AppID:           client.AppID,
		Version:         client.Version,
		SubscriptionKey: client.SubscriptionKey,
		httpClient:      &http.Client{},
	}
}

// SendRequest with specified method, path and payload, if resposne code does not match the expected it will fail
func (c *Client) SendRequest(method string, path string, payload interface{}, statusCode int) (body []byte, err error) {
	url := fmt.Sprintf("https://%s/luis/authoring/v3.0-preview/apps/%s/versions/%s/", c.Endpoint, c.AppID, c.Version) + path
	client := &http.Client{}

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", c.SubscriptionKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if statusCode != 0 {
		if resp.StatusCode != statusCode {
			return nil, fmt.Errorf("[ERROR] unexpected status code got: %v expected: %v \n %v", resp.StatusCode, statusCode, string(body))
		}
	}

	return body, nil
}
