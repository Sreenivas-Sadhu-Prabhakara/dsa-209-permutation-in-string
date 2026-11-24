package main

import "testing"

func TestCheckInclusion(t *testing.T) {
	cases := []struct {
		s1, s2 string
		want   bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"a", "a", true},
		{"abc", "ab", false}, // pattern longer than text
		{"adc", "dcda", true},
		{"hello", "ooolleoooleh", false},
		{"aa", "baa", true}, // match at the very end
		{"ab", "ab", true},
	}
	for _, c := range cases {
		if got := CheckInclusion(c.s1, c.s2); got != c.want {
			t.Errorf("CheckInclusion(%q,%q) = %v, want %v", c.s1, c.s2, got, c.want)
		}
	}
}
