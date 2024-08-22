package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// For example, in a simplified 6x6 grid, the light marked A
// has the neighbors numbered 1 through 8, and the light
// marked B, which is on an edge, only has the neighbors
// marked 1 through 5:
//
// 1B5...
// 234...
// ......
// ..123.
// ..8A4.
// ..765.
// The state a light should have next is based on its current
// state (on or off) plus the number of neighbors that are on:
//
// A light which is on stays on when 2 or 3 neighbors are on,
// and turns off otherwise.
// A light which is off turns on if exactly 3 neighbors are on,
// and stays off otherwise.
// All of the lights update simultaneously; they all consider
// the same current state before moving to the next.
//
// Here's a few steps from an example configuration of
// another 6x6 grid:
//
// Initial state:
// .#.#.#
// ...##.
// #....#
// ..#...
// #.#..#
// ####..
//
// After 1 step:
// ..##..
// ..##.#
// ...##.
// ......
// #.....
// #.##..
//
// After 2 steps:
// ..###.
// ......
// ..###.
// ......
// .#....
// .#....
//
// After 3 steps:
// ...#..
// ......
// ...#..
// ..##..
// ......
// ......
//
// After 4 steps:
// ......
// ......
// ..##..
// ..##..
// ......
// ......
// After 4 steps, this example has four lights on.
//
// In your grid of 100x100 lights, given your initial
// configuration, how many lights are on after 100 steps?
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner) // [][]string of input
	fmt.Printf("Input: %v \n", input)
	steps := 4
	for i := 0; i < steps; i++ {
		input = animateGrid(input)
	}
	fmt.Printf("Input: %v \n", input)
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
