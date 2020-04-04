package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_getUSD(t *testing.T) {

	tests := []struct {
		name string
		args []string
		want bool
	}{
		{"test dollar only", []string{"دولار امريكي", "الدولار", "دولار&nbsp;الامريكي<"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getUSD(tt.args)
			if got != tt.want {
				t.Errorf("getUSD() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ebs_extractEBS(t *testing.T) {
	type args struct {
		data io.Reader
	}

	f, _ := ioutil.ReadFile("index.html")
	data := bytes.NewReader(f)

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"testing reader", args{data: data}, []string{"data"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ebs{}
			if got := e.extractEBS(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ebs.extractEBS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ebs_extractEBSInternet(t *testing.T) {
	type args struct {
		data io.Reader
	}

	f2, err := http.Get("https://standforsudan.ebs-sd.com/StandForSudan/")
	if err != nil {
		t.Fatalf("Error in http.Get: %v", err)
	}

	defer f2.Body.Close()
	// internet, _ := ioutil.ReadAll(f2.Body)

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test from internet", args{data: f2.Body}, []string{"data1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ebs{}
			if got := e.extractEBS(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ebs.extractEBS() = %v, want %v", got, tt.want)
			}
		})
	}
}
