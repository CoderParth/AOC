package main

import (
	"fmt"
	"log"
	"strconv"
)

// --- Day 4: Secure Container ---
// You arrive at the Venus fuel depot only to discover it's protected by
// a password. The Elves had written the password on a sticky note,
// but someone threw it out.
//
// However, they do remember a few key facts about the password:
//
// It is a six-digit number.
// The value is within the range given in your puzzle input.
// Two adjacent digits are the same (like 22 in 122345).
// Going from left to right, the digits never decrease; they only
// ever increase or stay the same (like 111123 or 135679).
// Other than the range rule, the following are true:
//
// 111111 meets these criteria (double 11, never decreases).
// 223450 does not meet these criteria (decreasing pair of digits 50).
// 123789 does not meet these criteria (no double).
// How many different passwords within the range given in your
// puzzle input meet these criteria?
//
// Your puzzle input is 273025-767253.
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
	hasDouble := false
	prevNum := convStrToInt(string(s[0]))
	for i := 1; i < n; i++ {
		currNum := convStrToInt(string(s[i]))
		if prevNum < currNum {
			return false
		}
		if prevNum == currNum {
			hasDouble = true
			continue
		}
	}
	return hasDouble
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
