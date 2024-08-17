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

// A gear is any * symbol that is adjacent
// to exactly two part numbers. Its gear ratio
// is the result of multiplying those two numbers together.
//
//
//In this schematic, there are two gears. The first is in
//the top left; it has part numbers 467 and 35, so its gear ratio
//is 16345. The second gear is in the lower right; its gear ratio
// is 451490. (The * adjacent to 617 is not a gear because it
// is only adjacent to one part number.) Adding up all of the
// gear ratios produces 467835.

// What is the sum of all of the gear ratios in your engine schematic?

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

type IdxsOfSymbol struct {
	First  int
	Second int
}

func main() {
	allPartNums := findPartNums()
	totalSumOfGearRatios := findTotalSumOfGearRatios(allPartNums)
	fmt.Printf("Total sum of gear ratios: %v\n", totalSumOfGearRatios)
}

func findTotalSumOfGearRatios(allPartNums map[IdxsOfSymbol][]int) int {
	total := 0
	for _, arr := range allPartNums {
		currTotal := 1
		if len(arr) > 1 {
			for _, num := range arr {
				currTotal *= num
			}
			total += currTotal
		}
	}
	return total
}

func findPartNums() map[IdxsOfSymbol][]int {
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

func checkForPartNumsInMatrix(matrix [][]string) map[IdxsOfSymbol][]int {
	partNums := make(map[IdxsOfSymbol][]int)
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
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
			currNum := getCurrNum(idxOfCurrNum, &matrix, i)
			isPartNum, idxOfAdjSymbol := checkIfCurrNumIsPartNum(idxOfCurrNum, &matrix, i, n)
			if isPartNum {
				first, second := idxOfAdjSymbol[0], idxOfAdjSymbol[1]
				if matrix[first][second] == "*" {
					idxs := IdxsOfSymbol{first, second}
					partNums[idxs] = append(partNums[idxs], currNum)
				}
			}
		}
	}
	return partNums
}

// create a map where (i, j) is an index of where the symbol was found.
// the value will be the part num.
func checkIfCurrNumIsPartNum(idxOfCurrNum []int, matrix *[][]string, i, n int) (bool, []int) {
	for _, idx := range idxOfCurrNum {
		// check if top, topleft diagonal, or topright diaognal has a symbol
		if ok, idxs := hasSymbolAtTopOrTopDiagonals(matrix, i, idx, n); ok {
			return true, idxs
		}

		// check if a symbol is present either on the left or on the right
		if ok, idxs := hasSymbolAtEitherLeftOrRight(matrix, i, idx, n); ok {
			return true, idxs
		}

		// check if a symbol is present at either bottomleft diagonal, bottomright diagonal
		// or at the bottom
		if ok, idxs := hasSymbolAtBottomOrBottomDiagonals(matrix, i, idx, n); ok {
			return true, idxs
		}
	}
	return false, []int{}
}

func hasSymbolAtTopOrTopDiagonals(matrix *[][]string, i, idx, n int) (bool, []int) {
	if i > 0 {
		if _, ok := mpOfSymbols[(*matrix)[i-1][idx]]; ok {
			return true, []int{i - 1, idx}
		}
		if idx > 0 {
			if _, ok := mpOfSymbols[(*matrix)[i-1][idx-1]]; ok {
				return true, []int{i - 1, idx - 1}
			}
		}
		if idx < n-1 {
			if _, ok := mpOfSymbols[(*matrix)[i-1][idx+1]]; ok {
				return true, []int{i - 1, idx + 1}
			}
		}
	}
	return false, []int{}
}

func hasSymbolAtEitherLeftOrRight(matrix *[][]string, i, idx, n int) (bool, []int) {
	if idx > 0 {
		if _, ok := mpOfSymbols[(*matrix)[i][idx-1]]; ok {
			return true, []int{i, idx - 1}
		}
	}
	if idx < n-1 {
		if _, ok := mpOfSymbols[(*matrix)[i][idx+1]]; ok {
			return true, []int{i, idx + 1}
		}
	}
	return false, []int{}
}

func hasSymbolAtBottomOrBottomDiagonals(matrix *[][]string, i, idx, n int) (bool, []int) {
	if i < n-1 {
		if idx > 0 {
			if _, ok := mpOfSymbols[(*matrix)[i+1][idx-1]]; ok {
				return true, []int{i + 1, idx - 1}
			}
		}
		if idx < n-1 {
			if _, ok := mpOfSymbols[(*matrix)[i+1][idx+1]]; ok {
				return true, []int{i + 1, idx + 1}
			}
		}
		if _, ok := mpOfSymbols[(*matrix)[i+1][idx]]; ok {
			return true, []int{i + 1, idx}
		}
	}
	return false, []int{}
}

func getCurrNum(idxOfCurrNum []int, matrix *[][]string, i int) int {
	currNumInStr := ""
	for _, s := range idxOfCurrNum {
		valInTheIdx := string((*matrix)[i][s])
		currNumInStr += valInTheIdx
	}
	currNum, err := strconv.Atoi(currNumInStr)
	if err != nil {
		log.Fatal(err)
	}
	return currNum
}
