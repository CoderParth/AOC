package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// --- Part Two ---
// Seeing how reindeer move in bursts, Santa decides he's not pleased
// with the old scoring system.
//
// Instead, at the end of each second, he awards one point to the
// reindeer currently in the lead. (If there are multiple reindeer
// tied for the lead, they each get one point.) He keeps the
// traditional 2503 second time limit, of course, as doing otherwise
// would be entirely ridiculous.
//
// Given the example reindeer from above, after the first second,
// Dancer is in the lead and gets one point. He stays in the lead
// until several seconds into Comet's second burst: after the 140th
// second, Comet pulls into the lead and gets his first point. Of
// course, since Dancer had been in the lead for the 139 seconds
// before that, he has accumulated 139 points by the 140th second.
//
// After the 1000th second, Dancer has accumulated 689 points,
// while poor Comet, our old champion, only has 312. So, with the
// new scoring system, Dancer would win (if the race ended at 1000 seconds).
//
// Again given the descriptions of each reindeer (in your puzzle input),
// after exactly 2503 seconds, how many points does the winning reindeer have?
func main() {
	input := parseInput()
	totalTime := 2503
	flyingInfo := calculateFlyingRecords(input, totalTime)
	maxDistance := math.MinInt
	for k, v := range flyingInfo {
		fmt.Printf("Bird info: %v, %v \n", k, v)
		maxDistance = max(maxDistance, v.distanceTravelled)
	}
	fmt.Printf("Max distance: %v \n", maxDistance)
}

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
	for i := 1; i <= totalTime; i++ {
		for bird, attribute := range input {
			if fr[bird].flightTime == attribute.flyingTime {
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
				fr[bird].elapsedRestTime = 0
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
