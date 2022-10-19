package main

func FindLongestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		// single char check
		single := 0
		for i-single-1 >= 0 && i+single+1 < len(s) && s[i-single-1] == s[i+single+1] {
			single++
		}
		if 1+(single*2) > len(result) {
			result = s[i-single : i+single+1]
		}

		// double chars check
		if i+1 >= len(s) || s[i] != s[i+1] {
			continue
		}
		double := 0
		for i-double-1 >= 0 && i+double+2 < len(s) && s[i-double-1] == s[i+double+2] {
			double++
		}
		if 2+(double*2) > len(result) {
			result = s[i-double : i+double+2]
		}
	}
	return result
}
