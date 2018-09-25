package flightstats

import (
	"github.com/cmodk/jsontime"
)

type ScheduledFlight struct {
	CarrierFsCode          string           `json:"carrierFsCode"`
	FlightNumber           string           `json:"flightNumber"`
	DepartureAirportFsCode string           `json:"departureAirportFsCode"`
	ArrivalAirportFsCode   string           `json:"arrivalAirportFsCode"`
	Stops                  int              `json:"stops"`
	ArrivalTerminal        string           `json:"arrivalTerminal"`
	DepartureTime          jsontime.ISO8601 `json:"departureTime"`
	ArrivalTime            jsontime.ISO8601 `json:"arrivalTime"`
}
