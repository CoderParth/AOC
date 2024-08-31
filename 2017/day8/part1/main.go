package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// --- Day 8: I Heard You Like Registers ---
// You receive a signal directly from the CPU. Because of your recent
// assistance with jump instructions, it would like you to compute
// the result of a series of unusual register instructions.
//
// Each instruction consists of several parts: the register to modify,
// whether to increase or decrease that register's value, the amount
// by which to increase or decrease it, and a condition. If the
// condition fails, skip the instruction without modifying the register.
// The registers all start at 0. The instructions look like this:
//
// b inc 5 if a > 1
// a inc 1 if b < 5
// c dec -10 if a >= 1 // decrementing a negative value is incrementing
// c inc -20 if c == 10
// These instructions would be processed as follows:
//
// Because a starts at 0, it is not greater than 1, and so b is not modified.
// a is increased by 1 (to 1) because b is less than 5 (it is 0).
// c is decreased by -10 (to 10) because a is now greater than or equal to 1 (it is 1).
// c is increased by -20 (to -10) because c is equal to 10.
// After this process, the largest value in any register is 1.
//
// You might also encounter <= (less than or equal to) or != (not equal to).
// However, the CPU doesn't have the bandwidth to tell you what all the registers
// are named, and leaves that to you to determine.
//
// What is the largest value in any register after completing the instructions
// in your puzzle input?
func main() {
	fileScanner := createFileScanner()
	arrOfInput, mpOfRegisters := parseInput(fileScanner)
	followConditions(arrOfInput, &mpOfRegisters)
	largest := findLargestRegister(mpOfRegisters)
	fmt.Printf("Largest: %v \n", largest)
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

func parseInput(fileScanner *bufio.Scanner) ([][]string, map[string]int) {
	arrOfInput := [][]string{}
	mpOfRegisters := make(map[string]int)
	for fileScanner.Scan() {
		currArr := parseCurrLine(fileScanner.Text())
		mpOfRegisters[currArr[0]] = 0
		mpOfRegisters[currArr[4]] = 0
		arrOfInput = append(arrOfInput, currArr)
	}
	return arrOfInput, mpOfRegisters
}

func parseCurrLine(line string) []string {
	arr := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
				arr = append(arr, curr)
				i = j
				break
			}
			if j == n-1 {
				curr += string(line[j])
				arr = append(arr, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return arr
}

func followConditions(arrOfInput [][]string, mpOfRegisters *map[string]int) {
	n := len(arrOfInput)
	for i := 0; i < n; i++ {
		firstRegister := arrOfInput[i][0]
		incOrDec := arrOfInput[i][1]
		value := convStrToInt(arrOfInput[i][2])
		secondRegister := arrOfInput[i][4]
		condition := arrOfInput[i][5]
		conditionVal := convStrToInt(arrOfInput[i][6])
		if conditionMatch(secondRegister, condition, conditionVal, mpOfRegisters) {
			if incOrDec == "inc" {
				(*mpOfRegisters)[firstRegister] += value
			}
			if incOrDec == "dec" {
				(*mpOfRegisters)[firstRegister] -= value
			}
		}
	}
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func conditionMatch(register string, condition string, conditionalVal int, mpOfRegisters *map[string]int) bool {
	regValue := (*mpOfRegisters)[register]
	switch condition {
	case ">":
		return regValue > conditionalVal
	case "<":
		return regValue < conditionalVal
	case ">=":
		return regValue >= conditionalVal
	case "<=":
		return regValue <= conditionalVal
	case "==":
		return regValue == conditionalVal
	case "!=":
		return regValue != conditionalVal
	}
	return false
}

func findLargestRegister(mp map[string]int) int {
	largest := math.MinInt
	for _, v := range mp {
		if v > largest {
			largest = v
		}
	}
	return largest
}
