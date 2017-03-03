package main

type TrendData struct {
	Ago   string  `json:"ago"`
	Now   string  `json:"now"`
	Trend float64 `json:"trend"`
}

type CityData struct {
	CityCode int    `json:"citycode"`
	CityName string `json:"cityname"`
}

type StationData struct {
	StationCode  string  `json:"stationcode"`
	PositionName string  `json:"positionname"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

type AQIData struct {
	Time float64 `json:"time"`
	Area string  `json:"area"`
	AQI  float64 `json:"aqi"`
	O3   float64 `json:"o3"`
	CO   float64 `json:"co"`
	SO2  float64 `json:"so2"`
	NO2  float64 `json:"no2"`
	PM25 float64 `json:"pm25"`
	PM10 float64 `json:"pm10"`
}

type AQIDatas []interface{}
