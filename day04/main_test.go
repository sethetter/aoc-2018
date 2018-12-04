package main

import "testing"

func TestSomething(t *testing.T) {
	tests := []struct {
		in     string
		expect string
	}{
		{"yo", "yo"},
		{"sup", "sup"},
	}
	for i, test := range tests {
		got := test.in
		if got != test.expect {
			t.Errorf("#%d: expected %v, got %v", i, test.expect, got)
		}
	}
}
