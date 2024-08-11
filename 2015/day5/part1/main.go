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
	fileScanner := createFileScanner()
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
