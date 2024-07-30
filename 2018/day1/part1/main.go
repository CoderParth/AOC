package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// other example situations:
//
// +1, +1, +1 results in  3
// +1, +1, -2 results in  0
// -1, -2, -3 results in -6
// Starting with a frequency of zero,
// what is the resulting frequency after
// all of the changes in frequency have been applied?

func main() {
	resultingFrequency := calcResultingFreq()
	fmt.Printf("Resulting Freq: %v \n", resultingFrequency)
}

func calcResultingFreq() int {
	total := 0
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		currNum, sign := findNumAndSign(currLine, n)
		if sign == "+" {
			total += currNum
		} else {
			total -= currNum
		}
	}
	return total
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
