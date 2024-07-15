package main

type XMLPayload struct {
	IconUrlBase string       `xml:"icon_url_base"`
	IconFormat  string       `xml:"icon_format"`
	MetData     []XMLMetData `xml:"metData"`
}

type XMLMetData struct {
	City              string  `xml:"domain_longTitle"`
	Country           string  `xml:"domain_countryIsoCode2"`
	Valid             string  `xml:"valid"`
	Day               string  `xml:"valid_day"`
	WeatherIcon       string  `xml:"nn_icon"` // kratica (ime ikone, ki jo uporablja meteo.si) za količino oblačnosti oz. videz neba;
	TempMin           int64   `xml:"tnsyn_degreesC"`
	TempMax           int64   `xml:"txsyn_degreesC"`
	TempUnit          string  `xml:"tnsyn_var_unit"`
	WindIcon          string  `xml:"ddff_icon"` // kratica (ime ikone, ki jo uporablja meteo.si) za vremenski pojav (pri napovedih samo osnovni vremenski pojav, pri opazovanjih pa je vključena tudi moč pojava):
	WindSpeed         float64 `xml:"ff_val"`
	WindUnit          string  `xml:"ff_var_unit"`
	WindDirection     string  `xml:"dd_shortText"`
	WindDirectionLong string  `xml:"dd_longText"`
}
