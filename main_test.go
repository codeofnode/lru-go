package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	var cache, _ = New(10)
	var tests = []struct {
		a    string
		want bool
	}{
    // add
		{"00", false},
		{"01", false},
		{"02", false},
		{"03", false},
		{"04", false},
		{"05", false},
		{"06", false},
		{"07", false},
		{"08", false},
		{"09", false},
    // update 
		{"00", false},
    // add with eviction
		{"10", true},
	}

	var qtests = []struct {
		a    string
		ret  string
		want bool
	}{
    // fetch
		{"00", "00", true},
    // not present
		{"01", "", false},
    // fetch others
		{"02", "02", true},
		{"03", "03", true},
		{"04", "04", true},
		{"05", "05", true},
		{"06", "06", true},
		{"07", "07", true},
		{"08", "08", true},
		{"09", "09", true},
		{"10", "10", true},
	}
	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := cache.Add(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
	for _, tt := range qtests {

		testname := fmt.Sprintf("%v", tt.want)
		t.Run(testname, func(t *testing.T) {
			ans, ok := cache.Query(tt.a)
			if ok != tt.want {
				t.Errorf("key mismatch, got %v, want %v", ok, tt.want)
			}
			if ans != tt.ret {
				t.Errorf("value mismatch, got %s, want %s", ans, tt.ret)
			}
		})
	}
}
