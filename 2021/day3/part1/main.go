package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Day 3: Binary Diagnostic ---
// The submarine has been making some odd creaking noises, so you ask
// it to produce a diagnostic report just in case.
//
// The diagnostic report (your puzzle input) consists of a list of binary
// numbers which, when decoded properly, can tell you many useful things
// about the conditions of the submarine. The first parameter to check is
// the power consumption.
//
// You need to use the binary numbers in the diagnostic report to generate
// two new binary numbers (called the gamma rate and the epsilon rate).
// The power consumption can then be found by multiplying the gamma rate
// by the epsilon rate.
//
// Each bit in the gamma rate can be determined by finding the most common
// bit in the corresponding position of all numbers in the diagnostic report.
// For example, given the following diagnostic report:
//
// 00100
// 11110
// 10110
// 10111
// 10101
// 01111
// 00111
// 11100
// 10000
// 11001
// 00010
// 01010
//
// Considering only the first bit of each number, there are five 0 bits
// and seven 1 bits. Since the most common bit is 1, the first bit of
// the gamma rate is 1.
//
// The most common second bit of the numbers in the diagnostic report is
// 0, so the second bit of the gamma rate is 0.
//
// The most common value of the third, fourth, and fifth bits are 1, 1,
// and 0, respectively, and so the final three bits of the gamma rate are 110.
//
// So, the gamma rate is the binary number 10110, or 22 in decimal.
//
// The epsilon rate is calculated in a similar way; rather than use
// the most common bit, the least common bit from each position is used.
// So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma
// rate (22) by the epsilon rate (9) produces the power consumption, 198.
//
// Use the binary numbers in your diagnostic report to calculate the gamma
// rate and epsilon rate, then multiply them together. What is the power
// consumption of the submarine? (Be sure to represent your answer in
// decimal, not binary.)
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner)
	fmt.Printf("Input: %v \n", input)
	power := calculatePower(input)
	fmt.Printf("Power: %v \n", power)
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

func parseInput(fileScanner *bufio.Scanner) [][]string {
	input := [][]string{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		currArr := parse(currLine)
		input = append(input, currArr)
	}
	return input
}

func parse(line string) []string {
	n := len(line)
	arr := []string{}
	for i := 0; i < n; i++ {
		arr = append(arr, string(line[i]))
	}
	return arr
}

func calculatePower(input [][]string) int {
	m, n := len(input), len(input[0])
	gamma, epsilon := "", ""

	for i := 0; i < n; i++ {
		freqMap := make(map[string]int)
		for j := 0; j < m; j++ {
			freqMap[input[j][i]]++
		}
		binaryGamma, binaryEpsilon := analyseBits(freqMap)
		gamma += binaryGamma
		epsilon += binaryEpsilon
	}
	decimalGamma, decimalEpsilon := convIntoDecimal(gamma), convIntoDecimal(epsilon)
	return decimalGamma * decimalEpsilon
}

func analyseBits(freqMap map[string]int) (string, string) {
	if freqMap["1"] > freqMap["0"] {
		return "1", "0"
	}
	return "0", "1"
}

func convIntoDecimal(binary string) int {
	n, remainingLength := len(binary), len(binary)-1
	decimal := 0
	for i := 0; i < n; i++ {
		currBinary := convStrToInt(string(binary[i]))
		decimal += twoToThePowerOf(remainingLength) * currBinary
		remainingLength--
	}
	return decimal
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func twoToThePowerOf(num int) int {
	if num == 0 {
		return 1
	}
	res := 1
	for i := 1; i <= num; i++ {
		res = 2 * res
	}
	return res
}
