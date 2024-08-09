package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Now, given the same instructions, find the
// position of the first character that causes
// him to enter the basement (floor -1). The
// first character in the instructions has position 1,
// the second character has position 2, and so on.
//
// For example:
//
// ) causes him to enter the basement at character position 1.
// ()()) causes him to enter the basement at character position 5.
// What is the position of the character that causes Santa to first enter the basement?
func main() {
	charPos := findCharPos()
	fmt.Printf("Character position which takes santa to the basement: %v \n", charPos)
}

func findCharPos() int {
	fileScanner := createFileScanner()
	floor := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		floor = calculateParanthesis(currLine, n)
	}
	return floor
}

func calculateParanthesis(line string, n int) int {
	curr := 0
	for i := 0; i < n; i++ {
		if curr == -1 {
			return i
		}
		if string(line[i]) == "(" {
			curr++
		} else if string(line[i]) == ")" {
			curr--
		}
	}
	return n
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
