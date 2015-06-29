package api

type DataCenter struct {
	DataCenterId int
	Location     string
	Abbreviation string
}

var Dallas = DataCenter{
	DataCenterId: 2,
	Location:     "Dallas, TX, USA",
	Abbreviation: "dallas",
}

var Fremont = DataCenter{
	DataCenterId: 3,
	Location:     "Fremont, CA, USA",
	Abbreviation: "fremont",
}

var Atlanta = DataCenter{
	DataCenterId: 4,
	Location:     "Atlanta, GA, USA",
	Abbreviation: "atlanta",
}

var Newark = DataCenter{
	DataCenterId: 6,
	Location:     "Newark, NJ, USA",
	Abbreviation: "newark",
}

var London = DataCenter{
	DataCenterId: 7,
	Location:     "London, England, UK",
	Abbreviation: "london",
}

var Tokyo = DataCenter{
	DataCenterId: 8,
	Location:     "Tokyo, JP",
	Abbreviation: "tokyo",
}

var Singapore = DataCenter{
	DataCenterId: 9,
	Location:     "Singapore, SG",
	Abbreviation: "singapore",
}

var DataCenters = map[string]DataCenter{
	"Dallas":    Dallas,
	"Fremont":   Fremont,
	"Atlanta":   Atlanta,
	"Newark":    Newark,
	"London":    London,
	"Tokyo":     Tokyo,
	"Singapore": Singapore,
}