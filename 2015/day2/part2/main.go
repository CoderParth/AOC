package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// The elves are also running low on ribbon. Ribbon is all the same width,
// so they only have to worry about the length they need to order, which
// they would again like to be exact.
//
// The ribbon required to wrap a present is the shortest distance around
// its sides, or the smallest perimeter of any one face. Each present
// also requires a bow made out of ribbon as well; the feet of ribbon
// required for the perfect bow is equal to the cubic feet of volume of
// the present. Don't ask how they tie the bow, though; they'll never tell.
//
// For example:
//
// A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to
// wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
// A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to
// wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
// How many total feet of ribbon should they order?
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
	return calculateTotal(lNum, wNum, hNum)
}

func calculateTotal(a, b, c int) int {
	p1 := 2 * (a + b)
	p2 := 2 * (b + c)
	p3 := 2 * (a + c)
	smallest := findSmallest(p1, p2, p3)
	v := a * b * c
	return smallest + v
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
