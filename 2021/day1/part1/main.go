package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 199
// 200
// 208
// 210
// 200
// 207
// 240
// 269
// 260
// 263

// 199 (N/A - no previous measurement)
// 200 (increased)
// 208 (increased)
// 210 (increased)
// 200 (decreased)
// 207 (increased)
// 240 (increased)
// 269 (increased)
// 260 (decreased)
// 263 (increased)
// In this example, there are 7 measurements that are larger than the previous measurement.
//
// How many measurements are larger than the previous measurement?

func main() {
	arrOfMeasurements := createArrOfMeasurements()
	numOfMeasurementsLargerThanPrevious := findTotalLargerMeasurements(arrOfMeasurements)
	fmt.Printf("Number of measurements: %v \n", numOfMeasurementsLargerThanPrevious)
}

func createArrOfMeasurements() []int {
	fileScanner := createFileScanner()
	arr := []int{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		num := getNumFromCurrLine(currLine)
		arr = append(arr, num)
	}
	return arr
}

func getNumFromCurrLine(currLine string) int {
	n := len(currLine)
	currNumInStr := ""
	for i := 0; i < n; i++ {
		currNumInStr += string(currLine[i])
	}
	currNum, err := strconv.Atoi(currNumInStr)
	if err != nil {
		log.Fatal(err)
	}
	return currNum
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

func findTotalLargerMeasurements(arr []int) int {
	total, prev := 0, arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > prev {
			total++
		}
		prev = arr[i]
	}
	return total
}
