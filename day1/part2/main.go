package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var mp = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen
//
// the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.
func main() {
	total := sumOfAllCalibrationValues()
	fmt.Println(total)
}

func sumOfAllCalibrationValues() int {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		n := len(line)

		totalInString := ""
		firstStrInt := findFirstStrInt(line, n)
		lastStrInt := findLastStrInt(line, n)
		totalInString = firstStrInt + lastStrInt

		i, err := strconv.Atoi(totalInString)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		total += i

	}
	return total
}

func findFirstStrInt(line string, n int) string {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ch := rune(line[i])
			curr := line[i : j+1]
			if _, ok := mp[curr]; ok {
				return mp[curr]
			}
			if unicode.IsNumber(ch) {
				return string(ch)
			}
		}
	}
	return ""
}

func findLastStrInt(line string, n int) string {
	for i := n - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			ch := rune(line[i])
			curr := line[j : i+1]
			if _, ok := mp[curr]; ok {
				return mp[curr]
			}
			if unicode.IsNumber(ch) {
				return string(ch)
			}
		}
	}
	return ""
}
