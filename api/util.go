package api

import (
	"net/url"
	"os"
	"strings"
)

// GetAPIKeyFromEnvironment ...
func GetAPIKeyFromEnvironment() string {
	return os.Getenv("LINODE_KEY")
}

// ValuesWithActionAndKey ...
func ValuesWithActionAndKey(key string, action string) *url.Values {
	params := url.Values{}
	params.Add("api_key", key)
	params.Add("api_action", action)
	return &params
}

// UpperCaseFirstLetter ...
func UpperCaseFirstLetter(name string) string {
	letter := strings.ToUpper(string(name[0]))
	return letter + name[1:]
}

func LabelVariableSafe(label string) string {
	label = strings.Replace(label, " ", "_", -1)
	label = strings.Replace(label, ".", "_", -1)
	label = strings.Replace(label, "-", "_", -1)
	return label
}

func GoSafeString(label string) string {
	return UpperCaseFirstLetter(LabelVariableSafe(label))
}

func IntToBool(value float64) string {
	if value == 0 {
		return "false"
	}
	return "true"
}
