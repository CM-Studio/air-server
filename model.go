package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var (
	driverName     string
	dataSourceName string
)

func init() {
	driverName = "mysql"
	dataSourceName = "root:@tcp(127.0.0.1:3306)/airx?charset=utf8"
}

func Conn() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
	}
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func OneCitySingleData(t string, location string) (datas AQIDatas) {
	db := Conn()
	defer db.Close()

	rows, err := db.Query("SELECT avg(TimePoint),Area,avg(AQI),avg(O3),avg(CO),avg(SO2),avg(NO2),avg(PM2_5),avg(PM10) FROM airx.raw where TimePoint='" + t + "'  and Area='" + location + "';")
	checkErr(err)
	defer rows.Close()

	var (
		time float64
		area string
		aqi  float64
		o3   float64
		co   float64
		so2  float64
		no2  float64
		pm25 float64
		pm10 float64
	)

	var s sql.NullString
	err = db.QueryRow("SELECT Area FROM airx.raw where TimePoint=? and Area=? limit 1;", t, location).Scan(&s)

	if s.Valid {
		for rows.Next() {
			err := rows.Scan(&time, &area, &aqi, &o3, &co, &so2, &no2, &pm25, &pm10)
			checkErr(err)
			data := AQIData{time, area, aqi, o3, co, so2, no2, pm25, pm10}
			datas = append(datas, data)
		}
	} else {
		errMsg := jsonErr{Code: 404, Text: "没有" + location + "此时的空气质量数据！"}
		datas = append(datas, errMsg)
	}

	return
}

/*
	输入时间戳和城市名，返回一个城市当天每个时间段的AQI信息，没有则返回错误信息
	示例：
		[
			...
			{
				"code": 404,
				"text": "没有成都市2017-03-02 11:00的空气质量数据！"
			},
			{
				"time": 20170302120000,
				"area": "成都市",
				"aqi": 71.75,
				"o3": 112.375,
				"co": 0.9375,
				"so2": 13.75,
				"no2": 71.5,
				"pm25": 46.375,
				"pm10": 92.5
			},
			...
		]
*/
func OneCityAllDataToday(t string, location string) (datas AQIDatas) {
	db := Conn()
	defer db.Close()

	tmp := time.Now()
	for i := 0; i <= tmp.Hour(); i++ {
		// 针对不同情况拼接字符串
		var h string
		if i < 10 {
			h = t + " " + "0" + strconv.Itoa(i) + ":00"
		} else {
			h = t + " " + strconv.Itoa(i) + ":00"
		}

		rows, err := db.Query("SELECT avg(TimePoint),Area,avg(AQI),avg(O3),avg(CO),avg(SO2),avg(NO2),avg(PM2_5),avg(PM10) FROM airx.raw where TimePoint='" + h + "'  and Area='" + location + "';")
		checkErr(err)
		defer rows.Close()

		var (
			time float64
			area string
			aqi  float64
			o3   float64
			co   float64
			so2  float64
			no2  float64
			pm25 float64
			pm10 float64
		)

		var s sql.NullString
		err = db.QueryRow("SELECT Area FROM airx.raw where TimePoint=? and Area=? limit 1;", h, location).Scan(&s)

		if s.Valid {
			for rows.Next() {
				err := rows.Scan(&time, &area, &aqi, &o3, &co, &so2, &no2, &pm25, &pm10)
				checkErr(err)
				data := AQIData{time, area, aqi, o3, co, so2, no2, pm25, pm10}
				datas = append(datas, data)
			}
		} else {
			errMsg := jsonErr{Code: 404, Text: "没有" + location + h + "的空气质量数据！"}
			datas = append(datas, errMsg)
		}

	}
	return
}

/*
	输入时间戳和城市名称数组，返回数组中城市的当前AQI数据
	单个条目：
			{
				"time": 20170302130000,
				"area": "成都市",
				"aqi": 75.375,
				"o3": 112.375,
				"co": 0.9625,
				"so2": 15.5,
				"no2": 74.875,
				"pm25": 50.25,
				"pm10": 97.625
			}

	没有数据则返回错误信息
*/
func CompareDataOfCities(t string, locations []string) (datas AQIDatas) {
	db := Conn()
	defer db.Close()

	for i := 0; i < len(locations); i++ {
		var (
			time float64
			area string
			aqi  float64
			o3   float64
			co   float64
			so2  float64
			no2  float64
			pm25 float64
			pm10 float64
		)

		rows, err := db.Query("SELECT avg(TimePoint),Area,avg(AQI),avg(O3),avg(CO),avg(SO2),avg(NO2),avg(PM2_5),avg(PM10) FROM airx.raw where TimePoint='" + t + "'  and Area='" + locations[i] + "';")

		checkErr(err)
		defer rows.Close()

		var s sql.NullString
		err = db.QueryRow("SELECT Area FROM airx.raw where TimePoint=? and Area=? limit 1;", t, locations[i]).Scan(&s)

		if s.Valid {
			for rows.Next() {
				err := rows.Scan(&time, &area, &aqi, &o3, &co, &so2, &no2, &pm25, &pm10)
				checkErr(err)
				data := AQIData{time, area, aqi, o3, co, so2, no2, pm25, pm10}
				datas = append(datas, data)
			}
		} else {
			errMsg := jsonErr{Code: 404, Text: "没有" + locations[i] + "此时的空气质量数据！"}
			datas = append(datas, errMsg)
		}
	}
	return
}

/*
	输入字符串格式的时间戳和城市名
	返回：
		{
			"ago":		前一小时的时间戳
			"now": 		当前时间戳
			"trend":	与前一小时相比变化的AQI数值，负数代表下降
		}

	默认如果没有当前或前一小时的AQI数据，trend的值为0

	不论大小写，函数中参数名带"b"为"before"，代指前一个小时的相关参数；"n"为"now"，代指当前时间的相关参数
*/
func TrendDataNow(t string, location string) (datas AQIDatas) {
	db := Conn()
	defer db.Close()

	var aqiN, aqiB int

	tnh := time.Now().Hour()
	tbh := tnh - 1

	tn := t + " " + strconv.Itoa(tnh) + ":00"
	tb := t + " " + strconv.Itoa(tbh) + ":00"

	rowsN, err := db.Query("SELECT avg(AQI) FROM airx.raw where TimePoint='" + tn + "'  and Area='" + location + "';")
	checkErr(err)
	defer rowsN.Close()

	var sN sql.NullString
	err = db.QueryRow("SELECT Area FROM airx.raw where TimePoint=? and Area=? limit 1;", tn, location).Scan(&sN)

	if sN.Valid {
		err := rowsN.Scan(&aqiN)
		checkErr(err)

		rowsB, err := db.Query("SELECT avg(AQI) FROM airx.raw where TimePoint='" + tb + "'  and Area='" + location + "';")
		checkErr(err)
		defer rowsB.Close()

		var sB sql.NullString
		err = db.QueryRow("SELECT Area FROM airx.raw where TimePoint=? and Area=? limit 1;", tb, location).Scan(&sB)

		if sB.Valid {
			err := rowsB.Scan(&aqiB)
			checkErr(err)

			data := TrendData{tb, tn, (aqiN - aqiB)}
			datas = append(datas, data)
		} else {
			// 没有上一小时数据
			data := TrendData{tb, tn, 0}
			datas = append(datas, data)
		}
	} else {
		// 没有本时间的数据
		data := TrendData{tb, tn, 0}
		datas = append(datas, data)
	}
	return
}

/*
	输入时间戳，返回此时的城市表
	单个条目：
			{
				"code": 城市代码（有的城市没有，为0）
				"city": 城市名称
			}

	没有当前城市表则返回相关错误信息
*/
func CityTable() (datas AQIDatas) {
	db := Conn()
	defer db.Close()

	// 这里可以有个排序，按照城市代码
	rows, err := db.Query("SELECT distinct CityCode,Area FROM airx.raw;")
	checkErr(err)
	defer rows.Close()

	if rows.Next() {
		// 这里输出第一行数据
		var (
			code int
			city string
		)
		err := rows.Scan(&code, &city)
		checkErr(err)

		data := CityData{code, city}
		datas = append(datas, data)

		for rows.Next() {
			err := rows.Scan(&code, &city)
			checkErr(err)

			data := CityData{code, city}
			datas = append(datas, data)
		}
	} else {
		errMsg := jsonErr{Code: 404, Text: "暂时没有此时段的城市表，请稍后重试"}
		datas = append(datas, errMsg)
	}
	return
}

func StationTable(location string) (datas AQIDatas) {
	db := Conn()
	defer db.Close()

	// 这里可以有个排序，按照城市代码
	rows, err := db.Query("SELECT distinct StationCode,PositionName,Longitude,Latitude FROM airx.raw where Area=?;", location)
	checkErr(err)
	defer rows.Close()

	if rows.Next() {
		// 这里输出第一行数据
		var (
			stationcode, positionname string
			longitude, latitude       float64
		)
		err := rows.Scan(&stationcode, &positionname, &longitude, &latitude)
		checkErr(err)

		data := StationData{stationcode, positionname, longitude, latitude}
		datas = append(datas, data)

		for rows.Next() {
			err := rows.Scan(&stationcode, &positionname, &longitude, &latitude)
			checkErr(err)

			data := StationData{stationcode, positionname, longitude, latitude}
			datas = append(datas, data)
		}
	} else {
		errMsg := jsonErr{Code: 404, Text: "暂时没有" + location + "的监测站信息，请稍后重试"}
		datas = append(datas, errMsg)
	}
	return
}
