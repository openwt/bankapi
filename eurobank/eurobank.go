package eurobank

import (
	"encoding/xml"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	EuroBankDayRateURL = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	EuroBank90RateURL  = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
)

type Currency struct {
	Rate float32
	Id   string
}

type DayRate struct {
	Day   time.Time
	Rates []Currency
}

func newRateFromXMLInput(r io.Reader) []*DayRate {

	var dayRates = make([]*DayRate, 0)
	var dayRate *DayRate = nil

	isInDayRate := false
	isInCurrencyValue := false
	isInDayRateList := false

	decoder := xml.NewDecoder(r)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch xmlNode := token.(type) {

		case xml.StartElement:
			if xmlNode.Name.Local == "Cube" {

				isInDayRateList = true

				if len(xmlNode.Attr) == 1 {

					isInDayRate = true
					dayTimeValue := xmlNode.Attr[0].Value
					dayTime, _ := time.Parse("2006-01-02", dayTimeValue)
					dayRate = &DayRate{Day: dayTime}
				} else if len(xmlNode.Attr) == 2 {

					isInCurrencyValue = true

					rate, _ := strconv.ParseFloat(xmlNode.Attr[1].Value, 32)

					currency := Currency{Id: xmlNode.Attr[0].Value, Rate: float32(rate)}
					dayRate.Rates = append(dayRate.Rates, currency)
				}
			}
		case xml.EndElement:

			if xmlNode.Name.Local == "Cube" {

				if isInCurrencyValue && isInDayRate && isInDayRateList {
					isInCurrencyValue = false
				} else if !isInCurrencyValue && isInDayRate && isInDayRateList {
					dayRates = append(dayRates, dayRate)
					isInDayRate = false
				} else if !isInCurrencyValue && !isInDayRate && isInDayRateList {
					isInDayRateList = true
				}

			}
		}
	}
	return dayRates
}

func GetDayRate() (*DayRate, error) {

	resp, err := http.Get(EuroBankDayRateURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return newRateFromXMLInput(resp.Body)[0], nil

}

func Get90DayRates() ([]*DayRate, error) {

	resp, err := http.Get(EuroBank90RateURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return newRateFromXMLInput(resp.Body), nil

}
