package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// --- Part Two ---
// Of course, that would be the message - if you hadn't agreed to use a
// modified repetition code instead.
//
// In this modified code, the sender instead transmits what looks like
// random data, but for each character, the character they actually
// want to send is slightly less likely than the others. Even after
// signal-jamming noise, you can look at the letter distributions in
// each column and choose the least common letter to reconstruct the
// original message.
//
// In the above example, the least common character in the first column
// is a; in the second, d, and so on. Repeating this process for the
// remaining characters produces the original message, advent.
//
// Given the recording in your puzzle input and this new decoding
// methodology, what is the original message that Santa is trying to send?
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner)
	message := findTheMessage(input)
	fmt.Printf("Message: %v \n", message)
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

func parseInput(fileScanner *bufio.Scanner) [][]string {
	arr := [][]string{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		arrOfCharacters := parseLine(currLine)
		arr = append(arr, arrOfCharacters)
	}
	return arr
}

func parseLine(line string) []string {
	currArr := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		currArr = append(currArr, string(line[i]))
	}
	return currArr
}

func findTheMessage(input [][]string) string {
	message := ""
	m, n := len(input), len(input[0])
	for c := 0; c < n; c++ {
		mp := make(map[string]int)
		for r := 0; r < m; r++ {
			mp[string(input[r][c])]++
		}
		commonCharacter := findTheLeastCommonChar(mp)
		message += commonCharacter
	}
	return message
}

func findTheLeastCommonChar(mp map[string]int) string {
	common, freq := "", math.MaxInt
	for char, n := range mp {
		if n < freq {
			freq = n
			common = char
		}
	}
	return common
}
