package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// The next year, to speed up the process, Santa creates a robot
// version of himself, Robo-Santa, to deliver presents with him.
//
// Santa and Robo-Santa start at the same location (delivering two
// presents to the same starting house), then take turns moving
// based on instructions from the elf, who is eggnoggedly reading
// from the same script as the previous year.
//
// This year, how many houses receive at least one present?
//
// For example:
//
// ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
// ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
// ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
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
	x1, y1, x2, y2 := 0, 0, 0, 0
	arr := [2]int{0, 0}
	mp[arr] = 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		for i := 0; i < n; i++ {
			currSign := currLine[i]
			if i%2 == 0 { // handle Santa's move
				if string(currSign) == "^" {
					x1 += 1
				} else if string(currSign) == "v" {
					x1 -= 1
				} else if string(currSign) == "<" {
					y1 -= 1
				} else {
					y1 += 1
				}
				newArr := [2]int{x1, y1}
				mp[newArr] = 0
			} else { // handle Robo-Santa's move
				if string(currSign) == "^" {
					x2 += 1
				} else if string(currSign) == "v" {
					x2 -= 1
				} else if string(currSign) == "<" {
					y2 -= 1
				} else {
					y2 += 1
				}
				newArr := [2]int{x2, y2}
				mp[newArr] = 0
			}
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
