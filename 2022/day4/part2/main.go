package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// It seems like there is still quite a bit of duplicate work planned.
// Instead, the Elves would like to know the number of pairs that
// overlap at all.
//
// In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't
// overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6,
// and 2-6,4-8) do overlap:
//
// 5-7,7-9 overlaps in a single section, 7.
// 2-8,3-7 overlaps all of the sections 3 through 7.
// 6-6,4-6 overlaps in a single section, 6.
// 2-6,4-8 overlaps in sections 4, 5, and 6.
// So, in this example, the number of overlapping assignment pairs is 4.
//
// In how many assignment pairs do the ranges overlap?
func main() {
	fileScanner := createFileScanner()
	totalValidPairs := 0 // assignment pairs where one range overlaps the other
	for fileScanner.Scan() {
		inputArr := parseLine(fileScanner.Text())
		if isValidPair(inputArr) { // if one range fully overlaps the other
			totalValidPairs++
		}
	}
	fmt.Printf("Total Valid Pairs: %v \n", totalValidPairs)
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

func isValidPair(inputArr []int) bool {
	firstSecMin, firstSecMax := inputArr[0], inputArr[1]
	secondSecMin, secondSecMax := inputArr[2], inputArr[3]
	if firstSecMin <= secondSecMin && firstSecMax >= secondSecMin {
		return true
	}
	if secondSecMin <= firstSecMin && secondSecMax >= firstSecMin {
		return true
	}
	return false
}

func parseLine(input string) []int {
	inputArr := []int{}
	n := len(input)
	for i := 0; i < n; i++ {
		curr := ""
		for j := i; j < n; j++ {
			if string(input[j]) == "-" || string(input[j]) == "," {
				num := convStrToInt(curr)
				inputArr = append(inputArr, num)
				i = j
				break
			}
			if j == n-1 {
				curr += string(input[j])
				num := convStrToInt(curr)
				inputArr = append(inputArr, num)
				i = j
				break
			}
			curr += string(input[j])
		}
	}
	return inputArr
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
