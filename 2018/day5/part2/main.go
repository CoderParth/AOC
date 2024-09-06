package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// --- Part Two ---
// Time to improve the polymer.
//
// One of the unit types is causing problems; it's preventing the polymer
// from collapsing as much as it should. Your goal is to figure out which
// unit type is causing the most problems, remove all instances of it
// (regardless of polarity), fully react the remaining polymer, and measure its length.
//
// For example, again using the polymer dabAcCaCBAcCcaDA from above:
//
// Removing all A/a units produces dbcCCBcCcD. Fully reacting
// this polymer produces dbCBcD, which has length 6.
// Removing all B/b units produces daAcCaCAcCcaDA. Fully reacting
// this polymer produces daCAcaDA, which has length 8.
// Removing all C/c units produces dabAaBAaDA. Fully reacting this
// polymer produces daDA, which has length 4.
// Removing all D/d units produces abAcCaCBAcCcaA. Fully reacting
// this polymer produces abCBAc, which has length 6.
// In this example, removing all C/c units was best, producing the answer 4.
//
// What is the length of the shortest polymer you can produce by
// removing all units of exactly one type and fully reacting the result?
func main() {
	input := parseInput()
	unitTypesMap := initializeUnitTypes()
	shortestPolymer := findShortestPolymer(unitTypesMap, input)
	fmt.Printf("Shortest Polymer: %v \n", shortestPolymer)
}

func initializeUnitTypes() map[string]string {
	unitTypesMap := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
		"D": "d",
		"E": "e",
		"F": "f",
		"G": "g",
		"H": "h",
		"I": "i",
		"J": "j",
		"K": "k",
		"L": "l",
		"M": "m",
		"N": "n",
		"O": "o",
		"P": "p",
		"Q": "q",
		"R": "r",
		"S": "s",
		"T": "t",
		"U": "u",
		"V": "v",
		"W": "w",
		"X": "x",
		"Y": "y",
		"Z": "z",
	}
	return unitTypesMap
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

func findShortestPolymer(unitTypesMap map[string]string, input string) int {
	shortest := math.MaxInt
	for k, v := range unitTypesMap {
		newInput := removeThisPolymer(k, v, input)
		currUnitsLen := findRemainingUnits(newInput)
		fmt.Printf("Curr units length: %v\n", currUnitsLen)
		shortest = min(shortest, currUnitsLen)
	}
	return shortest
}

func removeThisPolymer(p1, p2, input string) string {
	newInput := ""
	n := len(input)
	for i := 0; i < n; i++ {
		if string(input[i]) == p1 || string(input[i]) == p2 {
			continue
		}
		newInput += string(input[i])
	}
	return newInput
}

func findRemainingUnits(input string) int {
	for {
		idx, ok := checkForReactors(input)
		if !ok {
			break
		}
		input = breakReactors(input, idx)
	}
	return len(input)
}

func checkForReactors(input string) (int, bool) {
	n := len(input)
	for i := 0; i < n-1; i++ {
		if string(input[i]) == string(input[i+1]) {
			continue
		}
		if strings.ToUpper(string(input[i])) == string(input[i+1]) {
			return i, true
		}
		if strings.ToUpper(string(input[i+1])) == string(input[i]) {
			return i, true
		}
	}
	return 0, false
}

func breakReactors(input string, i int) string {
	tmp := input
	n := len(tmp)
	input = tmp[0:i]
	input += tmp[i+2 : n]
	return input
}
