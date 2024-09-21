package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Day 4: Giant Squid ---
// You're already almost 1.5km (almost a mile) below the surface of the ocean,
// already so deep that you can't see any sunlight. What you can see, however,
// is a giant squid that has attached itself to the outside of your submarine.
//
// Maybe it wants to play bingo?
//
// Bingo is played on a set of boards each consisting of a 5x5 grid of numbers.
// Numbers are chosen at random, and the chosen number is marked on all boards
// on which it appears. (Numbers may not appear on all boards.) If all numbers
// in any row or any column of a board are marked, that board wins. (Diagonals
// don't count.)
//
// The submarine has a bingo subsystem to help passengers (currently, you and
// the giant squid) pass the time. It automatically generates a random order
// in which to draw numbers and a random set of boards (your puzzle input).
// For example:
//
// 7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1
//
// 22 13 17 11  0
//
//	8  2 23  4 24
//
// 21  9 14 16  7
//
//	6 10  3 18  5
//	1 12 20 15 19
//
//	3 15  0  2 22
//	9 18 13 17  5
//
// 19  8  7 25 23
// 20 11 10 24  4
// 14 21 16 12  6
//
// 14 21 17 24  4
// 10 16 15  9 19
// 18  8 23 26 20
// 22 11 13  6  5
//
//	2  0 12  3  7
//
// After the first five numbers are drawn (7, 4, 9, 5, and 11),
// there are no winners, but the boards are marked as follows
// (shown here adjacent to each other to save space):
//
// 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
//
//	8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
//
// 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
//
//	6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
//	1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
//
// After the next six numbers are drawn (17, 23, 2, 0, 14, and 21),
// there are still no winners:
//
// 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
//
//	8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
//
// 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
//
//	6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
//	1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
//
// Finally, 24 is drawn:
//
// 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
//
//	8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
//
// 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
//
//	6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
//	1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
//
// At this point, the third board wins because it has at least one
// complete row or column of marked numbers (in this case, the entire
// top row is marked: 14 21 17 24 4).
//
// The score of the winning board can now be calculated. Start by
// finding the sum of all unmarked numbers on that board; in this
// case, the sum is 188. Then, multiply that sum by the number that
// was just called when the board won, 24, to get the final score,
// 188 * 24 = 4512.
//
// To guarantee victory against the giant squid, figure out which board
// will win first. What will your final score be if you choose that board?

type Num struct {
	value  string
	marked bool
}

func main() {
	score := calculateScore()
	fmt.Printf("Score: %v \n", score)
}

func calculateScore() int {
	fileScanner := createFileScanner()
	randomNums := []string{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		randomNums = parseRandomNum(currLine)
		break
	}
	fmt.Printf("random nums: %v \n", randomNums)
	input := parseInput(fileScanner)
	fmt.Printf("Input: %v \n", input)
	score := findWinningBoard(randomNums, &input)
	return score
}

func findWinningBoard(randomNums []string, input *[][][]Num) int {
	randomNumsLen := len(randomNums)
	n := len(*input)
	for i := 0; i < randomNumsLen; i++ {
		currNum := randomNums[i]
		// mark curr num
		for j := 0; j < n; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if (*input)[j][k][l].value == currNum {
						(*input)[j][k][l].marked = true
					}
				}
			}
			// check if the curr board won
			// check row wise
			for k := 0; k < 5; k++ {
				rowsMarked := 0
				for l := 0; l < 5; l++ {
					if (*input)[j][k][l].marked {
						rowsMarked++
					}
				}
				if rowsMarked == 5 {
					remainingSum := sumOfAllUnmarkedNums((*input)[j])
					currNumAsInt := convStrToInt(currNum)
					return remainingSum * currNumAsInt
				}
			}
			// check column wise
			for k := 0; k < 5; k++ {
				colsMarked := 0
				for l := 0; l < 5; l++ {
					if (*input)[j][l][k].marked {
						colsMarked++
					}
				}
				if colsMarked == 5 {
					remainingSum := sumOfAllUnmarkedNums((*input)[j])
					currNumAsInt := convStrToInt(currNum)
					return remainingSum * currNumAsInt
				}
			}
		}
	}
	return 0
}

func sumOfAllUnmarkedNums(input [][]Num) int {
	total := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !(input)[i][j].marked {
				total += convStrToInt((input)[i][j].value)
			}
		}
	}
	return total
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseInput(fileScanner *bufio.Scanner) [][][]Num {
	input := [][][]Num{}
	currBoard := [][]Num{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		if len(currLine) == 0 && len(currBoard) == 0 {
			continue
		}
		if len(currLine) == 0 {
			input = append(input, currBoard)
			currBoard = [][]Num{}
			continue
		}
		parsedLine := parse(currLine)
		currBoard = append(currBoard, parsedLine)
	}
	return input
}

func parse(line string) []Num {
	arr := []Num{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if j == n-1 {
				curr += string(line[j])
				n := Num{
					value:  curr,
					marked: false,
				}
				arr = append(arr, n)
				i = j
				break
			}
			if string(line[j]) == " " {
				n := Num{
					value:  curr,
					marked: false,
				}
				arr = append(arr, n)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return arr
}

func parseRandomNum(line string) []string {
	randomNums := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		curr := ""
		for j := i; j < n; j++ {
			if j == n-1 {
				curr += string(line[j])
				randomNums = append(randomNums, curr)
				i = j
				break
			}
			if string(line[j]) == "," {
				randomNums = append(randomNums, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return randomNums
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
