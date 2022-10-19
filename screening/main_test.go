package main

import "testing"

func TestFindLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input string
		expect string
	}{
		{
			input: "babad",
			expect: "bab",
		},
		{
			input: "cbbd",
			expect: "bb",
		},
		{
			input: "abbacca",
			expect: "abba",
		},
	}
	for _, testCase := range testCases {
		actual := FindLongestPalindrome(testCase.input)
		if actual != testCase.expect {
			t.Errorf("expect '%s' to has palindrome '%s' but got '%s'", testCase.input, testCase.expect, actual)
		}
	}
}
