package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to use AQI api!\n")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 - Not Found")
}

func ReturnSperifiedData(w http.ResponseWriter, r *http.Request) {
	// 输入格式为"2006010215"
	t := mux.Vars(r)["time"]
	// 格式化时间戳为"2006-01-02 :15:00"
	ft := t[:4] + "-" + t[4:6] + "-" + t[6:8] + " " + t[8:] + ":00"
	city := mux.Vars(r)["city"]

	result := OneCitySingleData(ft, city)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func ReturnNowData(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15") + ":00"
	city := mux.Vars(r)["city"]
	result := OneCitySingleData(t, city)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func ReturnAllDataToday(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02")
	city := mux.Vars(r)["city"]

	result := OneCityAllDataToday(t, city)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func ReturnDataOfCities(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	var locations []string
	t := time.Now().Format("2006-01-02 15") + ":00"

	for _, v := range q {
		locations = append(locations, v[0])
	}
	result := CompareDataOfCities(t, locations)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func ReturnTrendData(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15") + ":00"
	city := mux.Vars(r)["city"]

	result := TrendDataNow(t, city)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func ReturnCityTable(w http.ResponseWriter, r *http.Request) {
	// t := time.Now().Format("2006-01-02 15") + ":00"

	result := CityTable()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func ReturnOneCityAllStation(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]

	result := StationTable(city)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
