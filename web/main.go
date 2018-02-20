package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable,omitempty"`
	DarkSkyStations    []string `json:"darksky-stations,omitempty"`
	DataPointStations  []string `json:"datapoint-stations,omitempty"`
	ISDStations        []string `json:"isds-stations,omitempty"`
	LAMPStations       []string `json:"lamp-stations,omitempty"`
	MADISStations      []string `json:"madis-stations,omitempty"`
	METARStations      []string `json:"metars-stations,omitempty"`
	METNOLicense       string   `json:"metnol-license,omitempty"`
	Sources            []string `json:"sources,omitempty"`
	Units              string   `json:"units,omitempty"`
}

type DataPoint struct {
	Time                       int64   `json:"time,omitempty"`
	Summary                    string  `json:"summary,omitempty"`
	Icon                       string  `json:"icon,omitempty"`
	SunriseTime                int64   `json:"sunriseTime,omitempty"`
	SunsetTime                 int64   `json:"sunsetTime,omitempty"`
	PrecipIntensity            float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityMax         float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime     int64   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability          float64 `json:"precipProbability,omitempty"`
	PrecipType                 string  `json:"precipType,omitempty"`
	PrecipAccumulation         float64 `json:"precipAccumulation,omitempty"`
	Temperature                float64 `json:"temperature,omitempty"`
	TemperatureMin             float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime         int64   `json:"temperatureMinTime,omitempty"`
	TemperatureMax             float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime         int64   `json:"temperatureMaxTime,omitempty"`
	ApparentTemperature        float64 `json:"apparentTemperature,omitempty"`
	ApparentTemperatureMin     float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime int64   `json:"apparentTemperatureMinTime,omitempty"`
	ApparentTemperatureMax     float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime int64   `json:"apparentTemperatureMaxTime,omitempty"`
	NearestStormBearing        float64 `json:"nearestStormBearing,omitempty"`
	NearestStormDistance       float64 `json:"nearestStormDistance,omitempty"`
	DewPoint                   float64 `json:"dewPoint,omitempty"`
	WindSpeed                  float64 `json:"windSpeed,omitempty"`
	WindBearing                float64 `json:"windBearing,omitempty"`
	CloudCover                 float64 `json:"cloudCover,omitempty"`
	Humidity                   float64 `json:"humidity,omitempty"`
	Pressure                   float64 `json:"pressure,omitempty"`
	Visibility                 float64 `json:"visibility,omitempty"`
	Ozone                      float64 `json:"ozone,omitempty"`
	MoonPhase                  float64 `json:"moonPhase,omitempty"`
}

type DataBlock struct {
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
	Data    []DataPoint `json:"data,omitempty"`
}

type alert struct {
	Title       string   `json:"title,omitempty"`
	Regions     []string `json:"regions,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Description string   `json:"description,omitempty"`
	Time        int64    `json:"time,omitempty"`
	Expires     float64  `json:"expires,omitempty"`
	URI         string   `json:"uri,omitempty"`
}

type Forecast struct {
	Latitude  float64   `json:"latitude,omitempty"`
	Longitude float64   `json:"longitude,omitempty"`
	Timezone  string    `json:"timezone,omitempty"`
	Offset    float64   `json:"offset,omitempty"`
	Currently DataPoint `json:"currently,omitempty"`
	Minutely  DataBlock `json:"minutely,omitempty"`
	Hourly    DataBlock `json:"hourly,omitempty"`
	Daily     DataBlock `json:"daily,omitempty"`
	Alerts    []alert   `json:"alerts,omitempty"`
	Flags     Flags     `json:"flags,omitempty"`
	APICalls  int       `json:"apicalls,omitempty"`
	Code      int       `json:"code,omitempty"`
}

func FromJSON(reader io.Reader) (*Forecast, error) {
	var f Forecast
	if err := json.NewDecoder(reader).Decode(&f); err != nil {
		return nil, err
	}

	return &f, nil
}

var t template.Template

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type Address struct {
	addr string
}

func main() {
	tmpl := template.Must(template.ParseFiles("web/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	t.Execute(w, Forecast{})
	log.Printf("0004\n\n")

	//response := ResponseMessage{Message: "Hello, " + name, Code: 200}
	//JsonResponseWrite(w, response, 200)
}
