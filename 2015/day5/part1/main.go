package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// A nice string is one with all of the following properties:
//
// It contains at least three vowels (aeiou only), like aei,
// xazegov, or aeiouaeiouaeiou.
//
// It contains at least one letter that appears twice in a row,
// like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
//
// It does not contain the strings ab, cd, pq, or xy, even if
// they are part of one of the other requirements.
// For example:
//
// ugknbfddgicrmopn is nice because it has at least three vowels
// (u...i...o...), a double letter (...dd...), and none of the
// disallowed substrings.
//
// aaa is nice because it has at least three vowels and a double
// letter, even though the letters used by different rules overlap.
//
// jchzalrnumimnmhp is naughty because it has no double letter.
// haegwjzuvuyypxyu is naughty because it contains the string xy.
// dvszwmarrgswjxmb is naughty because it contains only one vowel.
// How many strings are nice?
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
	containsThreeVowels := checkIfStrHasThreeVowels(line, n)
	letterAppearsTwice := checkIfLetterAppearsTwice(line, n)
	doesNotContain := doesNotContainSpecificStrings(line, n)
	if containsThreeVowels && letterAppearsTwice && doesNotContain {
		return true
	}
	return false
}

func doesNotContainSpecificStrings(line string, n int) bool {
	mp := map[string]int{
		"ab": 0,
		"cd": 0,
		"pq": 0,
		"xy": 0,
	}
	for i := 0; i < n-1; i++ {
		curr := ""
		curr += string(line[i]) + string(line[i+1])
		if _, ok := mp[curr]; ok {
			return false
		}
	}
	return true
}

func checkIfLetterAppearsTwice(line string, n int) bool {
	curr := string(line[0])
	for i := 1; i < n; i++ {
		if string(line[i]) == curr {
			return true
		}
		curr = string(line[i])
	}
	return false
}

func checkIfStrHasThreeVowels(line string, n int) bool {
	mp := map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}
	for i := 0; i < n; i++ {
		curr := string(line[i])
		if _, ok := mp[curr]; ok {
			mp[curr]++
		}
	}
	numOfVowels := 0
	for _, v := range mp {
		numOfVowels += v
	}
	return numOfVowels >= 3
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
