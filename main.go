package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const kelvin_zero float64 = -273.15
const mps_to_kmph float64 = 3.6

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	args := os.Args

	api_key := os.Getenv("OPEN_WEATHER_MAP_API")
	zip_code := strings.ToUpper(args[1])
	country_code := strings.ToUpper(args[2])

	api_url := "https://api.openweathermap.org/data/2.5/weather?appid=" + api_key + "&zip=" + zip_code + "," + country_code

	weather_data := getData(api_url)

	fmt.Println("==============================")
	fmt.Println("========  Go Weather  ========")
	fmt.Println("==============================")
	fmt.Println()

	city := weather_data.Name
	lat := weather_data.Coord.Lat
	lon := weather_data.Coord.Lon

	updated_datetime := time.Unix(int64(weather_data.Dt+weather_data.Timezone), 0)

	fmt.Printf("Weather for %s (%v, %v)\n", city, lat, lon)
	fmt.Printf("Last Updated: %v\n", updated_datetime)
	fmt.Println()

	temperature := int(math.Round(weather_data.Main.Temp + kelvin_zero))
	feels_like := int(math.Round(weather_data.Main.FeelsLike + kelvin_zero))
	conditions := weather_data.Weather[0].Main

	fmt.Printf("%vC (Feels like %vC) %s\n", temperature, feels_like, conditions)

	temp_high := int(math.Round(weather_data.Main.TempMax + kelvin_zero))
	temp_low := int(math.Round(weather_data.Main.TempMin + kelvin_zero))

	fmt.Printf("High: %vC  Low: %vC\n", temp_high, temp_low)

	wind_speed := math.Round(weather_data.Wind.Speed * mps_to_kmph)
	wind_dir := deg_to_cardinal(weather_data.Wind.Deg)

	fmt.Printf("Wind: %vkm/h %s\n", wind_speed, wind_dir)

	sunrise := time.Unix(int64(weather_data.Sys.Sunrise+weather_data.Timezone), 0)
	sunset := time.Unix(int64(weather_data.Sys.Sunset+weather_data.Timezone), 0)

	fmt.Printf("Sunrise: %v:%v  Sunset: %v:%v\n", sunrise.Hour(), sunrise.Minute(), sunset.Hour(), sunset.Minute())

}

func getData(api_url string) WeatherData {
	res, err := http.Get(api_url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var weather_data WeatherData

	json.Unmarshal([]byte(body), &weather_data)
	return weather_data
}

func deg_to_cardinal(deg int) string {
	dirs := [16]string{
		"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
	}

	ix := int(math.Round((float64(deg) + 11.25) / 22.5))
	dir_index := ix % 16
	cardinal := dirs[dir_index]
	return cardinal
}
