package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// While it appears you validated the passwords correctly, they don't seem
// to be what the Official Toboggan Corporate Authentication System is
// expecting.
//
// The shopkeeper suddenly realizes that he just accidentally explained
// the password policy rules from his old job at the sled rental place
// down the street! The Official Toboggan Corporate Policy actually works
// a little differently.
//
// Each policy actually describes two positions in the password, where 1
// means the first character, 2 means the second character, and so on.
// (Be careful; Toboggan Corporate Policies have no concept of "index zero"!)
// Exactly one of these positions must contain the given letter. Other
// occurrences of the letter are irrelevant for the purposes of policy
// enforcement.
//
// Given the same example list from above:
//
// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
// 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
// How many passwords are valid according to the new interpretation of the
// policies?
func main() {
	fileScanner := createFileScanner()
	totalValids := 0
	for fileScanner.Scan() {
		line := parseLine(fileScanner.Text())
		if isValid(line) {
			totalValids++
		}
	}
	fmt.Printf("Total Valids: %v \n", totalValids)
}

func isValid(line []string) bool {
	firstIdx, secondIdx, givenLetter, password := convStrToInt(line[0]), convStrToInt(line[1]), line[2], line[4]
	if string(password[firstIdx-1]) == givenLetter && string(password[secondIdx-1]) == givenLetter {
		return false
	}
	if string(password[firstIdx-1]) != givenLetter && string(password[secondIdx-1]) != givenLetter {
		return false
	}
	return true
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
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

func parseLine(line string) []string {
	arr := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == "-" || string(line[j]) == " " || string(line[j]) == ":" {
				arr = append(arr, curr)
				i = j
				curr = ""
				continue
			}
			if j == n-1 {
				curr += string(line[j])
				arr = append(arr, curr)
				i = j
				curr = ""
				continue
			}
			curr += string(line[j])
		}
	}
	return arr
}
