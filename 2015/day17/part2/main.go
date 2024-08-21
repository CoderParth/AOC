package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// --- Part Two ---
// While playing with all the containers in the kitchen,
// another load of eggnog arrives! The shipping and
// receiving department is requesting as many containers
// as you can spare.
//
// Find the minimum number of containers that can exactly
// fit all 150 liters of eggnog. How many different ways
// can you fill that number of containers and still hold
// exactly 150 litres?
//
// In the example above, the minimum number of containers
// was two. There were three ways to use that many
// containers, and so the answer there would be 3.
func main() {
	litersToStore := 25
	fileScanner := createFileScanner()
	input := parseInput(fileScanner) // create Array from input
	subsets := findSubsets(input)
	numOfMinCombinations := findNumOfMinCombinations(subsets, litersToStore)
	fmt.Printf("Num of combinations: %v \n", numOfMinCombinations)
}

func findNumOfMinCombinations(subsets [][]int, litersToStore int) int {
	mp := make(map[int]int) // num of containers, num of ways to use that many containers
	for i := 0; i < len(subsets); i++ {
		sum := calculateSum(subsets[i])
		if sum == litersToStore {
			n := len(subsets[i])
			mp[n]++
		}
	}
	fmt.Printf("mp: %v \n", mp)
	minComb := findMinimumKey(mp)
	return mp[minComb]
}

func findMinimumKey(mp map[int]int) int {
	lowest := math.MaxInt
	for k := range mp {
		if k < lowest {
			lowest = k
		}
	}
	return lowest
}

func calculateSum(nums []int) int {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	return totalSum
}

func findSubsets(input []int) [][]int {
	res := [][]int{}
	curr := []int{}
	backTrack(0, input, &curr, &res)
	return res
}

func backTrack(idx int, input []int, curr *[]int, res *[][]int) {
	copyOfCurr := make([]int, len(*curr))
	copy(copyOfCurr, (*curr))
	*res = append(*res, copyOfCurr)

	for i := idx; i < len(input); i++ {
		*curr = append(*curr, input[i])
		backTrack(i+1, input, curr, res)
		*curr = (*curr)[0 : len(*curr)-1]
	}
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
