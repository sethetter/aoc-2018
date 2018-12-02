package main

import (
	"testing"
)

func TestParseChange(t *testing.T) {
	test := []struct {
		in       string
		expected change
	}{
		{"+5", change{op: '+', val: 5}},
		{"+10", change{op: '+', val: 10}},
		{"-2", change{op: '-', val: 2}},
	}
	for i, test := range test {
		c, err := parseChange(test.in)
		if err != nil {
			t.Errorf("%d, failed: %v", i, err)
		}
		if c != test.expected {
			t.Errorf("%d, failed: expected %v, got %v", i, test.expected, c)
		}
	}
}

func TestProcessChange(t *testing.T) {
	test := []struct {
		c        change
		f        int
		expected int
	}{
		{change{op: '+', val: 5}, 0, 5},
		{change{op: '+', val: 3}, 1, 4},
		{change{op: '-', val: 3}, 1, -2},
	}
	for i, test := range test {
		err := applyChange(&test.f, test.c)
		if err != nil {
			t.Errorf("%d, failed: %v", i, err)
		}
		if test.f != test.expected {
			t.Errorf("%d, failed: expected %v, got %v", i, test.expected, test.f)
		}
	}
}

func TestFindFirstRepeat(t *testing.T) {
	tests := []struct {
		changes  []change
		expected int
	}{
		{[]change{change{'+', 5}, change{'-', 5}}, 0},
		{[]change{change{'+', 5}, change{'-', 10}, change{'+', 10}}, 5},
	}

	for i, test := range tests {
		got, err := findFirstRepeat(test.changes)
		if err != nil {
			t.Errorf("%d, failed: %v", i, err)
		}
		if got != test.expected {
			t.Errorf("%d, failed: expected %v, got %v", i, test.expected, got)
		}
	}
}
