package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 1721
// 979
// 366
// 299
// 675
// 1456
//
// In this list, the two entries that sum to 2020
// are 1721 and 299. Multiplying them together
// produces 1721 * 299 = 514579, so the correct
// answer is 514579.
//
// Of course, your expense report is much larger.
// Find the two entries that sum to 2020;
// what do you get if you multiply them together?

// I guess, create a map, save the numbers.
// then solve it like a simple two sum problem.
func main() {
	mpOfNums := createMapFromInput()
	twoNums := findTwoNumsThatSumTo2020(mpOfNums)
	multipleOfTwoNums := multiplyTwoNums(twoNums)
	fmt.Printf("multiple of two nums: %v \n", multipleOfTwoNums)
}

func createMapFromInput() map[int]int {
	mp := make(map[int]int)
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		num := getNumFromALine(currLine, n)
		mp[num] = 0
	}
	return mp
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

func getNumFromALine(currLine string, n int) int {
	currNumInStr := ""
	for i := 0; i < n; i++ {
		currNumInStr += string(currLine[i])
	}
	currNum, err := strconv.Atoi(currNumInStr)
	if err != nil {
		log.Fatal(err)
	}
	return currNum
}

func findTwoNumsThatSumTo2020(mp map[int]int) []int {
	for k := range mp {
		comp := 2020 - k
		if _, ok := mp[comp]; ok {
			return []int{k, comp}
		}
	}
	return []int{}
}

func multiplyTwoNums(arr []int) int {
	return arr[0] * arr[1]
}
