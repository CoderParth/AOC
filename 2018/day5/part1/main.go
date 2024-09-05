package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// --- Day 5: Alchemical Reduction ---
// You've managed to sneak in to the prototype suit manufacturing lab.
// The Elves are making decent progress, but are still struggling with
// the suit's size reduction capabilities.
//
// While the very latest in 1518 alchemical technology might have solved
// their problem eventually, you can do better. You scan the chemical
// composition of the suit's material and discover that it is formed by
// extremely long polymers (one of which is available as your puzzle input).
//
// The polymer is formed by smaller units which, when triggered, react with
// each other such that two adjacent units of the same type and opposite
// polarity are destroyed. Units' types are represented by letters; units'
// polarity is represented by capitalization. For instance, r and R are units
// with the same type but opposite polarity, whereas r and s are entirely
// different types and do not react.
//
// For example:
//
// In aA, a and A react, leaving nothing behind.
// In abBA, bB destroys itself, leaving aA. As above, this then destroys
// itself, leaving nothing.
// In abAB, no two adjacent units are of the same type, and so nothing happens.
// In aabAAB, even though aa and AA are of the same type, their polarities
// match, and so nothing happens.
// Now, consider a larger example, dabAcCaCBAcCcaDA:
//
// dabAcCaCBAcCcaDA  The first 'cC' is removed.
// dabAaCBAcCcaDA    This creates 'Aa', which is removed.
// dabCBAcCcaDA      Either 'cC' or 'Cc' are removed (the result is the same).
// dabCBAcaDA        No further actions can be taken.
// After all possible reactions, the resulting polymer contains 10 units.
//
// How many units remain after fully reacting the polymer you scanned?
// (Note: in this puzzle and others, the input is large; if you
// copy/paste your input, make sure you get the whole thing.)
func main() {
	input := parseInput()
	remainingUnits := findRemainingUnits(input)
	fmt.Printf("Number of remaining units: %v \n", remainingUnits)
}

func parseInput() string {
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		return currLine
	}
	return ""
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

func findRemainingUnits(input string) int {
	fmt.Printf("Starting input: %v \n", input)
	for hasReactors(input) {
		input = breakReactors(input)
	}
	return len(input)
}

func hasReactors(input string) bool {
	n := len(input)
	for i := 0; i < n-1; i++ {
		if strings.ToUpper(string(input[i])) == string(input[i+1]) {
			return true
		}
		if strings.ToUpper(string(input[i+1])) == string(input[i]) {
			return true
		}
	}
	return false
}

func breakReactors(input string) string {
	n := len(input)
	i, j := findReactorIdxs(input, n)
	tmp := input
	input = tmp[0:i]
	input += tmp[j+1 : n]
	return input
}

func findReactorIdxs(input string, n int) (int, int) {
	for i := 0; i < n-1; i++ {
		if strings.ToUpper(string(input[i])) == string(input[i+1]) {
			return i, i + 1
		}
		if strings.ToUpper(string(input[i+1])) == string(input[i]) {
			return i, i + 1
		}
	}
	return 0, 0
}
