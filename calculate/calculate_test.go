package calculate

import (
	"testing"
)

func TestResult(t *testing.T) {
	var tests = []struct {
		s 		string
		want 	string
	}{
		{`(4+(7-3)*5)-2`, "(4+(7-3)*5)-2=22"},
		{`(-1)-2`, "(-1)-2=invalid"},
		{`--1`, "--1=invalid"},
		{`-1`, "-1=invalid"},
		{`22`, "22=invalid"},
		{`(4+(3-7)*5)-2`, "(4+(3-7)*5)-2=-18"},
		{`1+2/3`, "1+2/3=1"},
		{`1+2*3`, "1+2*3=7"},
		{`-10`, "-10=invalid"},
	}
	for _, tt := range tests {
		t.Run("test calculate function", func(t *testing.T) {
			got := Result(tt.s)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

