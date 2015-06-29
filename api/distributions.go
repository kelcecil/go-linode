package api

import (
	"time"
)

type Distribution struct {
	Is64Bit            bool
	Label              string
	MinimumImageSize   int
	DistributionId     int
	CreateDateTime     time.Time
	RequiresVopsKernel bool
}

var Arch_Linux_2015_02_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2015-02-20 14:17:16.0")
var Arch_Linux_2015_02 = Distribution{
	Is64Bit:            true,
	Label:              "Arch_Linux_2015_02",
	MinimumImageSize:   800,
	DistributionId:     138,
	CreateDateTime:     Arch_Linux_2015_02_Time,
	RequiresVopsKernel: true,
}

var CentOS_7_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-07-08 10:07:21.0")
var CentOS_7 = Distribution{
	Is64Bit:            true,
	Label:              "CentOS_7",
	MinimumImageSize:   750,
	DistributionId:     129,
	CreateDateTime:     CentOS_7_Time,
	RequiresVopsKernel: true,
}

var Debian_7_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-09-24 13:59:32.0")
var Debian_7 = Distribution{
	Is64Bit:            true,
	Label:              "Debian_7",
	MinimumImageSize:   600,
	DistributionId:     130,
	CreateDateTime:     Debian_7_Time,
	RequiresVopsKernel: true,
}

var Debian_8_1_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2015-04-27 16:26:41.0")
var Debian_8_1 = Distribution{
	Is64Bit:            true,
	Label:              "Debian_8_1",
	MinimumImageSize:   900,
	DistributionId:     140,
	CreateDateTime:     Debian_8_1_Time,
	RequiresVopsKernel: true,
}

var Fedora_22_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2015-05-26 13:50:58.0")
var Fedora_22 = Distribution{
	Is64Bit:            true,
	Label:              "Fedora_22",
	MinimumImageSize:   650,
	DistributionId:     141,
	CreateDateTime:     Fedora_22_Time,
	RequiresVopsKernel: true,
}

var Gentoo_2014_12_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-12-24 18:00:09.0")
var Gentoo_2014_12 = Distribution{
	Is64Bit:            true,
	Label:              "Gentoo_2014_12",
	MinimumImageSize:   2000,
	DistributionId:     137,
	CreateDateTime:     Gentoo_2014_12_Time,
	RequiresVopsKernel: true,
}

var OpenSUSE_13_2_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-12-17 17:55:42.0")
var OpenSUSE_13_2 = Distribution{
	Is64Bit:            true,
	Label:              "OpenSUSE_13_2",
	MinimumImageSize:   700,
	DistributionId:     135,
	CreateDateTime:     OpenSUSE_13_2_Time,
	RequiresVopsKernel: true,
}

var Slackware_14_1_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2013-11-25 11:11:02.0")
var Slackware_14_1 = Distribution{
	Is64Bit:            true,
	Label:              "Slackware_14_1",
	MinimumImageSize:   875,
	DistributionId:     117,
	CreateDateTime:     Slackware_14_1_Time,
	RequiresVopsKernel: true,
}

var Ubuntu_14_04_LTS_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-04-17 15:42:07.0")
var Ubuntu_14_04_LTS = Distribution{
	Is64Bit:            true,
	Label:              "Ubuntu_14_04_LTS",
	MinimumImageSize:   750,
	DistributionId:     124,
	CreateDateTime:     Ubuntu_14_04_LTS_Time,
	RequiresVopsKernel: true,
}

var Ubuntu_15_04_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2015-04-23 11:08:05.0")
var Ubuntu_15_04 = Distribution{
	Is64Bit:            true,
	Label:              "Ubuntu_15_04",
	MinimumImageSize:   1200,
	DistributionId:     139,
	CreateDateTime:     Ubuntu_15_04_Time,
	RequiresVopsKernel: true,
}

var Arch_Linux_2014_10_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-10-06 15:32:20.0")
var Arch_Linux_2014_10 = Distribution{
	Is64Bit:            true,
	Label:              "Arch_Linux_2014_10",
	MinimumImageSize:   600,
	DistributionId:     132,
	CreateDateTime:     Arch_Linux_2014_10_Time,
	RequiresVopsKernel: true,
}

var CentOS_5_6_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2009-08-17 00:00:00.0")
var CentOS_5_6 = Distribution{
	Is64Bit:            true,
	Label:              "CentOS_5_6",
	MinimumImageSize:   950,
	DistributionId:     60,
	CreateDateTime:     CentOS_5_6_Time,
	RequiresVopsKernel: true,
}

var CentOS_6_5_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-04-28 15:19:34.0")
var CentOS_6_5 = Distribution{
	Is64Bit:            true,
	Label:              "CentOS_6_5",
	MinimumImageSize:   675,
	DistributionId:     127,
	CreateDateTime:     CentOS_6_5_Time,
	RequiresVopsKernel: true,
}

var Debian_6_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2011-02-08 16:54:31.0")
var Debian_6 = Distribution{
	Is64Bit:            true,
	Label:              "Debian_6",
	MinimumImageSize:   550,
	DistributionId:     78,
	CreateDateTime:     Debian_6_Time,
	RequiresVopsKernel: true,
}

var Fedora_20_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2013-01-27 10:00:00.0")
var Fedora_20 = Distribution{
	Is64Bit:            true,
	Label:              "Fedora_20",
	MinimumImageSize:   1536,
	DistributionId:     122,
	CreateDateTime:     Fedora_20_Time,
	RequiresVopsKernel: true,
}

var Fedora_21_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-12-10 16:56:28.0")
var Fedora_21 = Distribution{
	Is64Bit:            true,
	Label:              "Fedora_21",
	MinimumImageSize:   650,
	DistributionId:     134,
	CreateDateTime:     Fedora_21_Time,
	RequiresVopsKernel: true,
}

var Gentoo_2013_11_26_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2013-11-26 15:20:31.0")
var Gentoo_2013_11_26 = Distribution{
	Is64Bit:            true,
	Label:              "Gentoo_2013_11_26",
	MinimumImageSize:   3072,
	DistributionId:     118,
	CreateDateTime:     Gentoo_2013_11_26_Time,
	RequiresVopsKernel: true,
}

var OpenSUSE_13_1_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2013-12-02 12:53:29.0")
var OpenSUSE_13_1 = Distribution{
	Is64Bit:            true,
	Label:              "OpenSUSE_13_1",
	MinimumImageSize:   1024,
	DistributionId:     120,
	CreateDateTime:     OpenSUSE_13_1_Time,
	RequiresVopsKernel: true,
}

var Slackware_13_37_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2011-06-05 15:11:59.0")
var Slackware_13_37 = Distribution{
	Is64Bit:            true,
	Label:              "Slackware_13_37",
	MinimumImageSize:   600,
	DistributionId:     87,
	CreateDateTime:     Slackware_13_37_Time,
	RequiresVopsKernel: true,
}

var Ubuntu_12_04_LTS_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-04-28 14:16:59.0")
var Ubuntu_12_04_LTS = Distribution{
	Is64Bit:            true,
	Label:              "Ubuntu_12_04_LTS",
	MinimumImageSize:   550,
	DistributionId:     126,
	CreateDateTime:     Ubuntu_12_04_LTS_Time,
	RequiresVopsKernel: true,
}

var Ubuntu_14_10_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2014-10-24 15:48:04.0")
var Ubuntu_14_10 = Distribution{
	Is64Bit:            true,
	Label:              "Ubuntu_14_10",
	MinimumImageSize:   650,
	DistributionId:     133,
	CreateDateTime:     Ubuntu_14_10_Time,
	RequiresVopsKernel: true,
}

var Slackware_13_37_32bit_Time, _ = time.Parse("2007-04-18 00:00:00.0", "2011-06-05 15:11:59.0")
var Slackware_13_37_32bit = Distribution{
	Is64Bit:            false,
	Label:              "Slackware_13_37_32bit",
	MinimumImageSize:   600,
	DistributionId:     86,
	CreateDateTime:     Slackware_13_37_32bit_Time,
	RequiresVopsKernel: true,
}

var Distributions = map[string]Distribution{
	"Arch_Linux_2015_02":    Arch_Linux_2015_02,
	"CentOS_7":              CentOS_7,
	"Debian_7":              Debian_7,
	"Debian_8_1":            Debian_8_1,
	"Fedora_22":             Fedora_22,
	"Gentoo_2014_12":        Gentoo_2014_12,
	"OpenSUSE_13_2":         OpenSUSE_13_2,
	"Slackware_14_1":        Slackware_14_1,
	"Ubuntu_14_04_LTS":      Ubuntu_14_04_LTS,
	"Ubuntu_15_04":          Ubuntu_15_04,
	"Arch_Linux_2014_10":    Arch_Linux_2014_10,
	"CentOS_5_6":            CentOS_5_6,
	"CentOS_6_5":            CentOS_6_5,
	"Debian_6":              Debian_6,
	"Fedora_20":             Fedora_20,
	"Fedora_21":             Fedora_21,
	"Gentoo_2013_11_26":     Gentoo_2013_11_26,
	"OpenSUSE_13_1":         OpenSUSE_13_1,
	"Slackware_13_37":       Slackware_13_37,
	"Ubuntu_12_04_LTS":      Ubuntu_12_04_LTS,
	"Ubuntu_14_10":          Ubuntu_14_10,
	"Slackware_13_37_32bit": Slackware_13_37_32bit,
}