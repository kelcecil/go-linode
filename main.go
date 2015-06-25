package main

import (
	"github.com/kelcecil/go-linode/api"
)

func main() {
	client := api.NewLinodeClient()
	response, err := client.ListLinodes()
	if err != nil {
		print(err)
	}
	print(string(response[:]))
}
