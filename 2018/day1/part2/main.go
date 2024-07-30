package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// other example situations:
// // +1, -1 first reaches 0 twice.
// +3, +3, +4, -2, -4 first reaches 10 twice.
// -6, +3, +8, +5, -6 first reaches 5 twice.
// +7, +7, -2, -7, -4 first reaches 14 twice.
// What is the first frequency your device reaches twice?

// Start at 0
// Apply +1 → frequency is 1
// Apply -2 → frequency is -1
// Apply +3 → frequency is 2
// Apply +1 → frequency is 3
// Repeat from the beginning
// Apply +1 again → frequency is 4
// Apply -2 → frequency is 2 (which we have seen before)
// The first frequency reached twice is 2.
func main() {
	resultingFrequency := calcResultingFreq()
	fmt.Printf("Resulting Freq: %v \n", resultingFrequency)
}

func calcResultingFreq() int {
	mp := make(map[int]int)
	total := 0
	mp[0] = 0
	fileScanner := createFileScanner()
	for {
		for fileScanner.Scan() {
			currLine := fileScanner.Text()
			n := len(currLine)
			currNum, sign := findNumAndSign(currLine, n)
			if sign == "+" {
				total += currNum
			} else {
				total -= currNum
			}
			if _, ok := mp[total]; ok {
				return total
			}
			mp[total] = 0
		}
		fileScanner = createFileScanner()
	}
}

func findNumAndSign(currLine string, n int) (int, string) {
	sign := string(currLine[0])
	currNumInStr := ""
	for i := 1; i < n; i++ {
		currNumInStr += string(currLine[i])
	}
	currNum, err := strconv.Atoi(currNumInStr)
	if err != nil {
		log.Fatal(err)
	}
	return currNum, sign
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
