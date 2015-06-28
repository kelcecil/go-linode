package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type LinodeClient struct {
	APIKey       string
	batchEnabled bool
}

func NewLinodeClient() *LinodeClient {
	return NewLinodeClientWithKey(GetAPIKeyFromEnvironment())
}

func NewLinodeClientWithKey(apiKey string) *LinodeClient {
	return &LinodeClient{
		APIKey: apiKey,
	}
}

func (c *LinodeClient) InitialValues(action string) url.Values {
	params := url.Values{}
	params.Add("api_key", c.APIKey)
	params.Add("api_action", "linode.list")
	return params
}

func (c *LinodeClient) HandleCall(query string) ([]byte, error) {
	return c.ImmediateCall(query)
}

func (c *LinodeClient) ImmediateCall(query string) ([]byte, error) {
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

func (c *LinodeClient) CallAction(action string) ([]byte, error) {
	params := ValuesWithActionAndKey(c.APIKey, action)
	return c.HandleCall(params.Encode())
}
