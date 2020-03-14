package main

import "testing"

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
