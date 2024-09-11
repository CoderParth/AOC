package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Day 6: Signals and Noise ---
// Something is jamming your communications with Santa. Fortunately,
// your signal is only partially jammed, and protocol in situations
// like this is to switch to a simple repetition code to get the
// message through.
//
// In this model, the same message is sent repeatedly. You've recorded
// the repeating message signal (your puzzle input), but the data seems
// quite corrupted - almost too badly to recover. Almost.
//
// All you need to do is figure out which character is most frequent
// for each position. For example, suppose you had recorded the
// following messages:
//
// eedadn
// drvtee
// eandsr
// raavrd
// atevrs
// tsrnev
// sdttsa
// rasrtv
// nssdts
// ntnada
// svetve
// tesnvt
// vntsnd
// vrdear
// dvrsen
// enarar
// The most common character in the first column is e; in the second, a;
// in the third, s, and so on. Combining these characters returns the
// error-corrected message, easter.
//
// Given the recording in your puzzle input, what is the error-corrected
// version of the message being sent?
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
		commonCharacter := findTheCommonChar(mp)
		message += commonCharacter
	}
	return message
}

func findTheCommonChar(mp map[string]int) string {
	common, freq := "", 0
	for char, n := range mp {
		if n > freq {
			freq = n
			common = char
		}
	}
	return common
}
