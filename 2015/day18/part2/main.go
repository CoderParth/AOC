package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// You flip the instructions over; Santa goes on to point out
// that this is all just an implementation of Conway's Game of
// Life. At least, it was, until you notice that something's
// wrong with the grid of lights you bought: four lights,
// one in each corner, are stuck on and can't be turned off.
// The example above will actually run like this:
//
// Initial state:
// ##.#.#
// ...##.
// #....#
// ..#...
// #.#..#
// ####.#
//
// After 1 step:
// #.##.#
// ####.#
// ...##.
// ......
// #...#.
// #.####
//
// After 2 steps:
// #..#.#
// #....#
// .#.##.
// ...##.
// .#..##
// ##.###
//
// After 3 steps:
// #...##
// ####.#
// ..##.#
// ......
// ##....
// ####.#
//
// After 4 steps:
// #.####
// #....#
// ...#..
// .##...
// #.....
// #.#..#
//
// After 5 steps:
// ##.###
// .##..#
// .##...
// .##...
// #.#...
// ##...#
// After 5 steps, this example now has 17 lights on.
//
// In your grid of 100x100 lights, given your initial
// configuration, but with the four corners always
// in the on state, how many lights are on after 100 steps?
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner) // [][]string of input
	fmt.Printf("Input: %v \n", input)
	steps := 100
	for i := 0; i < steps; i++ {
		input = animateGrid(input)
	}
	fmt.Printf("Final Input: %v \n", input)
	totalOnLights := calculateNumOfOnLights(input)
	fmt.Printf("Total lights that are on: %v \n", totalOnLights)
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

func parseInput(fileScanner *bufio.Scanner) [][]string {
	inp := [][]string{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		currArr := createArrFrom(currLine)
		inp = append(inp, currArr)
	}
	return inp
}

func createArrFrom(line string) []string {
	arr := []string{}
	for i := 0; i < len(line); i++ {
		arr = append(arr, string(line[i]))
	}
	return arr
}

func calculateNumOfOnLights(inp [][]string) int {
	m, n := len(inp), len(inp[0])
	totalOnLights := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if inp[i][j] == "#" {
				totalOnLights++
			}
		}
	}
	return totalOnLights
}

func animateGrid(inp [][]string) [][]string {
	animatedGrid := make([][]string, len(inp))
	m, n := len(inp), len(inp[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			currLight := inp[i][j]
			on := false
			if currLight == "#" {
				on = true
			}
			numOfOnNeighbours := calculateNumOfOnNeighbours(inp, i, j)
			if on && (numOfOnNeighbours == 2 || numOfOnNeighbours == 3) {
				animatedGrid[i] = append(animatedGrid[i], "#")
				continue
			}
			if !on && (numOfOnNeighbours == 3) {
				animatedGrid[i] = append(animatedGrid[i], "#")
				continue
			}
			animatedGrid[i] = append(animatedGrid[i], ".")
		}
	}
	return animatedGrid
}

func calculateNumOfOnNeighbours(inp [][]string, i, j int) int {
	num, m, n := 0, len(inp), len(inp[0])
	if i > 0 {
		// calculate top left diagonal
		if j > 0 {
			topLeft := inp[i-1][j-1]
			if topLeft == "#" {
				num++
			}
		}

		// calculate top
		top := inp[i-1][j]
		if top == "#" {
			num++
		}

		// calculate top-right diagonal
		if j < n-1 {
			topRight := inp[i-1][j+1]
			if topRight == "#" {
				num++
			}
		}
	}

	// calculate left
	if j > 0 {
		left := inp[i][j-1]
		if left == "#" {
			num++
		}
	}

	// calculate right
	if j < n-1 {
		right := inp[i][j+1]
		if right == "#" {
			num++
		}
	}

	if i < m-1 {
		// calculate bottom left
		if j > 0 {
			bottomLeft := inp[i+1][j-1]
			if bottomLeft == "#" {
				num++
			}
		}

		// calculate bottom
		bottom := inp[i+1][j]
		if bottom == "#" {
			num++
		}

		// calculate bottom right
		if j < n-1 {
			bottomRight := inp[i+1][j+1]
			if bottomRight == "#" {
				num++
			}
		}
	}
	return num
}
