package main

import (
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
	lines = lines[0 : len(lines)-1] // Get rid of blank line

	claims := make([]claim, len(lines))
	for i, l := range lines {
		c, err := parseClaim(l)
		if err != nil {
			fail(err, "Error parsing claims")
		}
		claims[i] = c
	}

	grid := populateGrid(claims)
	overlaps := countOverlappingSquares(grid)

	fmt.Printf("Part1: %d\n", overlaps)
	fmt.Printf("Part2: %s\n", nonOverlappingClaim(grid, claims))
}

func fail(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(1)
}

type claim struct {
	id string
	x  int
	y  int
	w  int
	h  int
}

func parseClaim(in string) (claim, error) {
	parts := strings.Split(in, " @ ")
	id := parts[0]

	parts = strings.Split(parts[1], ": ")

	xy := strings.Split(parts[0], ",")
	wh := strings.Split(parts[1], "x")

	x, err := strconv.Atoi(xy[0])
	if err != nil {
		return claim{}, err
	}

	y, err := strconv.Atoi(xy[1])
	if err != nil {
		return claim{}, err
	}

	w, err := strconv.Atoi(wh[0])
	if err != nil {
		return claim{}, err
	}

	h, err := strconv.Atoi(wh[1])
	if err != nil {
		return claim{}, err
	}

	return claim{id, x, y, w, h}, nil
}

func populateGrid(claims []claim) [1000][1000][]string {
	var grid [1000][1000][]string
	// grid := make([][][]string, 1000, 1000)
	for _, c := range claims {
		for x := c.x; x < c.x+c.w; x++ {
			if x > 999 {
				break
			}
			for y := c.y; y < c.y+c.h; y++ {
				if y > 999 {
					break
				}
				grid[x][y] = append(grid[x][y], c.id)
			}
		}
	}
	return grid
}

func countOverlappingSquares(grid [1000][1000][]string) int {
	count := 0
	for x, _ := range grid {
		for y, _ := range grid[x] {
			if len(grid[x][y]) > 1 {
				count = count + 1
			}
		}
	}
	return count
}

func nonOverlappingClaim(grid [1000][1000][]string, claims []claim) string {
	overlaps := make(map[string]bool)
	for x, _ := range grid {
		for y, _ := range grid[x] {
			if len(grid[x][y]) > 1 {
				for _, id := range grid[x][y] {
					overlaps[id] = true
				}
			}
		}
	}
	for _, c := range claims {
		if _, ok := overlaps[c.id]; !ok {
			return c.id
		}
	}
	return ""
}
