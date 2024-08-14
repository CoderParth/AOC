package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Look-and-say sequences are generated iteratively,
// using the previous value as input for the next step.
// For each step, take the previous value, and replace
// each run of digits (like 111) with the number of digits
// (3) followed by the digit itself (1).
//
// For example:
//
// 1 becomes 11 (1 copy of digit 1).
// 11 becomes 21 (2 copies of digit 1).
// 21 becomes 1211 (one 2 followed by one 1).
// 1211 becomes 111221 (one 1, one 2, and two 1s).
// 111221 becomes 312211 (three 1s, two 2s, and one 1).
// Starting with the digits in your puzzle input, apply
// this process 40 times. What is the length of the result?
func main() {
	input := parseInput()
	for i := 0; i < 50; i++ {
		fmt.Printf("i: %v\n", i)
		sequence := createSequence(&input)
		input = sequence
	}
	fmt.Printf("Sequence: %v\n", input)
	fmt.Printf("Length of the result: %v \n", len(input))
}

func parseInput() string {
	fileScanner := createFileScanner()
	curr := ""
	for fileScanner.Scan() {
		curr = fileScanner.Text()
	}
	return curr
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

func createSequence(input *string) string {
	seq := ""
	n := len(*input)
	for i := 0; i < n; i++ {
		currNum := (*input)[i]
		numOfRepeats, newIdx := findNumOfRepeats(input, i, n)
		seq += string(numOfRepeats) + string(currNum)
		i = newIdx - 1
	}
	return seq
}

func findNumOfRepeats(input *string, idx, n int) (string, int) {
	numOfRepeats := 1
	i := idx + 1
	for ; i < n; i++ {
		if string((*input)[i]) != string((*input)[idx]) {
			break
		}
		numOfRepeats++
	}
	repeats := strconv.Itoa(numOfRepeats)
	return repeats, i
}
