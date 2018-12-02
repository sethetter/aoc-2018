package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	cs := checksum(lines)
	fmt.Printf("Part1: %d\n", cs)

	common, err := commonIDString(lines)
	if err != nil {
		fail(err, "error on part 2")
	}
	fmt.Printf("Part2: %s\n", common)
}

func fail(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(1)
}

func checksum(lines []string) int {
	pairs, triples := 0, 0
	for _, l := range lines {
		if hasSetOfN([]byte(l), 2) {
			pairs = pairs + 1
		}
		if hasSetOfN([]byte(l), 3) {
			triples = triples + 1
		}
	}
	return pairs * triples
}

func hasSetOfN(in []byte, n int) bool {
	counts := make(map[byte]int)
	for _, l := range in {
		counts[l] = counts[l] + 1
	}
	for _, c := range counts {
		if c == n {
			return true
		}
	}
	return false
}

func commonIDString(ids []string) (string, error) {
	idPair := make([]string, 2)
	sort.Strings(ids)
	for i, id1 := range ids {
		for _, id2 := range ids[i+1:] {
			if onlyOneMismatch(id1, id2) {
				idPair[0], idPair[1] = id1, id2
				break
			}
		}
		// Did we find our pair? Stop searching
		if idPair[0] != "" || idPair[1] != "" {
			break
		}
	}
	if idPair[0] == "" || idPair[1] == "" {
		return "", errors.New("Did not find ID pair")
	}
	return commonString(idPair[0], idPair[1]), nil
}

func onlyOneMismatch(id1, id2 string) bool {
	mismatches := 0
	for j, _ := range id1 {
		if id1[j] != id2[j] {
			mismatches = mismatches + 1
		}
		if mismatches > 1 {
			break
		}
	}
	return mismatches == 1
}

func commonString(s1, s2 string) string {
	out := ""
	for i, _ := range s1 {
		if s1[i] == s2[i] {
			out = out + string(s1[i])
		}
	}
	return out
}
