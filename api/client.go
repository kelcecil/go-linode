package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LinodeClient struct {
	APIEndpoint  string
	APIKey       string
	batchEnabled bool
}

func NewLinodeClient() *LinodeClient {
	return NewLinodeClientWithKey(GetAPIKeyFromEnvironment())
}

func NewLinodeClientWithKey(apiKey string) *LinodeClient {
	return &LinodeClient{
		APIEndpoint: "https://api.linode.com",
		APIKey:      apiKey,
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
	response, err := http.Get(c.APIEndpoint + "/?" + query)
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

func GetRawJson(response []byte) (json.RawMessage, error) {
	topJson := struct {
		ACTION string
		DATA   json.RawMessage
	}{
		"",
		nil,
	}
	err := json.Unmarshal(response, &topJson)
	if err != nil {
		return nil, err
	}
	return topJson.DATA, nil
}

func GetRawJsons(response []byte) ([]json.RawMessage, error) {
	topJson := struct {
		ACTION string
		DATA   []json.RawMessage
	}{
		"",
		nil,
	}
	err := json.Unmarshal(response, &topJson)
	if err != nil {
		return nil, err
	}
	return topJson.DATA, nil
}
