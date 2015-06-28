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

type Linode struct {
	TotalXfer             int    `json:"TOTALXFER"`
	BackupsEnabled        bool   `json:"BACKUPSENABLED"`
	WatchdogEnabled       bool   `json:"WATCHDOG"`
	LpmDisplayGroup       string `json:"LPM_DISPLAYGROUP"`
	Status                LinodeStatus
	StatusString          string `json:"STATUS"`
	TotalRAM              int    `json:"TOTALRAM"`
	TotalHardDisk         int    `json:"TOTALHD"`
	BackupWindow          bool   `json:"BACKUPWINDOW"`
	Label                 string `json:"LABEL"`
	BackupWeeklyDay       int    `json:"BAKCUPWEEKLYDAY"`
	rawHostingDataCenter  int    `json:"DATACENTERID"`
	Id                    int    `json:"LINODEID"`
	rawCreateDateTime     string `json:"CREATE_DT"`
	PlanId                int    `json:"PLANID"`
	DistributionVendor    string `json:"DISTRIBUTIONVENDOR"`
	BandwidthQuotaEnabled bool   `json:"ALERT_BWQUOTA_ENABLED"`
	DiskIOThreshold       int    `json:"ALERT_DISKIO_THRESHOLD"`
	DiskIOEnabled         bool   `json:"ALERT_DISKIO_ENABLED"`
	BandwithOutEnabled    bool   `json:"ALERT_BWOUT_ENABLED"`
	BandwidthOutThreshold int    `json:"ALERT_BWOUT_THRESHOLD"`
	BandwidthInThreshold  int    `json:"ALERT_BWIN_THRESHOLD"`
	BandwidthInEnabled    bool   `json:"ALERT_BWIN_ENABLED"`
	CPUThreshold          int    `json:"ALERT_CPU_THRESHOLD"`
}

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
	data, err := getRawJsons(response)
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

func getRawJsons(response []byte) ([]json.RawMessage, error) {
	topJson := struct {
		ACTION string
		DATA   []json.RawMessage
	}{
		"",
		nil,
	}
	err := json.Unmarshal(response, &topJson)
	if err != nil {
		return nil, err
	}
	return topJson.DATA, nil
}

func createLinode(raw json.RawMessage) Linode {
	println(string(raw))
	newLinode := Linode{}
	json.Unmarshal(raw, &newLinode)
	return newLinode
}
