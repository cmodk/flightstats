package flightstats

type Appendix struct {
	Airlines   []AirlineAppendix   `json:"airlines"`
	Airports   []AirportAppendix   `json:"airports"`
	Equipments []EquipmentAppendix `json:"equipments"`
}

type AirlineAppendix struct {
	Fs     string `json:"fs"`
	IATA   string `json:"iata"`
	ICAO   string `json:"icao"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type AirportAppendix struct {
	Fs        string  `json:"fs"`
	IATA      string  `json:"iata"`
	ICAO      string  `json:"icao"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EquipmentAppendix struct {
}
