package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	fmt.Printf("input: %v \n", input)
	totalTime := 1000
	flyingInfo := calculateFlyingRecords(input, totalTime)
	fmt.Printf("Comet info: %v \n", flyingInfo["Comet"])
	fmt.Printf("Dancer info: %v \n", flyingInfo["Dancer"])
}

// calculate the distance
// keep a record of distance, restPeriod, isResting, elapsedRestTime
//

type Attributes struct {
	speed         int
	flyingTime    int
	restingPeriod int
}

type FlyingRecord struct {
	distanceTravelled int
	flightTime        int
	isResting         bool
	elapsedRestTime   int
}

func calculateFlyingRecords(input map[string]Attributes, totalTime int) map[string]*FlyingRecord {
	fr := initializeFlyingRecords(input)
	for i := 0; i < totalTime; i++ {
		for bird, attribute := range input {
			fmt.Printf("Bird: %v\n", bird)
			fmt.Printf("Attribute: %v\n", attribute)
			if fr[bird].flightTime == attribute.flyingTime+1 {
				fr[bird].flightTime = 0
				fr[bird].isResting = true
			}

			if !fr[bird].isResting {
				fr[bird].distanceTravelled += attribute.speed
				fr[bird].flightTime++
			}

			if fr[bird].isResting {
				fr[bird].elapsedRestTime++
			}

			if fr[bird].elapsedRestTime == attribute.restingPeriod {
				fr[bird].flightTime = 0
				fr[bird].isResting = false
			}
		}
	}
	return fr
}

func initializeFlyingRecords(input map[string]Attributes) map[string]*FlyingRecord {
	mp := make(map[string]*FlyingRecord)
	for k := range input {
		initial := &FlyingRecord{
			distanceTravelled: 0,
			flightTime:        0,
			isResting:         false,
			elapsedRestTime:   0,
		}
		mp[k] = initial
	}
	return mp
}

// Parse input
func parseInput() map[string]Attributes {
	mp := make(map[string]Attributes)
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		arr := extractData(currLine, n)
		filterData(arr, mp)
	}
	return mp
}

func filterData(arr []string, mp map[string]Attributes) {
	name := arr[0]
	speed := convertStrToInt(arr[3])
	flyingTime := convertStrToInt(arr[6])
	restingPeriod := convertStrToInt(arr[13])
	a := Attributes{
		speed,
		flyingTime,
		restingPeriod,
	}
	mp[name] = a
}

func convertStrToInt(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func extractData(line string, n int) []string {
	arr := []string{}
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " || j == n-1 {
				arr = append(arr, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return arr
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
