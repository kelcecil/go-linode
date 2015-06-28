package main

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/kelcecil/go-linode/api"
)

var templateFunc = map[string]interface{}{
	"UpperCaseFirstLetter": api.UpperCaseFirstLetter,
	"LabelVariableSafe":    api.LabelVariableSafe,
	"IntToBool":            api.IntToBool,
	"GoSafeString":         api.GoSafeString,
}

// GenerateEnum ...
func GenerateEnum(client *api.LinodeClient, action string, metadata LinodeGenerateMetadata) {
	response, err := client.CallAction(action)
	if err != nil {
		panic(err.Error())
	}
	var result interface{}
	err = json.Unmarshal(response, &result)
	rows := result.(map[string]interface{})
	tmpl, err := template.New("").Funcs(templateFunc).ParseFiles("./generate/datacenters.template", "./generate/distributions.template")

	if err != nil {
		panic(err.Error())
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "distributions.template", rows)
	if err != nil {
		panic(err.Error())
	}
}
