package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
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

func (c *LinodeClient) APICall(query string) ([]byte, error) {
	response, err := http.Get(API_ENDPOINT + "/?" + query)
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

func (c *LinodeClient) ListLinodes() ([]byte, error) {
	params := url.Values{}
	params.Add("api_key", c.APIKey)
	params.Add("api_action", "linode.list")
	return c.APICall(query)
}
