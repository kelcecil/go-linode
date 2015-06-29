package main

import (
	"bytes"
	"encoding/json"
	"github.com/kelcecil/go-linode/api"
	"go/format"
	"io/ioutil"
	"text/template"
)

var templateFunc = map[string]interface{}{
	"UpperCaseFirstLetter": api.UpperCaseFirstLetter,
	"LabelVariableSafe":    api.LabelVariableSafe,
	"IntToBool":            api.IntToBool,
	"GoSafeString":         api.GoSafeString,
}

func main() {
	client := api.NewLinodeClient()
	templates := loadTemplates()
	GenerateAvailData(client, templates, "avail.datacenters", "datacenters.template", "./api/datacenters.go")
	GenerateAvailData(client, templates, "avail.distributions", "distributions.template", "./api/distributions.go")
	GenerateAvailData(client, templates, "avail.linodeplans", "plans.template", "./api/plans.go")
}

func GenerateAvailData(client *api.LinodeClient, templates *template.Template, action string, tmpl string, fileDestination string) {
	data, err := client.CallAction(action)
	var buf bytes.Buffer
	if err != nil {
		return
	}
	spec := struct {
		DATA interface{}
	}{
		nil,
	}
	json.Unmarshal(data, &spec)
	err = templates.ExecuteTemplate(&buf, tmpl, spec)
	if err != nil {
		println(err.Error())
	}
	finalSource, err := format.Source(buf.Bytes())
	if err != nil {
		println(err.Error())
		return
	}
	ioutil.WriteFile(fileDestination, finalSource, 0777)
	println(string(finalSource))
}

func loadTemplates() *template.Template {
	tmpl, err := template.New("").Funcs(templateFunc).ParseGlob("./generate/*.template")
	if err != nil {
		panic(err.Error())
	}
	return tmpl
}
