package main

import (
	"encoding/json"
	"github.com/kelcecil/go-linode/api"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type LinodeSpec struct {
	ERRORARRAY []string
	DATA       SpecData
	ACTION     string
}

type SpecData struct {
	VERSION float64
	METHODS map[string]Method
}

type Method struct {
	DESCRIPTION string
	PARAMETERS  map[string]Parameter
	THROWS      string
}

type Parameter struct {
	NAME        string
	DESCRIPTION string
	TYPE        string
	REQUIRED    bool
}

type LinodeGenerateMetadata struct {
	Avail map[string]LinodeAvail
}

type LinodeAvail struct {
	Structname string
	Type       string
	Key        string
	Value      string
}

func main() {
	yamlFile, err := ioutil.ReadFile("./info.yaml")
	if err != nil {
		print(err.Error())
	}
	var metadata LinodeGenerateMetadata
	err = yaml.Unmarshal(yamlFile, &metadata)
	if err != nil {
		print(err.Error())
	}
	client := api.NewLinodeClient()
	response, err := client.GetSpec()
	if err != nil {
		print(err)
	}
	var result LinodeSpec
	err = json.Unmarshal(response, &result)
	if err != nil {
		print(err.Error())
	}
	GenerateEnum(client, "avail.datacenters", metadata)
}
