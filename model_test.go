package main

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOneCityLatestData(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name      string
		args      args
		wantDatas AQIDatas
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatas := OneCityLatestData(tt.args.location); !reflect.DeepEqual(gotDatas, tt.wantDatas) {
				t.Errorf("OneCityLatestData() = %v, want %v", gotDatas, tt.wantDatas)
			}
		})
	}
}

func TestOneCitySingleData(t *testing.T) {
	type args struct {
		t        string
		location string
	}
	tests := []struct {
		name      string
		args      args
		wantDatas AQIDatas
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatas := OneCitySingleData(tt.args.t, tt.args.location); !reflect.DeepEqual(gotDatas, tt.wantDatas) {
				t.Errorf("OneCitySingleData() = %v, want %v", gotDatas, tt.wantDatas)
			}
		})
	}
}

func TestOneCityAllDataToday(t *testing.T) {
	type args struct {
		t        string
		location string
	}
	tests := []struct {
		name      string
		args      args
		wantDatas AQIDatas
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatas := OneCityAllDataToday(tt.args.t, tt.args.location); !reflect.DeepEqual(gotDatas, tt.wantDatas) {
				t.Errorf("OneCityAllDataToday() = %v, want %v", gotDatas, tt.wantDatas)
			}
		})
	}
}

func TestCompareDataOfCities(t *testing.T) {
	type args struct {
		locations []string
	}
	tests := []struct {
		name      string
		args      args
		wantDatas AQIDatas
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatas := CompareDataOfCities(tt.args.locations); !reflect.DeepEqual(gotDatas, tt.wantDatas) {
				t.Errorf("CompareDataOfCities() = %v, want %v", gotDatas, tt.wantDatas)
			}
		})
	}
}

func TestTrendDataNow(t *testing.T) {
	type args struct {
		t        string
		location string
	}
	tests := []struct {
		name      string
		args      args
		wantDatas AQIDatas
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatas := TrendDataNow(tt.args.t, tt.args.location); !reflect.DeepEqual(gotDatas, tt.wantDatas) {
				t.Errorf("TrendDataNow() = %v, want %v", gotDatas, tt.wantDatas)
			}
		})
	}
}

func TestCityTable(t *testing.T) {
	tests := []struct {
		name      string
		wantDatas AQIDatas
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatas := CityTable(); !reflect.DeepEqual(gotDatas, tt.wantDatas) {
				t.Errorf("CityTable() = %v, want %v", gotDatas, tt.wantDatas)
			}
		})
	}
}

func TestStationTable(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name      string
		args      args
		wantDatas AQIDatas
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatas := StationTable(tt.args.location); !reflect.DeepEqual(gotDatas, tt.wantDatas) {
				t.Errorf("StationTable() = %v, want %v", gotDatas, tt.wantDatas)
			}
		})
	}
}
