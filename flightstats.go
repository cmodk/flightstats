package flightstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cmodk/go-simplehttp"
	"github.com/sirupsen/logrus"
)

type FlightStats struct {
	lg     *logrus.Logger
	sh     simplehttp.SimpleHttp
	key    string
	app_id string
	debug  bool
}

func New(id string, k string, logger *logrus.Logger) *FlightStats {
	return &FlightStats{
		sh:     simplehttp.New("https://api.flightstats.com/flex", logger),
		lg:     logger,
		key:    k,
		app_id: id,
	}
}

func (fs *FlightStats) SetDebug(d bool) {
	fs.debug = d
	fs.sh.SetDebug(d)
}

type FlightStatsSchedule struct {
	ScheduledFlights []ScheduledFlight `json:"scheduledFlights"`
	Appendix         Appendix          `json:"appendix"`
}

func (fss FlightStatsSchedule) GetAirport(code string) (AirportAppendix, error) {
	for _, aa := range fss.Appendix.Airports {
		if aa.Fs == code {
			return aa, nil
		}
	}

	return AirportAppendix{}, errors.New(fmt.Sprintf("Unknown airport code: %s\n", code))
}

func (fs *FlightStats) GetScheduleArrival(arrival time.Time, airline string, flight_number int) (FlightStatsSchedule, error) {

	url := fmt.Sprintf("/schedules/rest/v1/json/flight/%s/%d/arriving/%d/%d/%d?appId=%s&appKey=%s",
		airline,
		flight_number,
		arrival.Year(),
		arrival.Month(),
		arrival.Day(),
		fs.app_id,
		fs.key)

	resp, err := fs.sh.Get(url)
	if err != nil {
		return FlightStatsSchedule{}, err
	}

	var s FlightStatsSchedule

	if err := json.Unmarshal([]byte(resp), &s); err != nil {
		return FlightStatsSchedule{}, err
	}

	return s, nil
}
