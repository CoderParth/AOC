package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Day 3: No Matter How You Slice It ---
// The Elves managed to locate the chimney-squeeze prototype fabric for
// Santa's suit (thanks to someone who helpfully wrote its box IDs on the
// wall of the warehouse in the middle of the night). Unfortunately,
// anomalies are still affecting them - nobody can even agree on how to
// cut the fabric.
//
// The whole piece of fabric they're working on is a very large square -
// at least 1000 inches on each side.
//
// Each Elf has made a claim about which area of fabric would be ideal for
// Santa's suit. All claims have an ID and consist of a single rectangle
// with edges parallel to the edges of the fabric. Each claim's rectangle
// is defined as follows:
//
// The number of inches between the left edge of the fabric and the left
// edge of the rectangle.
// The number of inches between the top edge of the fabric and the top
// edge of the rectangle.
// The width of the rectangle in inches.
// The height of the rectangle in inches.
//
// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle
// 3 inches from the left edge, 2 inches from the top edge, 5 inches wide,
// and 4 inches tall. Visually, it claims the square inches of fabric
// represented by # (and ignores the square inches of fabric represented by
// .) in the diagram below:
//
// ...........
// ...........
// ...#####...
// ...#####...
// ...#####...
// ...#####...
// ...........
// ...........
// ...........
// The problem is that many of the claims overlap, causing two or more
// claims to cover part of the same areas. For example, consider the
// following claims:
//
// #1 @ 1,3: 4x4
// #2 @ 3,1: 4x4
// #3 @ 5,5: 2x2
// Visually, these claim the following areas:
//
// ........
// ...2222.
// ...2222.
// .11XX22.
// .11XX22.
// .111133.
// .111133.
// ........
// The four square inches marked with X are claimed by both 1 and 2.
// (Claim 3, while adjacent to the others, does not overlap either of them.)
//
// If the Elves all proceed with their own plans, none of them will have
// enough fabric. How many square inches of fabric are within two or more claims?

func main() {
	fileScanner := createFileScanner()
	fabric := initializeFabric()
	for fileScanner.Scan() {
		inputArr := parseInput(fileScanner.Text())
		claimTheArea(&fabric, inputArr)
	}
	totalSquareInches := calculateSquareInches(fabric)
	fmt.Printf("Total square inches: %v \n", totalSquareInches)
}

func calculateSquareInches(fabric [][]string) int {
	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] == "X" {
				total++
			}
		}
	}
	return total
}

func claimTheArea(fabric *[][]string, inputArr []string) {
	leftEdge, topEdge := convStrToInt(inputArr[1]), convStrToInt(inputArr[2])
	width, height := convStrToInt(inputArr[3]), convStrToInt(inputArr[4])
	for i := topEdge; i < topEdge+height; i++ {
		for j := leftEdge; j < leftEdge+width; j++ {
			if (*fabric)[i][j] != "." {
				(*fabric)[i][j] = "X"
				continue
			}
			(*fabric)[i][j] = inputArr[0]
		}
	}
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseInput(line string) []string {
	arr := []string{}
	n := len(line)
	i := 1

	// Parse claim Id
	claimId := ""
	for {
		if string(line[i]) == " " {
			i += 3
			break
		}
		claimId += string(line[i])
		i++
	}
	arr = append(arr, claimId)

	// parse left edge
	leftEdge := ""
	for {
		if string(line[i]) == "," {
			i++
			break
		}
		leftEdge += string(line[i])
		i++
	}
	arr = append(arr, leftEdge)

	// parse top edge
	topEdge := ""
	for {
		if string(line[i]) == ":" {
			i += 2
			break
		}
		topEdge += string(line[i])
		i++
	}
	arr = append(arr, topEdge)

	// parse width
	width := ""
	for {
		if string(line[i]) == "x" {
			i++
			break
		}
		width += string(line[i])
		i++
	}
	arr = append(arr, width)

	// parse height
	height := ""
	for {
		if i == n-1 {
			height += string(line[i])
			break
		}
		height += string(line[i])
		i++
	}
	arr = append(arr, height)
	return arr
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

func initializeFabric() [][]string {
	fabric := make([][]string, 1000)
	for i := 0; i < 1000; i++ {
		fabric[i] = make([]string, 1000)
		for j := 0; j < 1000; j++ {
			fabric[i][j] = "."
		}
	}
	return fabric
}
