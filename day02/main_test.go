package main

import "testing"

func Test_hasSetOfN(t *testing.T) {
	tests := []struct {
		in     []byte
		n      int
		expect bool
	}{
		{[]byte("yya"), 2, true},
		{[]byte("abbbbcdefghijaklmmmno"), 2, true},
		{[]byte("abc"), 2, false},
		{[]byte("abcdefghijk"), 2, false},
		{[]byte("abbcddefghijk"), 2, true},
		{[]byte("aaabcdefff"), 3, true},
		{[]byte("aabcdeff"), 3, false},
		{[]byte("abcdefff"), 3, true},
	}
	for i, test := range tests {
		got := hasSetOfN(test.in, test.n)
		if got != test.expect {
			t.Errorf("#%d: expected %v, got %v", i, test.expect, got)
		}
	}
}

func Test_checksup(t *testing.T) {
	tests := []struct {
		in     []string
		expect int
	}{
		{
			[]string{
				"abcdef",
				"bababc",
				"abbcde",
				"abcccd",
				"aabcdd",
				"abcdee",
				"ababab",
			},
			12,
		},
	}
	for i, test := range tests {
		got := checksum(test.in)
		if got != test.expect {
			t.Errorf("#%d: expected %v, got %v", i, test.expect, got)
		}
	}
}

func Test_commonIDString(t *testing.T) {
	tests := []struct {
		in     []string
		expect string
	}{
		{
			[]string{
				"abcde",
				"fghij",
				"klmno",
				"pqrst",
				"fguij",
				"axcye",
				"wvxyz",
			},
			"fgij",
		},
	}
	for i, test := range tests {
		got, err := commonIDString(test.in)
		if err != nil {
			t.Errorf("#%d: err: %v", i, err)
		}
		if got != test.expect {
			t.Errorf("#%d: expected %v, got %v", i, test.expect, got)
		}
	}
}

func Test_onlyOneMismatch(t *testing.T) {
	tests := []struct {
		in     []string
		expect bool
	}{
		{[]string{"fghij", "fguij"}, true},
		{[]string{"fghij", "fhuij"}, false},
	}
	for i, test := range tests {
		got := onlyOneMismatch(test.in[0], test.in[1])
		if got != test.expect {
			t.Errorf("#%d: expected %v, got %v", i, test.expect, got)
		}
	}
}

func Test_commonString(t *testing.T) {
	tests := []struct {
		in     []string
		expect string
	}{
		{[]string{"fghij", "fguij"}, "fgij"},
	}
	for i, test := range tests {
		got := commonString(test.in[0], test.in[1])
		if got != test.expect {
			t.Errorf("#%d: expected %v, got %v", i, test.expect, got)
		}
	}
}
