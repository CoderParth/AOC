package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// An opening parenthesis, (, means he should go up one floor,
// and a closing parenthesis, ), means he should go down one
// floor.
//
// The apartment building is very tall, and the basement is very deep;
// he will never find the top or bottom floors.
//
// For example:
//
// (()) and ()() both result in floor 0.
// ((( and (()(()( both result in floor 3.
// ))((((( also results in floor 3.
// ()) and ))( both result in floor -1 (the first basement level).
// ))) and )())()) both result in floor -3.
// To what floor do the instructions take Santa?
func main() {
	floorToGo := findFloorToGo()
	fmt.Printf("Floor to go: %v \n", floorToGo)
}

func findFloorToGo() int {
	fileScanner := createFileScanner()
	floor := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		floor = calculateParanthesis(currLine, n)
	}
	return floor
}

func calculateParanthesis(line string, n int) int {
	curr := 0
	for i := 0; i < n; i++ {
		if string(line[i]) == "(" {
			curr++
		} else if string(line[i]) == ")" {
			curr--
		}
	}
	return curr
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
