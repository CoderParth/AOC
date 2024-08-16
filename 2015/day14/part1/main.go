package main

import (
	"bufio"
	"log"
	"os"
)

// Reindeer can only either be flying (always at their top speed) or resting
// (not moving at all), and always spend whole seconds in either state.
//
// For example, suppose you have the following Reindeer:
//
// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
// Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
// After one second, Comet has gone 14 km, while Dancer has gone 16 km.
// After ten seconds, Comet has gone 140 km, while Dancer has gone 160 km.
// On the eleventh second, Comet begins resting (staying at 140 km), and
// Dancer continues on for a total distance of 176 km. On the 12th second,
// both reindeer are resting. They continue to rest until the 138th second,
// when Comet flies for another ten seconds. On the 174th second, Dancer
// flies for another 11 seconds.
//
// In this example, after the 1000th second, both reindeer are resting, and
// Comet is in the lead at 1120 km (poor Dancer has only gotten 1056 km by
// that point). So, in this situation, Comet would win (if the race ended
// at 1000 seconds).
//
// Given the descriptions of each reindeer (in your puzzle input), after
// exactly 2503 seconds, what distance has the winning reindeer traveled?
func main() {
	input := parseInput()
}

// Parse input
func parseInput() {
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		arr := extractData(currLine, n)
	}
}

func extractData(line string, n int) []string {
	arr := []string{}

	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		for j := i; j < n; j++ {
		}
	}
}

func createFileScanner() *bufio.Scanner {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}
