package main

import (
	"github.com/kelcecil/go-linode/api"
)

func main() {
	client := api.NewLinodeClient()
	result, err := client.ListLinodes(api.ListLinodesRequest{})
	if err != nil {
		println(err.Error())
	}
	println(result[0].BandwidthInEnabled)
}
