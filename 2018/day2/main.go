package main

import (
	"bufio"
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
	input := readInputAndCollectStrings()
}

func readInputAndCollectStrings() []string {
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
