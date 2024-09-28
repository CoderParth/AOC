package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// Time to check the rest of the slopes - you need to minimize the
// probability of a sudden arboreal stop, after all.
//
// Determine the number of trees you would encounter if, for each of the
// following slopes, you start at the top-left corner and traverse the
// map all the way to the bottom:
//
// Right 1, down 1.
// Right 3, down 1. (This is the slope you already checked.)
// Right 5, down 1.
// Right 7, down 1.
// Right 1, down 2.
// In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s)
// respectively; multiplied together, these produce the answer 336.
//
// What do you get if you multiply together the number of trees
// encountered on each of the listed slopes?
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner)
	fmt.Printf("Input: %v \n", input)
	numOfTrees := followSlope(input)
	fmt.Printf("Num of trees: %v \n", numOfTrees)
}

func followSlope(input [][]string) int {
	slopes := []int{}
	slopes = append(slopes, rightOneDownOne(input))
	slopes = append(slopes, rightThreeDownOne(input))
	slopes = append(slopes, rightFiveDownOne(input))
	slopes = append(slopes, rightSevenDownOne(input))
	slopes = append(slopes, rightOneDownTwo(input))
	product := 1
	for _, val := range slopes {
		product *= val
	}
	return product
}

func rightOneDownTwo(input [][]string) int {
	m, n := len(input), len(input[0])
	numOfTrees := 0
	i, j := 0, 0
	for {
		j++
		i += 2
		if i >= m || j >= n {
			break
		}
		if input[i][j] == "#" {
			numOfTrees++
		}
	}
	return numOfTrees
}

func rightSevenDownOne(input [][]string) int {
	m, n := len(input), len(input[0])
	numOfTrees := 0
	i, j := 0, 0
	for {
		j += 7
		i++
		if i >= m || j >= n {
			break
		}
		if input[i][j] == "#" {
			numOfTrees++
		}
	}
	return numOfTrees
}

func rightFiveDownOne(input [][]string) int {
	m, n := len(input), len(input[0])
	numOfTrees := 0
	i, j := 0, 0
	for {
		j += 5
		i++
		if i >= m || j >= n {
			break
		}
		if input[i][j] == "#" {
			numOfTrees++
		}
	}
	return numOfTrees
}

func rightOneDownOne(input [][]string) int {
	m, n := len(input), len(input[0])
	numOfTrees := 0
	i, j := 0, 0
	for {
		j++
		i++
		if i >= m || j >= n {
			break
		}
		if input[i][j] == "#" {
			numOfTrees++
		}
	}
	return numOfTrees
}

func rightThreeDownOne(input [][]string) int {
	m, n := len(input), len(input[0])
	numOfTrees := 0
	i, j := 0, 0
	for {
		j += 3
		i += 1
		if i >= m || j >= n {
			break
		}
		if input[i][j] == "#" {
			numOfTrees++
		}
	}
	return numOfTrees
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
	arr := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		currArr := parseLine(line)
		arr = append(arr, currArr)
	}
	return arr
}

func parseLine(line string) []string {
	n := len(line)
	currArr := []string{}
	for i := 0; i < 80; i++ {
		for j := 0; j < n; j++ {
			currArr = append(currArr, string(line[j]))
		}
	}
	return currArr
}
