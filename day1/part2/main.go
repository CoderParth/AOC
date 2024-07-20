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
	"zero":  "0",
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
		chars := []rune(line)
		n := len(chars)
		totalInString := ""

		foundFirstNum := false
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				ch := rune(line[i])
				curr := line[i : j+1]
				if _, ok := mp[curr]; ok {
					foundFirstNum = true
					totalInString += mp[curr]
					break
				}

				if unicode.IsNumber(ch) {
					foundFirstNum = true
					totalInString += string(ch)
					break
				}
			}
			if foundFirstNum {
				break
			}
		}

		foundLastNum := false
		for i := n - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				ch := rune(line[i])
				curr := line[j : i+1]
				if _, ok := mp[curr]; ok {
					foundLastNum = true
					totalInString += mp[curr]
					break
				}

				if unicode.IsNumber(ch) {
					foundLastNum = true
					totalInString += string(ch)
					break
				}
			}
			if foundLastNum {
				break
			}
		}

		i, err := strconv.Atoi(totalInString)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		total += i
	}

	return total
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
