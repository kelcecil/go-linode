package api

import (
	"io/ioutil"
	"net/http"
)

type LinodeClient struct {
	APIKey string
}

func NewLinodeClient() *LinodeClient {
	return NewLinodeClient()
}

func NewLinodeClientWithKey(apiKey string) *LinodeClient {
	return &LinodeClient{
		APIKey: apiKey,
	}
}

func (c *LinodeClient) ListLinodes() ([]byte, error) {
	response, err := http.Get(API_ENDPOINT + "/?api_key=" + c.APIKey +
		"&api_action=linode.list")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
