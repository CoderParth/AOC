package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// You finally arrive at the bathroom (it's a several minute walk from the
// lobby so visitors can behold the many fancy conference rooms and water
// coolers on this floor) and go to punch in the code. Much to your
// bladder's dismay, the keypad is not at all like you imagined it.
// Instead, you are confronted with the result of hundreds of
// man-hours of bathroom-keypad-design meetings:
//
//	  1
//	2 3 4
//
// 5 6 7 8 9
//
//	A B C
//	  D
//
// You still start at "5" and stop when you're at an edge, but given
// the same instructions as above, the outcome is very different:
//
// You start at "5" and don't move at all (up and left are both edges),
// ending at 5.
// Continuing from "5", you move right twice and down three times
// (through "6", "7", "B", "D", "D"), ending at D.
// Then, from "D", you move five more times (through "D", "B",
// "C", "C", "B"), ending at B.
// Finally, after five more moves, you end at 3.
// So, given the actual keypad layout, the code would be 5DB3.
//
// Using the same instructions in your puzzle input, what is the
// correct bathroom code?
func main() {
	keypad := initializeKeyPad()
	fileScanner := createFileScanner()
	code := findCode(keypad, fileScanner)
	fmt.Printf("Code: %v \n", code)
}

func initializeKeyPad() [][]string {
	keypad := [][]string{
		{".", ".", "1", ".", "."},
		{".", "2", "3", "4", "."},
		{"5", "6", "7", "8", "9"},
		{".", "A", "B", "C", "."},
		{".", ".", "D", ".", "."},
	}
	return keypad
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

func findCode(keypad [][]string, fileScanner *bufio.Scanner) string {
	code := ""
	i, j := 2, 0 // starting idxs
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		button := findButton(currLine, keypad, &i, &j)
		code += button
	}
	return code
}

func findButton(line string, keypad [][]string, i, j *int) string {
	fmt.Printf("Curr i: %v \n j: %v \n", *i, *j)
	m, n := len(keypad), len(keypad[0])
	lenOfLine := len(line)
	for l := 0; l < lenOfLine; l++ {
		switch string(line[l]) {
		case "U":
			if *i > 0 && keypad[*i-1][*j] != "." {
				*i--
			}
		case "D":
			if *i < m-1 && keypad[*i+1][*j] != "." {
				*i++
			}
		case "L":
			if *j > 0 && keypad[*i][*j-1] != "." {
				*j--
			}
		case "R":
			if *j < n-1 && keypad[*i][*j+1] != "." {
				*j++
			}
		}
	}

	return keypad[*i][*j]
}
