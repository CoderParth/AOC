package main

import (
	"fmt"
	"log"
	"strconv"
)

// --- Part Two ---
// An Elf just remembered one more important detail: the two adjacent
// matching digits are not part of a larger group of matching digits.
//
// Given this additional criterion, but still ignoring the range rule,
// the following are now true:
//
// 112233 meets these criteria because the digits never decrease and
// all repeated digits are exactly two digits long.
// 123444 no longer meets the criteria (the repeated 44 is part of a
// larger group of 444).
// 111122 meets the criteria (even though 1 is repeated more than twice,
// it still contains a double 22).
// How many different passwords within the range given in your puzzle
// input meet all of the criteria?

func main() {
	lowRange := 273025
	highRange := 767253
	total := findValidPasswords(lowRange, highRange) // find total number of different passwords
	fmt.Printf("Total: %v \n", total)
}

func findValidPasswords(lowRange, highRange int) int {
	total := 0
	for lowRange <= highRange {
		if meetsCriteria(lowRange) {
			total++
		}
		lowRange++
	}
	return total
}

func meetsCriteria(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	// prevNum := convStrToInt(string(s[0]))
	prevNum := s[0]

	// Check if the digits are decreasing
	for i := 1; i < n; i++ {
		currNum := s[i]
		if prevNum > currNum {
			return false // digit is decreasing
		}
		prevNum = currNum
	}

	if s[0] == s[1] && s[0] != s[2] {
		return true
	}

	if s[n-1] == s[n-2] && s[n-1] != s[n-3] {
		return true
	}

	for i := 1; i < n-2; i++ {
		if s[i] == s[i+1] && s[i] != s[i+2] && s[i] != s[i-1] {
			return true
		}
	}

	return false
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
