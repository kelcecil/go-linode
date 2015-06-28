package api

import (
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
	TotalXfer          int
	BackupsEnabled     bool
	WatchdogEnabled    bool
	LpmDisplayGroup    string
	Status             LinodeStatus
	TotalRAM           int
	TotalHardDisk      int
	BackupWindow       bool
	Label              string
	BackupWeeklyDay    int
	HostingDataCenter  int
	Id                 int
	CreateDateTime     string
	PlanId             int
	DistributionVendor string
	Alerts             AlertInformation
}

type AlertInformation struct {
	BandwidthQuotaEnabled bool
	DiskIOThreshold       int
	DiskIOEnabled         bool
	BandwithOutEnabled    bool
	BandwidthOutThreshold int
	BWinThreshold         int
	BWinEnabled           bool
	CPUThreshold          int
}

func (c *LinodeClient) ListLinodes(request ListLinodesRequest) ([]byte, error) {
	params := c.InitialValues("linode.list")
	if request.LinodeId != 0 {
		params.Add("LinodeId", string(request.LinodeId))
	}
	return c.ImmediateCall(params.Encode())
}
