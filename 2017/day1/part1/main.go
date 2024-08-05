package main

import (
	"bufio"
	"log"
	"os"
)

// The captcha requires you to review a sequence of digits
// (your puzzle input) and find the sum of all digits that
// match the next digit in the list. The list is circular,
// so the digit after the last digit is the first digit in the list.
//
// For example:
//
// 1122 produces a sum of 3 (1 + 2) because the first digit (1)
// matches the second digit and the third digit (2) matches
// the fourth digit.
//
// 1111 produces 4 because each digit (all 1) matches the next.
// 1234 produces 0 because no digit matches the next.
// 91212129 produces 9 because the only digit that matches
// the next one is the last digit, 9.
//
// What is the solution to your captcha?

func main() {
	listOfMatchinDigits := findTheArrOfMatchingDigits()
	total := addAllNums(listOfMatchinDigits)
}

func findTheArrOfMatchingDigits() []int {
	fileScanner := createFileScanner()
	arr := []int{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		sumForCurrLine := calcSumForCurrLine(currLine, n)
		arr = append(arr, sumForCurrLine)
	}
	return arr
}

func calcSumForCurrLine(currLine string, n int) int {
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

func addAllNums(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}
