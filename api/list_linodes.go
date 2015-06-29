package api

import (
	"encoding/json"
	_ "time"
)

type ListLinodesRequest struct {
	LinodeId int
}

type LinodeStatus string

const (
	BeingCreated LinodeStatus = "-1"
	BrandNew     LinodeStatus = "0"
	Running      LinodeStatus = "1"
	PoweredOff   LinodeStatus = "2"
)

func (c *LinodeClient) ListLinodes(request ListLinodesRequest) ([]Linode, error) {
	params := c.InitialValues("linode.list")
	if request.LinodeId != 0 {
		params.Add("LinodeId", string(request.LinodeId))
	}
	response, err := c.ImmediateCall(params.Encode())
	if err != nil {
		return nil, err
	}
	activeLinodes, err := createLinodeObjects(response)
	if err != nil {
		return nil, err
	}
	return activeLinodes, nil
}

func createLinodeObjects(response []byte) ([]Linode, error) {
	linodes := make([]Linode, 0)
	data, err := GetRawJsons(response)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(data); i++ {
		linodeBox := createLinode(data[i])
		if err != nil {
			continue
		}
		linodes = append(linodes, linodeBox)
	}
	return linodes, nil
}

func createLinode(raw json.RawMessage) Linode {
	println(string(raw))
	newLinode := Linode{}
	json.Unmarshal(raw, &newLinode)
	return newLinode
}
