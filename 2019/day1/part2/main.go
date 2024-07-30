package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module,
// take its mass, divide by three, round down, and subtract 2.
//
// For example:
// A module of mass 14 requires 2 fuel. This fuel requires
// no further fuel (2 divided by 3 and rounded down is 0,
// which would call for a negative fuel), so the total fuel
// required is still just 2.
//
// At first, a module of mass 1969 requires 654 fuel. Then,
// this fuel requires 216 more fuel (654 / 3 - 2). 216 then
// requires 70 more fuel, which requires 21 fuel, which requires
// 5 fuel, which requires no further fuel. So, the total fuel
// required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.
// The fuel required by a module of mass 100756 and its
// fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.
//
//
// What is the sum of the fuel requirements for all of the
// modules on your spacecraft when also taking into account the
// mass of the added fuel? (Calculate the fuel requirements
//

func main() {
	arrOfFuels := getFuelsRequired()
	totalFuel := sumOfFuels(arrOfFuels)
	fmt.Printf("Total Fuel: %v \n", totalFuel)
}

func getFuelsRequired() []float64 {
	arr := []float64{}
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

func calcFuelRequired(n int) float64 {
	total, curr := float64(0), float64(0)
	for n > 0 {
		curr = math.Floor(float64(n/3) - 2)
		total += curr
	}

	return total
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

func sumOfFuels(arr []float64) int {
	curr := 0.0
	for _, v := range arr {
		curr += v
	}
	return int(curr)
}
