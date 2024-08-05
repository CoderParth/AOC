package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Now, instead of considering the next digit, it wants you to
// consider the digit halfway around the circular list. That is,
// if your list contains 10 items, only include a digit in your
// sum if the digit 10/2 = 5 steps forward matches it. Fortunately,
// your list has an even number of elements.
//
// For example:
//
// 1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
// 1221 produces 0, because every comparison is between a 1 and a 2.
// 123425 produces 4, because both 2s match each other, but no other digit has a match.
// 123123 produces 12.
// 12131415 produces 4.
// What is the solution to your new captcha?

func main() {
	listOfMatchinDigits := findTheArrOfMatchingDigits()
	total := addAllNums(listOfMatchinDigits)
	fmt.Printf("total: %v \n", total)
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
	count, half := 0, n/2
	for i := 0; i < half; i++ {
		if string(currLine[i]) == string(currLine[i+half]) {
			currNum, err := strconv.Atoi(string(currLine[i]))
			if err != nil {
				log.Fatal(err)
			}
			count += currNum
		}
	}
	idx := 0
	for i := half; i < n; i++ {
		if string(currLine[i]) == string(currLine[idx]) {
			currNum, err := strconv.Atoi(string(currLine[i]))
			if err != nil {
				log.Fatal(err)
			}
			count += currNum
		}
		idx++
	}
	return count
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
