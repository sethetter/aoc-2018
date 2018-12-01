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

	changes := make([]change, len(lines)-1)
	for i, l := range lines {
		if l == "" {
			continue
		}
		change, err := parseLine(l)
		if err != nil {
			fail(err, "Error parsing number from input")
		}
		changes[i] = change
	}

	// Part 1
	var freq int
	for _, change := range changes {
		err := processChange(&freq, change)
		if err != nil {
			fail(err, "err")
		}
	}

	// Part 2
	firstRepeat, err := findFirstRepeat(changes)
	if err != nil {
		fail(err, "err")
	}

	fmt.Printf("Part 1: %v\n", freq)
	fmt.Printf("Part 2: %v\n", firstRepeat)
}

type change struct {
	op  byte
	val int
}

func parseLine(in string) (change, error) {
	val, err := strconv.Atoi(in[1:])
	if err != nil {
		return change{}, err
	}
	return change{op: in[0], val: val}, nil
}

func fail(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(1)
}

func processChange(f *int, c change) error {
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
		err := processChange(&freq, changes[i])
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
