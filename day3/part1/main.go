package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
//

// In this schematic, two numbers are not part numbers
// because they are not adjacent to a symbol: 114 (top right)
// and 58 (middle right). Every other number is adjacent to a
// symbol and so is a part number; their sum is 4361.
//
// Of course, the actual engine schematic is much larger.
// What is the sum of all of the part numbers in the engine schematic?

var mpOfSymbols = map[string]int{
	"~":  0,
	"!":  0,
	"@":  0,
	"#":  0,
	"$":  0,
	"%":  0,
	"^":  0,
	"&":  0,
	"*":  0,
	"(":  0,
	")":  0,
	"_":  0,
	"+":  0,
	"`":  0,
	"-":  0,
	"=":  0,
	"{":  0,
	"}":  0,
	"|":  0,
	"[":  0,
	"]":  0,
	"\\": 0,
	":":  0,
	"\"": 0,
	";":  0,
	"'":  0,
	",":  0,
	"<":  0,
	">":  0,
	"/":  0,
	"?":  0,
}

func main() {
	allPartNums := findPartNums()
	fmt.Printf("Part Nums: %v\n", allPartNums)
	totalSumOfPartNums := findTotalSum(allPartNums)
	fmt.Printf("Total sum of part nums: %v\n", totalSumOfPartNums)
}

func findPartNums() []int {
	fileScanner := createFileScanner()

	n := totalLenOfStrInALine(fileScanner)     // total len of string in each line
	totalLines := totalNumOfLines(fileScanner) // total number of lines  present in input
	matrix := createMatrix(totalLines, n)

	secondFileScanner := createFileScanner() // create filescanner again for same input file
	matrix = fillMatrix(secondFileScanner, &matrix, n)

	partNums := checkForPartNumsInMatrix(matrix)
	return partNums
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

func totalLenOfStrInALine(fileScanner *bufio.Scanner) int {
	n := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n = len(currLine)
		break
	}
	return n
}

func totalNumOfLines(fileScanner *bufio.Scanner) int {
	totalLines := 1
	for fileScanner.Scan() {
		totalLines++
	}
	return totalLines
}

func createMatrix(totalLines, n int) [][]string {
	matrix := make([][]string, totalLines)
	for i := 0; i < totalLines; i++ {
		matrix[i] = make([]string, n)
	}
	return matrix
}

func fillMatrix(secondFileScanner *bufio.Scanner, matrix *[][]string, n int) [][]string {
	currLineNum := 0
	for secondFileScanner.Scan() {
		currLine := secondFileScanner.Text()
		for i := 0; i < n; i++ {
			currCharacter := currLine[i]
			(*matrix)[currLineNum][i] = string(currCharacter)

		}
		currLineNum++
	}
	return *matrix
}

func checkForPartNumsInMatrix(matrix [][]string) []int {
	partNums := []int{}
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			fmt.Printf("curr matrix: %v \n", matrix[i])
			idxOfCurrNum := []int{}
			for k := j; k < n; k++ {
				if _, err := strconv.Atoi(matrix[i][k]); err != nil {
					j = k
					break
				}
				idxOfCurrNum = append(idxOfCurrNum, k)
			}

			if len(idxOfCurrNum) == 0 {
				continue
			}

			currNumInStr := ""
			for _, s := range idxOfCurrNum {
				valInTheIdx := string(matrix[i][s])
				fmt.Printf("val in the idx: %v \n", valInTheIdx)
				currNumInStr += valInTheIdx
			}

			fmt.Printf("curr num in str: %v \n ", currNumInStr)
			currNum, err := strconv.Atoi(currNumInStr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("curr num is: %v \n", currNum)

			// search around each index
			isPartNum := false
			for _, idx := range idxOfCurrNum {
				if i > 0 {
					// check if top, topleft diagonal, or topright diaognal has a symbol
					if _, ok := mpOfSymbols[matrix[i-1][idx]]; ok {
						isPartNum = true
						break
					}
					if idx > 0 {
						if _, ok := mpOfSymbols[matrix[i-1][idx-1]]; ok {
							isPartNum = true
							break
						}
					}
					if idx < n-1 {
						if _, ok := mpOfSymbols[matrix[i-1][idx+1]]; ok {
							isPartNum = true
							break
						}
					}
				}
				// check if a symbol is present either on the left or on the right
				if idx > 0 {
					if _, ok := mpOfSymbols[matrix[i][idx-1]]; ok {
						isPartNum = true
						break
					}
				}
				if idx < n-1 {
					if _, ok := mpOfSymbols[matrix[i][idx+1]]; ok {
						isPartNum = true
						break
					}
				}
				// check if a symbol is present at either bottomleft diagonal, bottomright diagonal
				// or at bottom
				if i < n-1 {
					if idx > 0 {
						if _, ok := mpOfSymbols[matrix[i+1][idx-1]]; ok {
							isPartNum = true
							break
						}
					}
					if idx < n-1 {
						if _, ok := mpOfSymbols[matrix[i+1][idx+1]]; ok {
							isPartNum = true
							break
						}
					}
					if _, ok := mpOfSymbols[matrix[i+1][idx]]; ok {
						isPartNum = true
						break
					}
				}
			}

			if isPartNum {
				partNums = append(partNums, currNum)
			}
		}
	}

	return partNums
}

func findTotalSum(partNums []int) int {
	curr := 0
	for _, p := range partNums {
		curr += p
	}
	return curr
}
