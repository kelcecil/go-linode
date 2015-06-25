package main

import (
	"encoding/json"
	"github.com/kelcecil/go-linode/api"
	"os"
	"text/template"
)

const (
	sourceTemplate = "package api\n\nconst (\n{{range $value := .DATA}}\t{{.ABBR}} = {{.DATACENTERID}}\n{{end}})"
)

func GenerateEnum(client *api.LinodeClient, action string, metadata LinodeGenerateMetadata) {
	response, err := client.CallAction(action)
	if err != nil {
		panic(err.Error())
	}
	var result interface{}
	err = json.Unmarshal(response, &result)
	rows := result.(map[string]interface{})
	tmpl, err := template.New("avail").Parse(sourceTemplate)
	if err != nil {
		panic(err.Error())
	}
	err = tmpl.Execute(os.Stdout, rows)
	if err != nil {
		panic(err.Error())
	}
}
