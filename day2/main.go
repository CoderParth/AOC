package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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
	possibleGames := findPossibleGameIds()
	fmt.Println(possibleGames)
}

func findPossibleGameIds() []int {
	ids := []int{}
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)

		currGameId := string(currLine[5])

		currMap := make(map[string]int)
		foundSet := false
		for i := 8; i < n; i++ {
			if unicode.IsNumber(rune(currLine[i])) {
				currNum, err := strconv.Atoi(string(currLine[i]))
				if err != nil {
					log.Fatal(err)
				}
				for j := i + 2; j < n; j++ {
					if string(currLine[j]) == "," {
						currColor := currLine[i+2 : j]
						currMap[currColor] = currNum
						i = j + 2
						break
					}
					if string(currLine[j]) == ";" {
						currColor := currLine[i+2 : j]
						currMap[currColor] = currNum
						i = j + 2
						foundSet = true
						break
					}
				}

			}
		}

		i, err := strconv.Atoi(currGameId)
		if err != nil {
			log.Fatal(err)
		}
		ids = append(ids, i)

		// for i := 0; i < n; i++ {
		// 	if rune(currLine[i]) == ':' {
		// 		currLineNum := string(currLine[i-1])
		// 		i, err := strconv.Atoi(currLineNum)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 		ids = append(ids, i)
		// 		break
		// 	}
		// }

	}

	return ids
}
