package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

func fetchCityPayload(city string) (*XMLPayload, int, error) {
	url := fmt.Sprintf("https://meteo.arso.gov.si/uploads/probase/www/fproduct/text/sl/forecast_%s_latest.xml", city)
	response, err := http.Get(url)
	if err != nil {
		return nil, 500, fmt.Errorf("error occurred while making a request to the api: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, response.StatusCode, fmt.Errorf("error occured while fetching data from the api: %s", response.Status)
	}

	payload := &XMLPayload{}
	err = xml.NewDecoder(response.Body).Decode(payload)
	if err != nil {
		return nil, 500, fmt.Errorf("error occurred while decoding the response body: %s", err)
	}

	return payload, 200, nil
}

func payloadToData(p *XMLPayload) City {
	var W_s []Weather
	for _, m := range p.MetData {
		W := Weather{
			Day:   strings.Split(m.Day, " ")[0],
			Valid: m.Valid,
			Icon:  p.IconUrlBase + m.WeatherIcon + "." + p.IconFormat,
			Temp: WeatherTemp{
				Low:  m.TempMin,
				Max:  m.TempMax,
				Unit: m.TempUnit,
			},
			Wind: WeatherWind{
				Icon:          p.IconUrlBase + m.WindIcon + "." + p.IconFormat,
				Direction:     m.WindDirection,
				DirectionLong: m.WindDirectionLong,
				Speed:         m.WindSpeed,
				Unit:          m.WindUnit,
			},
		}
		W_s = append(W_s, W)
	}

	C := City{
		City:    p.MetData[0].City,
		Country: p.MetData[0].Country,
		Weather: W_s,
	}

	return C
}
