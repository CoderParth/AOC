package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// --- Day 17: No Such Thing as Too Much ---
// The elves bought too much eggnog again - 150 liters this time.
// To fit it all into your refrigerator, you'll need to move it
// into smaller containers. You take an inventory of the capacities
// of the available containers.
//
// For example, suppose you have containers of size 20, 15, 10, 5,
// and 5 liters. If you need to store 25 liters, there are four ways
// to do it:
//
// 15 and 10
// 20 and 5 (the first 5)
// 20 and 5 (the second 5)
// 15, 5, and 5
// Filling all containers entirely, how many different combinations
// of containers can exactly fit all 150 liters of eggnog?
func main() {
	litersToStore := 25
	fileScanner := createFileScanner()
	input := parseInput(fileScanner) // create Array from input
}

func parseInput(fileScanner *bufio.Scanner) []int {
	input := []int{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		numAsStr := getCurrNum(currLine)
		numAsInt := convStrToInt(numAsStr)
		input = append(input, numAsInt)
	}
	return input
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

func getCurrNum(line string) string {
	n := ""
	for i := 0; i < len(line); i++ {
		n += string(line[i])
	}
	return n
}

func convStrToInt(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
