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
	fmt.Printf("possible games: %v\n", idsOfPossibleGames)
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
	lineNum := 0

	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		fmt.Printf("curr line: %v\n", currLine)
		lineNum++
		n := len(currLine)

		currGameId := lineNum
		currMap := make(map[string]int)
		foundSet := false
		isCurrLineValid := true

		fmt.Printf("value of n: %v\n", n)
		fmt.Printf("curr Game ID : %v\n", currGameId)

		idxOfFirstNum := 0
		for i := 0; i < n; i++ {
			if string(currLine[i]) != ":" {
				continue
			}
			idxOfFirstNum = i + 2
		}

		for i := idxOfFirstNum; i < n; i++ {
			k := i + 1
			for string(currLine[k]) != " " {
				k++
			}

			fmt.Printf("curr unicode - i: %v\n", currLine[i:k])
			currNum, err := strconv.Atoi(currLine[i:k])
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Printf("curr unicode - i: %v\n", string(currLine[i]))
			// currNum, err := strconv.Atoi(string(currLine[i]))
			// if err != nil {
			// 	log.Fatal(err)
			// }
			for j := k + 1; j < n; j++ {
				if string(currLine[j]) == "," {
					fmt.Println("grabbing word")
					currColor := currLine[k+1 : j]
					currMap[currColor] = currNum
					fmt.Printf("Curr color: %v\n", currColor)
					i = j + 1

					fmt.Printf("curr map: %v\n", currMap)
					break
				}
				if string(currLine[j]) == ";" {
					fmt.Println("found ;")
					fmt.Println("grabbing word")
					currColor := currLine[k+1 : j]
					fmt.Printf("Curr color: %v\n", currColor)
					currMap[currColor] = currNum
					i = j + 1

					fmt.Printf("curr map: %v\n", currMap)
					foundSet = true
					break
				}
				if j == n-1 {
					fmt.Println("found ;")
					fmt.Println("grabbing word")
					currColor := currLine[k+1 : j+1]
					fmt.Printf("Curr color: %v\n", currColor)
					currMap[currColor] = currNum
					i = j + 1

					fmt.Printf("curr map: %v\n", currMap)
					foundSet = true
					break
				}

			}
			if foundSet {
				foundSet = false
				setValid := true
				fmt.Println("checking set validity")
				// check if the set is valid
				fmt.Printf("config map: %v\n", configurationMap)
				for k, v := range currMap {
					if v > configurationMap[k] {
						isCurrLineValid = false
						setValid = false
						fmt.Println("set not valid")
						fmt.Printf("value found: %v\n", v)
						break
					}
				}

				if setValid {
					clear(currMap)
					fmt.Println("set valid")
					fmt.Printf("cleared curr map: %v\n", currMap)
				}
			}

			if !isCurrLineValid {
				break
			}
		}

		if isCurrLineValid {
			ids = append(ids, lineNum)
			fmt.Println("Line valid")
			fmt.Println("found correct line")
			fmt.Println(ids)
		}

	}

	return ids
}

func findTotalSumOfIds(ids []int) int {
	curr := 0

	for _, i := range ids {
		curr += i
	}

	return curr
}
