package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var inFile *string = flag.String("input", "input.txt", "the input file")

func main() {
	flag.Parse()

	input, err := ioutil.ReadFile(*inFile)
	if err != nil {
		fail(err, "Couldn't read input file")
	}

	lines := strings.Split(string(input), "\n")

	changes, err := parseChanges(lines)
	if err != nil {
		fail(err, "Error parsing lines")
	}

	// Part 1
	var freq int
	for _, change := range changes {
		err := applyChange(&freq, change)
		if err != nil {
			fail(err, "Error processing changes")
		}
	}

	// Part 2
	firstRepeat, err := findFirstRepeat(changes)
	if err != nil {
		fail(err, "Error finding first repeat")
	}

	fmt.Printf("Part 1: %v\n", freq)
	fmt.Printf("Part 2: %v\n", firstRepeat)
}

type change struct {
	op  byte
	val int
}

func parseChanges(lines []string) ([]change, error) {
	changes := make([]change, len(lines)-1)
	for i, l := range lines {
		change, err := parseChange(l)
		if err != nil {
			return nil, err
		}
		changes[i] = change
	}
	return changes, nil
}

func parseChange(in string) (change, error) {
	val, err := strconv.Atoi(in[1:])
	if err != nil {
		return change{}, err
	}
	return change{op: in[0], val: val}, nil
}

func applyChange(f *int, c change) error {
	switch c.op {
	case '+':
		*f = *f + c.val
		return nil
	case '-':
		*f = *f - c.val
		return nil
	default:
		return errors.New("Invalid operation")
	}
}

func findFirstRepeat(changes []change) (int, error) {
	var freq int
	var history []int
	var i int

	for {
		if i >= len(changes) {
			i = 0
		}
		history = append(history, freq)
		err := applyChange(&freq, changes[i])
		if err != nil {
			return 0, err
		}
		for _, h := range history {
			if h == freq {
				return freq, nil
			}
		}
		i++
	}
}

func fail(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(1)
}
