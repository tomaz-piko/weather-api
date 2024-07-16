package main

// Use of a generic type to support []City | City | nil
type Response[T any] struct {
	Status int  `json:"status"`
	Error  bool `json:"error"`
	Data   *T   `json:"data,omitempty"`
}

type City struct {
	City    string    `json:"city"`
	Country string    `json:"country"`
	Weather []Weather `json:"weather"`
}

type Weather struct {
	Day   string      `json:"day"`
	Valid string      `json:"valid"`
	Icon  string      `json:"icon"`
	Temp  WeatherTemp `json:"temp"`
	Wind  WeatherWind `json:"wind"`
}

type WeatherTemp struct {
	Low  int16  `json:"low"`
	Max  int16  `json:"max"`
	Unit string `json:"unit"`
}

type WeatherWind struct {
	Icon          string  `json:"icon"`
	Direction     string  `json:"direction"`
	DirectionLong string  `json:"directionLong"`
	Speed         float32 `json:"speed"`
	Unit          string  `json:"unit"`
}
