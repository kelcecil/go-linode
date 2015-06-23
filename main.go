package main

import (
	"github.com/kelcecil/go-linode/api"
)

func main() {
	client := api.NewLinodeClientWithKey("xzDXtXFwWWxLlDSAWqmwT8p2YOxZy9qFt485srdmiY9dhKWQGYBevzdDRVFPaZ1G")
	response, err := client.ListLinodes()
	if err != nil {
		print(err)
	}
	print(string(response[:]))
}
