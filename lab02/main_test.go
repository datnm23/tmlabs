package main

import "testing"

type testCase struct {
	input  string
	expect string
}

func TestReplaceAllSensitiveWords(t *testing.T) {
	testCases := []testCase{
		{
			input:  "skill",
			expect: "skill",
		},
		{
			input:  "   kill   ",
			expect: "   k*ll   ",
		},
		{
			input:  ".kill",
			expect: ".k*ll",
		},
		{
			input:  "kill!",
			expect: "k*ll!",
		},
		{
			input:  ".Kill!",
			expect: ".K*ll!",
		},
		{
			input:  "   sex   ",
			expect: "   s*x   ",
		},
		{
			input:  ",sex",
			expect: ",s*x",
		},
		{
			input:  "drug!",
			expect: "dr*g!",
		},
		{
			input:  ".fuck!",
			expect: ".f*ck!",
		},
	}

	for i, tc := range testCases {
		t.Logf("Testing case: %02d", i)
		result := replaceAllSensitiveWords(tc.input)
		if result != tc.expect {
			t.Errorf("case %d, want %s but got %s\n", i, tc.expect, result)
		}
	}
}
