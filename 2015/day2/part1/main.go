package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l.
// The elves also need a little extra paper for each present: the
// area of the smallest side.
//
// For example:
//
// A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52
// square feet of wrapping paper plus 6 square feet of slack, for
// a total of 58 square feet.
// A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42
// square feet of wrapping paper plus 1 square foot of slack, for a
// total of 43 square feet.
//
// All numbers in the elves' list are in feet. How many total square
// feet of wrapping paper should they order?
func main() {
	elvesList := findElvesList()
	total := findTotalSquare(elvesList)
	fmt.Printf("Total: %v \n", total)
}

func findElvesList() []int {
	list := []int{}
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		area := findArea(currLine, n)
		list = append(list, area)
	}
	return list
}

func findArea(line string, n int) int {
	l, w, h := "", "", ""
	i := 0
	// find l
	for ; i < n; i++ {
		if string(line[i]) == "x" {
			i++
			break
		}
		l += string(line[i])
	}

	// find w
	for ; i < n; i++ {
		if string(line[i]) == "x" {
			i++
			break
		}
		w += string(line[i])
	}

	// find h
	for ; i < n; i++ {
		if string(line[i]) == "x" {
			i++
			break
		}
		h += string(line[i])
	}

	lNum, wNum, hNum := convertStrToNum(l, w, h)
	a, b, c := lNum*wNum, wNum*hNum, hNum*lNum
	smallest := findSmallest(a, b, c)
	return calculateTotal(a, b, c, smallest)
}

func calculateTotal(a, b, c, smallest int) int {
	totalSquareFeet := 2 * (a + b + c)
	return totalSquareFeet + smallest
}

func findSmallest(vars ...int) int {
	s := vars[0]
	for i := 1; i < len(vars); i++ {
		if vars[i] < s {
			s = vars[i]
		}
	}
	return s
}

func convertStrToNum(l, w, h string) (int, int, int) {
	lNum, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal(err)
	}

	wNum, err := strconv.Atoi(w)
	if err != nil {
		log.Fatal(err)
	}

	hNum, err := strconv.Atoi(h)
	if err != nil {
		log.Fatal(err)
	}

	return lNum, wNum, hNum
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

func findTotalSquare(arr []int) int {
	total := arr[0]
	for i := 1; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}
