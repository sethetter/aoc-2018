package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

	records := make([]record, len(lines))
	for i, l := range lines {
		rec := parseLine(l)
		records[i] = rec
	}

	sortRecords(records)
	sched := buildSchedule(records)

	mostMinutesIDStr := mostMinutesAsleep(sched)
	mostCommonMinute := mostCommonMinute(sched[mostMinutesIDStr])
	mostMinutesID, err := strconv.Atoi(mostMinutesIDStr)
	if err != nil {
		fail(err, "Error parsing mostMinutesID to int")
	}
	fmt.Printf("Part1: %v\n", mostMinutesID*mostCommonMinute)

	highestMinuteIDStr, highestMinute := mostConsistentSleepingMinute(sched)
	highestMinuteID, err := strconv.Atoi(highestMinuteIDStr)
	if err != nil {
		fail(err, "Error parsing highestMinuteID to int")
	}
	fmt.Printf("Part2: %v\n", highestMinute*highestMinuteID)
}

func fail(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(1)
}

type schedule [][]bool
type schedules map[string]schedule

type record struct {
	t    time.Time
	text string
}

func parseLine(line string) record {
	tstring := line[1:17]
	rest := line[19:]
	t, err := time.Parse("2006-01-02 15:04", tstring)
	if err != nil {
		fail(err, "")
	}
	return record{t, rest}
}

func sortRecords(recs []record) {
	sort.Slice(recs, func(i, j int) bool {
		return recs[j].t.After(recs[i].t)
	})
}

func buildSchedule(recs []record) schedules {
	sched := make(schedules)

	guard := ""
	asleep := 0
	day := 0

	for _, r := range recs {
		switch {
		case r.text[0:5] == "Guard":
			guard = r.text[7 : strings.Index(r.text[7:], " begins")+7]
			day = len(sched[guard])
			sched[guard] = append(sched[guard], make([]bool, 60))
			break
		case r.text == "falls asleep":
			asleep = r.t.Minute()
		case r.text == "wakes up":
			for i := asleep; i < r.t.Minute(); i++ {
				sched[guard][day][i] = true
			}
		}
	}

	return sched
}

func mostMinutesAsleep(sched schedules) string {
	counts := make(map[string]int, len(sched))
	highest := 0
	highestID := ""
	for g, days := range sched {
		for _, d := range days {
			for _, m := range d {
				if m {
					counts[g]++
					if counts[g] > highest {
						highest = counts[g]
						highestID = g
					}
				}
			}
		}
	}
	return highestID
}

func mostCommonMinute(sched schedule) int {
	counts := make(map[int]int)
	highest := 0
	highestI := 0
	for _, day := range sched {
		for i, asleep := range day {
			if asleep {
				counts[i]++
				if counts[i] > highest {
					highest = counts[i]
					highestI = i
				}
			}
		}
	}
	return highestI
}

func mostConsistentSleepingMinute(scheds schedules) (string, int) {
	results := make(map[string][]int, len(scheds))
	highestMinute := 0
	highestMinuteCount := 0
	highestMinuteID := ""
	for g, sched := range scheds {
		results[g] = make([]int, 60)
		for _, day := range sched {
			for m, asleep := range day {
				if asleep {
					results[g][m]++
					if results[g][m] > highestMinuteCount {
						highestMinuteCount = results[g][m]
						highestMinute = m
						highestMinuteID = g
					}
				}
			}
		}
	}
	return highestMinuteID, highestMinute
}
