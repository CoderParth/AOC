package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// The Document indicates that you should start
// at the given coordinates (where you just landed)
// and face North. Then, follow the provided sequence:
// either turn left (L) or right (R) 90 degrees, then
// walk forward the given number of blocks, ending at
// a new intersection.  There's no time to follow such
// ridiculous instructions on foot, though, so you take
// a moment and work out the destination. Given that you
// can only walk on the street grid of the city, how far
// is the shortest path to the destination?
//
// For example:
// Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
// R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
// R5, L5, R5, R3 leaves you 12 blocks away.
// How many blocks away is Easter Bunny HQ?

func main() {
	// Moving "up" increases the y-coordinate.
	// Moving "down" decreases the y-coordinate.
	// Moving "right" increases the x-coordinate.
	// Moving "left" decreases the x-coordinate.
	distance := findTheShortesPath()
	fmt.Printf("Total block away: %v \n", distance)
}

func findTheShortesPath() int {
	fileScanner := createFileScanner()
	x, y := 0, 0
	currDir := "N"
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		dir := string(currLine[0])
		dist := ""
		for i := 1; i < n; i++ {
			dist += string(currLine[i])
		}
		currDist, err := strconv.Atoi(dist)
		if err != nil {
			log.Fatal(err)
		}
		checkConditionsAndEvaluate(&x, &y, &currDir, dir, currDist)
	}
	return abs(x) + abs(y)
}

func checkConditionsAndEvaluate(x, y *int, currDir *string, dir string, currDist int) {
	if (*currDir) == "N" {
		if dir == "L" {
			// move left - decrease x-coordinate
			(*x) -= currDist
			(*currDir) = "W"
		}
		if dir == "R" {
			// move Right - increase x-coordinate
			(*x) += currDist
			(*currDir) = "E"
		}
		return
	}

	if (*currDir) == "E" {
		if dir == "L" {
			// move upwards - increase y-coordinate
			(*y) += currDist
			(*currDir) = "N"
		}
		if dir == "R" {
			// move downwards - decrease y-coordinate
			(*y) -= currDist
			(*currDir) = "W"
		}
		return
	}

	if (*currDir) == "S" {
		if dir == "L" {
			// move Right - increase x-coordinate
			(*x) += currDist
			(*currDir) = "E"
		}
		if dir == "R" {
			// move left - decrease x-coordinate
			(*x) -= currDist
			(*currDir) = "W"
		}
		return
	}

	if (*currDir) == "W" {
		if dir == "L" {
			// move downwards - decrease y-coordinate
			(*y) -= currDist
			(*currDir) = "W"
		}
		if dir == "R" {
			// move upwards - increase y-coordinate
			(*y) += currDist
			(*currDir) = "N"
		}
		return
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
