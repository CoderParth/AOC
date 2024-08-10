package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// He begins by delivering a present to the house at
// his starting location, and then an elf at the North
// Pole calls him via radio and tells him where to move
// next. Moves are always exactly one house to the
// north (^), south (v), east (>), or west (<). After
// each move, he delivers another present to the house
// at his new location.
//
// However, the elf back at the north pole has had
// a little too much eggnog, and so his directions are
// a little off, and Santa ends up visiting some houses
// more than once. How many houses receive at least one present?
//
// For example:
// > delivers presents to 2 houses: one at the starting location,
// and one to the east.
//
// ^>v< delivers presents to 4 houses in a square, including twice
// to the house at his starting/ending location.
//
// ^v^v^v^v^v delivers a bunch of presents to some very lucky
// children at only 2 houses

func main() {
	numOfHouses := findNumOfUniqueHouses()
	fmt.Printf("Unique houses that receive presents: %v \n", numOfHouses)
}

func findNumOfUniqueHouses() int {
	fileScanner := createFileScanner()
	cordsOfDeliveredHouses := findDeliveredHouses(fileScanner)
	return len(cordsOfDeliveredHouses)
}

func findDeliveredHouses(fileScanner *bufio.Scanner) map[[2]int]int {
	mp := make(map[[2]int]int)
	x, y := 0, 0
	arr := [2]int{x, y}
	mp[arr] = 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		for i := 0; i < n; i++ {
			currSign := currLine[i]
			if string(currSign) == "^" {
				x += 1
			} else if string(currSign) == "v" {
				x -= 1
			} else if string(currSign) == "<" {
				y -= 1
			} else {
				y += 1
			}
			newArr := [2]int{x, y}
			mp[newArr] = 0
		}
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
