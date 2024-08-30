package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// Now, the jumps are even stranger: after each jump, if the offset
// was three or more, instead decrease it by 1. Otherwise, increase
// it by 1 as before.
//
// Using this rule with the above example, the process now takes 10 steps,
// and the offset values after finding the exit are left as 2 3 2 3 -1.
//
// How many steps does it now take to reach the exit?
func main() {
	fileScanner := createFileScanner()
	arr := createArrayFromInput(fileScanner)
	numOfSteps := findNumOfSteps(arr)
	fmt.Printf("Arr: %v \n", arr)
	fmt.Printf("Num Of Steps: %v \n", numOfSteps)
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

func createArrayFromInput(fileScanner *bufio.Scanner) []int {
	arr := []int{}
	for fileScanner.Scan() {
		currStr := fileScanner.Text()
		currNum := convStrToInt(currStr)
		arr = append(arr, currNum)
	}
	return arr
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func findNumOfSteps(arr []int) int {
	steps, currIdx, n := 0, 0, len(arr)
	for currIdx < n {
		tmpIdx := currIdx
		currIdx += arr[currIdx]
		if arr[tmpIdx] >= 3 {
			arr[tmpIdx]--
			steps++
			continue
		}
		arr[tmpIdx]++
		steps++
	}
	return steps
}
