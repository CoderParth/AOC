package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// It sounds like the goal is to find the only two numbers in each row
// where one evenly divides the other - that is, where the result of
// the division operation is a whole number. They would like you to find
// those numbers on each line, divide them, and add up each line's result.
//
// For example, given the following spreadsheet:
//
// 5 9 2 8
// 9 4 7 3
// 3 8 6 5
//
// In the first row, the only two numbers that evenly divide are 8 and 2;
// the result of this division is 4.
// In the second row, the two numbers are 9 and 3; the result is 3.
// In the third row, the result is 2.
// In this example, the sum of the results would be 4 + 3 + 2 = 9.
//
// What is the sum of each row's result in your puzzle input?
func main() {
	listOfDifferences := findDifferences() // differences of the largest and the smalles values
	total := findTotal(listOfDifferences)
	fmt.Printf("Total: %v \n", total)
}

func findDifferences() []int {
	arr := []int{}
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		a := createArrOfNums(currLine, n)
		d := findTwoNumsThatEvenlyDivide(a) // Returns the divison of those numbers
		arr = append(arr, d)
	}
	return arr
}

func createArrOfNums(line string, n int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " && len(curr) > 0 {
				i = j
				currNum := convertStrToNum(curr)
				arr = append(arr, currNum)
				break
			}
			curr += string(line[j])
		}
	}
	return arr
}

func findTwoNumsThatEvenlyDivide(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := arr[i], arr[j]
			r1 := a % b
			r2 := b % a
			if r1 == 0 {
				return (a / b)
			}
			if r2 == 0 {
				return (b / a)
			}
		}
	}
	return 0
}

func convertStrToNum(curr string) int {
	currNum, err := strconv.Atoi(curr)
	if err != nil {
		log.Fatal(err)
	}
	return currNum
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

func findTotal(arr []int) int {
	curr := arr[0]
	for i := 1; i < len(arr); i++ {
		curr += arr[i]
	}
	return curr
}
