package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// --- Part Two ---
// To be safe, the CPU also needs to know the highest value held in any register
// during this process so that it can decide how much memory to allocate to these
// operations. For example, in the above instructions, the highest value ever held
// was 10 (in register c after the third instruction was evaluated).
func main() {
	fileScanner := createFileScanner()
	arrOfInput, mpOfRegisters := parseInput(fileScanner)
	highestValue := followConditions(arrOfInput, &mpOfRegisters)
	fmt.Printf("Highest value: %v \n", highestValue)
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

func followConditions(arrOfInput [][]string, mpOfRegisters *map[string]int) int {
	n := len(arrOfInput)
	highestValue := math.MinInt
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
				if highestValue < (*mpOfRegisters)[firstRegister] {
					highestValue = (*mpOfRegisters)[firstRegister]
				}
			}
			if incOrDec == "dec" {
				(*mpOfRegisters)[firstRegister] -= value
				if highestValue < (*mpOfRegisters)[firstRegister] {
					highestValue = (*mpOfRegisters)[firstRegister]
				}
			}
		}
	}
	return highestValue
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
