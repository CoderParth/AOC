package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		currMap := make(map[string]int)

		for i := 0; i < n; i++ {
			if unicode.IsSymbol(rune(currLine[i])) {
			}
		}

	}
}
