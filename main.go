package main

import (
	"encoding/json"
	"fmt"
	"github.com/kelcecil/go-linode/api"
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

func main() {
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
	for k, _ := range result.DATA.METHODS {
		fmt.Printf("%s\n", k)
	}
}
