package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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

func main() {
	allPartNums := findPartNums()
	fmt.Printf("Part Nums: %v\n", allPartNums)
	totalSumOfPartNums := findTotalSum(allPartNums)
	fmt.Printf("Total sum of part nums: %v\n", totalSumOfPartNums)
}

func findPartNums() []int {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	n := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n = len(currLine)
		break
	}

	totalLines := 1
	for fileScanner.Scan() {
		totalLines++
	}
	fmt.Printf("total lines: %v \n", totalLines)

	matrix := make([][]rune, totalLines)
	for i := 0; i < totalLines; i++ {
		matrix[i] = make([]rune, n)
	}

	// create filescanner again
	readFileSecond, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScannerSecond := bufio.NewScanner(readFileSecond)
	fileScannerSecond.Split(bufio.ScanLines)

	currLineNum := 0
	for fileScannerSecond.Scan() {
		currLine := fileScannerSecond.Text()
		fmt.Printf("Curr Line: %v \n", currLine)

		for i := 0; i < n; i++ {
			currCharacter := currLine[i]
			if string(currCharacter) != " " {
				matrix[currLineNum] = append(matrix[currLineNum], rune(currCharacter))
			}
		}
		fmt.Printf("curr matrix: %v \n", string(matrix[currLineNum]))
		currLineNum++
	}

	partNums := checkForPartNumsInMatrix(matrix)
	return partNums
	// return []int{}
}

func checkForPartNumsInMatrix(matrix [][]rune) []int {
	partNums := []int{}
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := j

			fmt.Printf("curr matrix: %v \n", string(matrix[i]))
			fmt.Printf("matrix[i][k]: %v \n", string(matrix[i][k]))
			idxOfCurrNum := []int{}
			for ; k < n; k++ {
				if unicode.IsSymbol(matrix[i][k]) {
					break
				}
				idxOfCurrNum = append(idxOfCurrNum, k)
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
			// for _, idx := range idxOfCurrNum {
			//
			// }

			partNums = append(partNums, currNum)

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
