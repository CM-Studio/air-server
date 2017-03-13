package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {
	type args struct {
		inner http.Handler
		name  string
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Logger(tt.args.inner, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}
