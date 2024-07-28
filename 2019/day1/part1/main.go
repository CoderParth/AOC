package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module,
// take its mass, divide by three, round down, and subtract 2.
//
// For example:
//
// For a mass of 12, divide by 3 and round down to get 4,
// then subtract 2 to get 2.
// For a mass of 14, dividing by 3 and rounding down
// still yields 4, so the fuel required is also 2.
// For a mass of 1969, the fuel required is 654.
// For a mass of 100756, the fuel required is 33583.
//
// The Fuel Counter-Upper needs to know the total fuel requirement.
// To find it, individually calculate the fuel needed for the
// mass of each module (your puzzle input), then add together
// all the fuel values.
//
// What is the sum of the fuel requirements for all of the modules on your spacecraft?

func main() {
	arrOfFuels := getFuelsRequired()
	totalFuel := sumOfFuels(arrOfFuels)
	fmt.Printf("Total Fuel: %v \n", totalFuel)
}

func getFuelsRequired() []int {
	arr := []int{}
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		currFuel := getFuelFromALine(currLine, n)
		fuelRequiredForModule := calcFuelRequired(currFuel)
		arr = append(arr, fuelRequiredForModule)
	}
	return arr
}

func calcFuelRequired(n int) int {
}

func getFuelFromALine(currLine string, n int) int {
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

func sumOfFuels(arr []int) int {
	curr := 0
	for _, v := range arr {
		curr += v
	}
	return curr
}
