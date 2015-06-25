package api

import (
	"os"
	"net/url"
)

func GetAPIKeyFromEnvironment() string {
	return os.Getenv("LINODE_KEY")
}

func ValuesWithActionAndKey(key string, action string) *url.Values {
	params := url.Values{}
	params.Add("api_key", key)
	params.Add("api_action", action)
	return &params
}
