package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// On the other hand, it might be wise to try a different strategy:
// let the giant squid win.
//
// You aren't sure how many bingo boards a giant squid could play
// at once, so rather than waste time counting its arms, the safe
// thing to do is to figure out which board will win last and
// choose that one. That way, no matter which boards it picks,
// it will win for sure.
//
// In the above example, the second board is the last to win, which
// happens after 13 is eventually called and its middle column is
// completely marked. If you were to keep playing until this point,
// the second board would have a sum of unmarked numbers equal to
// 148 for a final score of 148 * 13 = 1924.
//
// Figure out which board will win last. Once it wins, what would
// its final score be?
type Num struct {
	value  string
	marked bool
}

type Board struct {
	boardNum int
	won      bool
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
	input, numOfBoards := parseInput(fileScanner)
	fmt.Printf("Input: %v \n", input)
	boards := initializeBoards(numOfBoards)
	score := findWinningBoard(randomNums, &input, boards)
	return score
}

func initializeBoards(numOfBoards int) []Board {
	boards := make([]Board, numOfBoards)
	for i := 0; i < numOfBoards; i++ {
		boards[i] = Board{
			boardNum: i,
			won:      false,
		}
	}
	return boards
}

func findWinningBoard(randomNums []string, input *[][][]Num, boards []Board) int {
	randomNumsLen := len(randomNums)
	n := len(*input)
	lastResult := 0
	numOfWinners := 0
	totalBoards := len(boards)

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
					if hasWon := boards[j].won; !hasWon {
						boards[j].won = true
						numOfWinners++
						if numOfWinners == totalBoards {
							remainingSum := sumOfAllUnmarkedNums((*input)[j])
							currNumAsInt := convStrToInt(currNum)
							lastResult = remainingSum * currNumAsInt
						}
					}
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
					if hasWon := boards[j].won; !hasWon {
						boards[j].won = true
						numOfWinners++
						if numOfWinners == totalBoards {
							remainingSum := sumOfAllUnmarkedNums((*input)[j])
							currNumAsInt := convStrToInt(currNum)
							lastResult = remainingSum * currNumAsInt
						}
					}
				}
			}
		}
	}
	return lastResult
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

func parseInput(fileScanner *bufio.Scanner) ([][][]Num, int) {
	input := [][][]Num{}
	currBoard := [][]Num{}
	numOfBoards := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		if len(currLine) == 0 && len(currBoard) == 0 {
			continue
		}
		if len(currLine) == 0 {
			input = append(input, currBoard)
			currBoard = [][]Num{}
			numOfBoards++
			continue
		}
		parsedLine := parse(currLine)
		currBoard = append(currBoard, parsedLine)
	}
	return input, numOfBoards
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
