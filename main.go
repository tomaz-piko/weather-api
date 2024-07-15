package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

var supportedCities = map[string]string{
	"reka":  "RIJEKA_OMISALJ",
	"zadar": "ZADAR-PUN",
	"split": "SPLIT_MARJAN",
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/weather", getWeather)
	router.GET("/weather/:city", getCityWeather)

	router.Run("localhost:8080")
}

func getWeather(c *gin.Context) {
	var data []City
	for _, v := range supportedCities {
		payload, statusCode, err := fetchCityPayload(v)
		if err != nil {
			log.Println(err)
			c.JSON(statusCode, Response[City]{
				Status: statusCode,
				Error:  true,
				Data:   nil,
			})
			return
		}
		cityData := payloadToData(payload)
		data = append(data, cityData)
	}
	c.JSON(200, Response[[]City]{
		Status: 200,
		Error:  false,
		Data:   &data,
	})
}

func getCityWeather(c *gin.Context) {
	city := strings.ToLower(c.Param("city"))
	val, ok := supportedCities[city]
	if !ok {
		log.Println("Attempted to fetch weather for unsupported city.")
		c.JSON(404, Response[City]{
			Status: 404,
			Error:  true,
			Data:   nil,
		})
		return
	}

	payload, statusCode, err := fetchCityPayload(val)
	if err != nil {
		log.Println(err)
		c.JSON(statusCode, Response[City]{
			Status: statusCode,
			Error:  true,
			Data:   nil,
		})
		return
	}

	cityData := payloadToData(payload)
	c.JSON(200, Response[City]{
		Status: 200,
		Error:  false,
		Data:   &cityData,
	})
}
