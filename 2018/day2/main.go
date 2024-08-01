package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// For example, if you see the following box IDs:
//
// abcdef contains no letters that appear exactly two or three times.
// bababc contains two a and three b, so it counts for both.
// abbcde contains two b, but no letter appears exactly three times.
// abcccd contains three c, but no letter appears exactly two times.
// aabcdd contains two a and two d, but it only counts once.
// abcdee contains two e.
// ababab contains three a and three b, but it only counts once.
// Of these box IDs, four of them contain a letter which
// appears exactly twice, and three of them contain a letter
// which appears exactly three times. Multiplying these together
// produces a checksum of 4 * 3 = 12.
//
// What is the checksum for your list of box IDs?
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
		fmt.Printf("curr line: %v \n", currLine)
		isTwoPresent, isThreePresent := findTwosAndThrees(currLine, n)
		fmt.Printf("twos: %v \n", isTwoPresent)
		fmt.Printf("threes: %v \n", isThreePresent)
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
