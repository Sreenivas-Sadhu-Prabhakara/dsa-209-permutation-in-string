package main

import "fmt"

// CheckInclusion reports whether s2 contains any permutation of s1 as a
// contiguous substring — i.e. some window of s2 has exactly s1's letter
// counts. A fixed-size window slides, tracking how many of the 26 letter
// counts currently match.
func CheckInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var need, have [26]int
	for i := 0; i < n; i++ {
		need[s1[i]-'a']++
		have[s2[i]-'a']++
	}
	matches := 0
	for c := 0; c < 26; c++ {
		if need[c] == have[c] {
			matches++
		}
	}
	for right := n; right < m; right++ {
		if matches == 26 {
			return true
		}
		// Letter entering the window.
		in := s2[right] - 'a'
		have[in]++
		switch {
		case have[in] == need[in]:
			matches++
		case have[in] == need[in]+1:
			matches--
		}
		// Letter leaving the window.
		out := s2[right-n] - 'a'
		have[out]--
		switch {
		case have[out] == need[out]:
			matches++
		case have[out] == need[out]-1:
			matches--
		}
	}
	return matches == 26
}

func main() {
	fmt.Println(CheckInclusion("ab", "eidbaooo"))
	fmt.Println(CheckInclusion("ab", "eidboaoo"))
}
