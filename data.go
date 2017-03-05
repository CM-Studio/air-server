package main

type TrendData struct {
	Ago   string `json:"ago"`
	Now   string `json:"now"`
	Trend int    `json:"trend"`
}

type CityData struct {
	City string `json:"city"`
}

type StationData struct {
	StationCode  string  `json:"stationcode"`
	PositionName string  `json:"positionname"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

type AQIData struct {
	Time  string  `json:"time"`
	Area  string  `json:"area"`
	AQI   int     `json:"aqi"`
	Trend int     `json:"trend"`
	O3    int     `json:"o3"`
	CO    float64 `json:"co"`
	SO2   int     `json:"so2"`
	NO2   int     `json:"no2"`
	PM25  int     `json:"pm25"`
	PM10  int     `json:"pm10"`
}

type AQIDatas []interface{}
