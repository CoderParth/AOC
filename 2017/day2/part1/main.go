package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// The spreadsheet consists of rows of apparently-random numbers.
// To make sure the recovery process is on the right track, they
// need you to calculate the spreadsheet's checksum. For each row,
// determine the difference between the largest value and the smallest
// value; the checksum is the sum of all of these differences.
//
// For example, given the following spreadsheet:
//
// 5 1 9 5
// 7 5 3
// 2 4 6 8
// The first row's largest and smallest values are 9 and 1, and their difference is 8.
// The second row's largest and smallest values are 7 and 3, and their difference is 4.
// The third row's difference is 6.
// In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
//
// What is the checksum for the spreadsheet in your puzzle input?
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
		largest, smallest := findLargestAndSmallest(currLine, n)
		diff := largest - smallest
		arr = append(arr, diff)
	}
	return arr
}

func findLargestAndSmallest(line string, n int) (int, int) {
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
	largest, smallest := calculateMaxAndMin(arr)
	return largest, smallest
}

func calculateMaxAndMin(arr []int) (int, int) {
	l, s := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if l < arr[i] {
			l = arr[i]
			continue
		}
		if s > arr[i] {
			s = arr[i]
			continue
		}
	}
	return l, s
}

func convertStrToNum(curr string) int {
	currNum, err := strconv.Atoi(curr)
	if err != nil {
		log.Fatal(err)
	}
	return currNum
}

func createFileScanner() *bufio.Scanner {
	readFile, err := os.Open("in.txt")
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
