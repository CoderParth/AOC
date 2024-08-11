package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// Realizing the error of his ways, Santa has switched to a better
// model of determining whether a string is naughty or nice. None of
// the old rules apply, as they are all clearly ridiculous.
//
// Now, a nice string is one with all of the following properties:
//
// It contains a pair of any two letters that appears at least twice
// in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa),
// but not like aaa (aa, but it overlaps).
// It contains at least one letter which repeats with exactly one letter
// between them, like xyx, abcdefeghi (efe), or even aaa.
// For example:
//
// qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj)
// and a letter that repeats with exactly one letter between them (zxz).
// xxyxx is nice because it has a pair that appears twice and a letter that
// repeats with one between, even though the letters used by each rule overlap.
// uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with
// a single letter between them.
// ieodomkazucvgmuy is naughty because it has a repeating letter with one
// between (odo), but no pair that appears twice.
// How many strings are nice under these new rules?
func main() {
	totalNiceStrings := findNiceStrings()
	fmt.Printf("Total Nice Strings: %v \n", totalNiceStrings)
}

func findNiceStrings() int {
	total := 0
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		if isStringNice(currLine, n) {
			total++
		}
	}
	return total
}

func isStringNice(line string, n int) bool {
	containsPair := containsPairOfTwoLetters(line, n)      // without overlapping``
	containsOneLetter := containsAtleastOneLetter(line, n) // a letter that repeats with one letter between them
	if containsPair && containsOneLetter {
		return true
	}
	return false
}

func containsPairOfTwoLetters(line string, n int) bool {
	// aabcdefgaa (aa)
	for i := 0; i < n-1; i++ {
		curr := string(line[i]) + string(line[i+1])
		for j := i + 2; j < n-1; j++ {
			next := string(line[j]) + string(line[j+1])
			if curr == next {
				return true
			}
		}
	}
	return false
}

// contains at least one letter which repeats with exactly one
// letter between them, like xyx
func containsAtleastOneLetter(line string, n int) bool {
	for i := 1; i < n-2; i++ {
		if string(line[i]) == string(line[i+2]) {
			return true
		}
	}
	return false
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
