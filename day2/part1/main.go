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

// Determine which games would have
// been possible if the bag had been loaded with only 12 red cubes,
// 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?

var configurationMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	idsOfPossibleGames := findPossibleGameIds()
	fmt.Printf("Possible Games: %v\n", idsOfPossibleGames)
	totalSumOfIds := findTotalSumOfIds(idsOfPossibleGames)
	fmt.Println(totalSumOfIds)
}

func findPossibleGameIds() []int {
	ids := []int{}
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lineNum := 1

	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		fmt.Printf("Curr Line: %v\n", currLine)
		isCurrLineValid := checkCurrLineValidity(currLine)
		if isCurrLineValid {
			ids = append(ids, lineNum)
		}
		lineNum++
	}
	return ids
}

func checkCurrLineValidity(currLine string) bool {
	n := len(currLine)
	currMap := make(map[string]int)
	idxOfFirstNumOfALine := getIdxOfFirstNumOfALine(currLine, n)

	for i := idxOfFirstNumOfALine; i < n; i++ {
		currNum, k := getCurrNumAndIdxBeforeCurrColor(currLine, i)
		for j := k + 1; j < n; j++ {
			if string(currLine[j]) == "," {
				currColor := currLine[k+1 : j]
				currMap[currColor] = currNum
				i = j + 1
				break // break through the loop and analyse other colors of the same set
			}
			if string(currLine[j]) == ";" {
				currColor := currLine[k+1 : j]
				currMap[currColor] = currNum
				i = j + 1
				isSetValid := checkSetValidity(currMap)
				if !isSetValid {
					return false
				}
				clear(currMap)
				break // break through the loop and analyse other sets
			}
			if j == n-1 {
				currColor := currLine[k+1 : j+1]
				currMap[currColor] = currNum
				i = j + 1
				isSetValid := checkSetValidity(currMap)
				if !isSetValid {
					return false
				}
				clear(currMap)
				break // break through the loop and analyse other sets
			}
		}
	}
	return true
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

func checkSetValidity(currMap map[string]int) bool {
	for k, v := range currMap {
		if v > configurationMap[k] {
			fmt.Println("Set is not valid")
			return false
		}
	}
	return true
}

func findTotalSumOfIds(ids []int) int {
	curr := 0
	for _, i := range ids {
		curr += i
	}
	return curr
}
