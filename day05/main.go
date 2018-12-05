package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var inFile *string = flag.String("input", "input.txt", "the input file")

func main() {
	flag.Parse()

	input, err := ioutil.ReadFile(*inFile)
	if err != nil {
		fail(err, "Couldn't read input file")
	}
	input = input[:len(input)-1]

	result := processInput(input)
	shortest := findShortest(result)
	fmt.Printf("Part1: %v\n", len(result))
	fmt.Printf("Part2: %d\n", shortest)
}

func fail(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(1)
}

func processInput(input []byte) []byte {
	for i := 0; i < len(input)-1; i++ {
		l1 := bytes.ToUpper([]byte{input[i]})
		l2 := bytes.ToUpper([]byte{input[i+1]})
		if l1[0] == l2[0] && input[i] != input[i+1] {
			new := append(input[:i], input[i+2:]...)
			return processInput(new)
		}
	}
	return input
}

func removeAll(input []byte, target byte) []byte {
	var result []byte
	for _, c := range input {
		if bytes.ToUpper([]byte{c})[0] != bytes.ToUpper([]byte{target})[0] {
			result = append(result, c)
		}
	}
	return result
}

func findShortest(input []byte) int {
	alphabet := []byte("abcdefghijklmnopqrstuvwxyz")
	shortest := len(input)
	for _, c := range alphabet {
		l := len(processInput(removeAll(input, c)))
		if l < shortest {
			shortest = l
		}
	}
	return shortest
}
