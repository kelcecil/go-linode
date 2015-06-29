package api

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

func (c *Linode) Refresh() {
	// Call ListLinodes for the given object's ID
	// Refresh itself.
}
