package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// For example
// abcde
// fghij
// klmno
// pqrst
// fguij
// axcye
// wvxyz
//
// The IDs abcde and axcye are close, but they differ
// by two characters (the second and fourth). However,
// the IDs fghij and fguij differ by exactly one character,
// the third (h and u). Those must be the correct boxes.
//
// What letters are common between the two correct box IDs?
// (In the example above, this is found by removing the
// differing character from either ID, producing fgij.)
func main() {
	input := readInputAndCollectStrings()
	commonChars := findCommonChars(input)
	fmt.Printf("common ids: %v \n", commonChars)
}

func readInputAndCollectStrings() []string {
	arr := []string{}
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		arr = append(arr, currLine)
	}
	return arr
}

func findCommonChars(input []string) string {
	m, n := len(input), len(input[0])
	for i := 0; i < m-1; i++ {
		for j := i + 1; j < n; j++ {
			firstString, secString := input[i], input[j]
			hasOnlyOneDiffChars, idx := findNumOfDiffCharAndIdx(firstString, secString, n)
			if hasOnlyOneDiffChars {
				s := ""
				s += firstString[0:idx]
				s += firstString[idx+1 : n]
				return s
			}
		}
	}
	return ""
}

func findNumOfDiffCharAndIdx(s1, s2 string, n int) (bool, int) {
	numOfDiff, idxOfDiff := 0, 0
	for i := 0; i < n; i++ {
		if numOfDiff > 1 {
			return false, 0
		}
		if string(s1[i]) != string(s2[i]) {
			idxOfDiff = i
			numOfDiff++
		}
	}
	return true, idxOfDiff
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
