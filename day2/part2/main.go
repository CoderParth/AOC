package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Example Input:
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

// The power of a set of cubes is equal to the numbers of red, green, and blue
// cubes multiplied together. The power of the minimum set of cubes in game 1 is 48.
// In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five
// powers produces the sum 2286.
//
// For each game, find the minimum set of cubes that must have been present.
// What is the sum of the power of these sets?

func main() {
	listOfPowerPerLine := findPowerPerLine()
	fmt.Printf("Power Per line: %v\n", listOfPowerPerLine)
	totalSumOfPowers := findTotalSumOfPower(listOfPowerPerLine)
	fmt.Println(totalSumOfPowers)
}

func findPowerPerLine() []int {
	powers := []int{}
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		fmt.Printf("Curr Line: %v\n", currLine)
		powerOfThisLine := checkPowerForCurrLine(currLine)
		powers = append(powers, powerOfThisLine)
	}
	return powers
}

func checkPowerForCurrLine(currLine string) int {
	n := len(currLine)
	currMap := make(map[string]int)
	idxOfFirstNumOfALine := getIdxOfFirstNumOfALine(currLine, n)

	for i := idxOfFirstNumOfALine; i < n; i++ {
		currNum, k := getCurrNumAndIdxBeforeCurrColor(currLine, i)
		for j := k + 1; j < n; j++ {
			if string(currLine[j]) == "," {
				currColor := currLine[k+1 : j]
				if currNum > currMap[currColor] {
					currMap[currColor] = currNum
				}
				i = j + 1
				break // break through the loop and analyse other colors of the same set
			}
			if string(currLine[j]) == ";" {
				currColor := currLine[k+1 : j]
				if currNum > currMap[currColor] {
					currMap[currColor] = currNum
				}
				i = j + 1
				break // break through the loop and analyse other sets
			}
			if j == n-1 {
				currColor := currLine[k+1 : j+1]
				if currNum > currMap[currColor] {
					currMap[currColor] = currNum
				}
				i = j + 1
				break // break through the loop and analyse other lines
			}
		}
	}
	totalPower := calculatePowerOfALine(currMap)
	return totalPower
}

func calculatePowerOfALine(mp map[string]int) int {
	total := 1
	fmt.Println("calculating power of a line ")
	for k, v := range mp {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
		total *= v
	}
	return total
}

func getIdxOfFirstNumOfALine(line string, n int) int {
	for i := 0; i < n; i++ {
		if string(line[i]) != ":" {
			continue
		}
		return i + 2
	}
	return 0
}

func getCurrNumAndIdxBeforeCurrColor(currLine string, i int) (int, int) {
	k := i + 1
	for string(currLine[k]) != " " {
		k++
	}
	currNum, err := strconv.Atoi(currLine[i:k])
	if err != nil {
		log.Fatal(err)
	}
	return currNum, k
}

func findTotalSumOfPower(powers []int) int {
	curr := 0
	for _, p := range powers {
		curr += p
	}
	return curr
}
