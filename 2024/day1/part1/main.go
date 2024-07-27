package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 1000
// 2000
// 3000
//
// 4000
//
// 5000
// 6000
//
// 7000
// 8000
// 9000
//
// 10000
//
// In case the Elves get hungry and need extra snacks,
// they need to know which Elf to ask: they'd like to
// know how many Calories are being carried by the Elf
// carrying the most Calories. In the example above,
// this is 24000 (carried by the fourth Elf).
//
// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

func main() {
	listOfCalories := findTheListOfCalories()
	highestCal := calculatehighestCal(listOfCalories)
	fmt.Printf("highest calorie: %v \n", highestCal)
}

func findTheListOfCalories() []int {
	fileScanner := createFileScanner()
	listOfCalories := []int{}
	totalCalsPerElf := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		if n == 0 {
			listOfCalories = append(listOfCalories, totalCalsPerElf)
			totalCalsPerElf = 0
			continue
		}
		calorie := findCalorieForThisLine(currLine, n)
		totalCalsPerElf += calorie
	}
	return listOfCalories
}

func findCalorieForThisLine(currLine string, n int) int {
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

func createFileScanner() *bufio.Scanner {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func calculatehighestCal(cals []int) int {
	curr := cals[0]
	for i := 1; i < len(cals); i++ {
		if cals[i] > curr {
			curr = cals[i]
		}
	}
	return curr
}
