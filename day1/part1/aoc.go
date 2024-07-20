package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// input is a a string of multiple lines containing a word each line
//
// example:
// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
//
// the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.
func sumOfCalibrationValues() int {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		chars := []rune(line)
		n := len(chars)
		ch := ""
		for i := 0; i < len(chars); i++ {
			if unicode.IsNumber(chars[i]) {
				ch += string(chars[i])
				break
			}
		}

		for j := n - 1; j >= 0; j-- {
			if unicode.IsNumber(chars[j]) {
				ch += string(chars[j])
				break

			}
		}
		i, err := strconv.Atoi(ch)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		total += i
	}

	return total
}

func main() {
	total := sumOfCalibrationValues()
	fmt.Println(total)
}
