package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Calculate the horizontal position and depth you would have after following
// the planned course. What do you get if you multiply your final
// horizontal position by your final depth?
func main() {
	fileScanner := createFileScanner()
	horPos, depth := calculatePos(fileScanner)
	fmt.Printf("horizontal Position: %v \n", horPos)
	fmt.Printf("Depth: %v \n", depth)
	product := horPos * depth
	fmt.Printf("Product: %v \n", product)
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

func calculatePos(fileScanner *bufio.Scanner) (int, int) {
	horPos, depth := 0, 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		command, units := parseLine(currLine)
		switch command {
		case "forward":
			horPos += units
		case "down":
			depth += units
		case "up":
			depth -= units
		}
	}
	return horPos, depth
}

func parseLine(line string) (string, int) {
	n := len(line)
	i := 0
	command := ""
	for ; i < n; i++ {
		if string(line[i]) == " " {
			i++
			break
		}
		command += string(line[i])
	}
	unitsInStr := ""
	for ; i < n; i++ {
		unitsInStr += string(line[i])
	}
	units := convStrToInt(unitsInStr)
	return command, units
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
