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
// The IDs abcde and axcye are close, but they differ by two characters (the second and fourth). However, the IDs fghij and fguij differ by exactly one character, the third (h and u). Those must be the correct boxes.
//
// What letters are common between the two correct box IDs? (In the example above, this is found by removing the differing character from either ID, producing fgij.)
func main() {
	totalTwos, totalThrees := readInputAndCollectStrings()
	checkSum := totalTwos * totalThrees
	fmt.Printf("check sum: %v \n", checkSum)
}

func readInputAndCollectStrings() (int, int) {
	fileScanner := createFileScanner()
	totalTwos, totalThrees := 0, 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		isTwoPresent, isThreePresent := findTwosAndThrees(currLine, n)
		if isTwoPresent {
			totalTwos++
		}
		if isThreePresent {
			totalThrees++
		}
	}
	return totalTwos, totalThrees
}

func findTwosAndThrees(currLine string, n int) (bool, bool) {
	mp := make(map[string]int)
	twos, threes := false, false
	for i := 0; i < n; i++ {
		currChar := string(currLine[i])
		mp[currChar]++
	}

	for i := 0; i < n; i++ {
		currChar := string(currLine[i])
		if !twos && mp[currChar] == 2 {
			twos = true
		}
		if !threes && mp[currChar] == 3 {
			threes = true
		}
	}
	return twos, threes
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
