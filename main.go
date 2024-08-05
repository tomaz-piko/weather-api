package main

import (
	"log"
	"strings"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

var supportedCities = map[string]string{
	"reka":  "RIJEKA_OMISALJ",
	"zadar": "ZADAR-PUN",
	"split": "SPLIT_MARJAN",
}

var cache_purge_interval = 5 * time.Minute // How often to purge expired cache items
var cache_item_expiry = 1 * time.Minute    // How long to before cache items expires

func main() {
	gin.SetMode(gin.ReleaseMode) // Uncomment this to enable ReleaseMode
	//router := gin.Default()
	router := gin.New()
	// Disable logging for /weather endpoint
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/weather"),
		gin.Recovery(),
	)

	memoryStore := persist.NewMemoryStore(cache_purge_interval)
	// Routes
	router.GET("/weather", cache.CacheByRequestURI(memoryStore, cache_item_expiry), getWeather)
	router.GET("/weather/:city", cache.CacheByRequestURI(memoryStore, cache_item_expiry), getCityWeather)

	router.Run()
}

// Fetches the payload for all supported citites,
// transforms it into an array of City structs
// and returns it as a JSON response.
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

// Fetches the payload for a specific city,
// transforms it into a City struct
// and returns it as a JSON response.
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
