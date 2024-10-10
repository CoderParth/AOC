package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// Amidst the chaos, you notice that exactly one claim doesn't overlap by
// even a single square inch of fabric with any other claim. If you can
// somehow draw attention to it, maybe the Elves will be able to make
// Santa's suit after all!
//
// For example, in the claims above, only claim 3 is intact after all claims are made.
//
// What is the ID of the only claim that doesn't overlap?

func main() {
	fileScanner := createFileScanner()
	fabric := initializeFabric()
	inputArr := [][]string{}
	for fileScanner.Scan() {
		currInputArr := parseInput(fileScanner.Text())
		inputArr = append(inputArr, currInputArr)
		claimTheArea(&fabric, currInputArr)
	}
	intactClaimId := findIntactClaim(fabric, inputArr)
	fmt.Printf("Intact Claim Id: %v \n", intactClaimId)
}

func findIntactClaim(fabric [][]string, inputArr [][]string) string {
	n := len(inputArr)
	for i := 0; i < n; i++ {
		currClaimId := inputArr[i][0]
		leftEdge, topEdge := convStrToInt(inputArr[i][1]), convStrToInt(inputArr[i][2])
		width, height := convStrToInt(inputArr[i][3]), convStrToInt(inputArr[i][4])
		isIntact := true
		for i := topEdge; i < topEdge+height; i++ {
			for j := leftEdge; j < leftEdge+width; j++ {
				if fabric[i][j] == "X" {
					isIntact = false
					break
				}
			}
		}
		if isIntact {
			return currClaimId
		}
	}
	return ""
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

func claimTheArea(fabric *[][]string, currInputArr []string) {
	leftEdge, topEdge := convStrToInt(currInputArr[1]), convStrToInt(currInputArr[2])
	width, height := convStrToInt(currInputArr[3]), convStrToInt(currInputArr[4])
	for i := topEdge; i < topEdge+height; i++ {
		for j := leftEdge; j < leftEdge+width; j++ {
			if (*fabric)[i][j] != "." {
				(*fabric)[i][j] = "X"
				continue
			}
			(*fabric)[i][j] = currInputArr[0]
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
