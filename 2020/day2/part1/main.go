package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Day 2: Password Philosophy ---
// Your flight departs in a few days from the coastal airport; the easiest
// way down to the coast from here is via toboggan.
//
// The shopkeeper at the North Pole Toboggan Rental Shop is having a bad
// day. "Something's wrong with our computers; we can't log in!" You
// ask if you can take a look.
//
// Their password database seems to be a little corrupted: some of the
// passwords wouldn't have been allowed by the Official Toboggan Corporate
// Policy that was in effect when they were chosen.
//
// To try to debug the problem, they have created a list (your puzzle input)
// of passwords (according to the corrupted database) and the corporate
// policy when that password was set.
//
// For example, suppose you have the following list:
//
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
// Each line gives the password policy and then the password. The password
// policy indicates the lowest and highest number of times a given letter
// must appear for the password to be valid. For example, 1-3 a means that
// the password must contain a at least 1 time and at most 3 times.
//
// In the above example, 2 passwords are valid. The middle password, cdefg,
// is not; it contains no instances of b, but needs at least 1. The first
// and third passwords are valid: they contain one a or nine c, both within
// the limits of their respective policies.
//
// How many passwords are valid according to their policies?
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
	minNum, maxNum, givenLetter, password := convStrToInt(line[0]), convStrToInt(line[1]), line[2], line[4]
	passMap := make(map[string]int)
	n := len(password)
	passMap[givenLetter] = 0
	for i := 0; i < n; i++ {
		passMap[string(password[i])]++
	}
	for currChar, freq := range passMap {
		if currChar == givenLetter {
			if freq < minNum || freq > maxNum {
				return false
			}
			break
		}
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
