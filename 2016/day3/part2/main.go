package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// --- Part Two ---
// Now that you've helpfully marked up their design documents,
// it occurs to you that triangles are specified in groups of
// three vertically. Each set of three numbers in a column
// specifies a triangle. Rows are unrelated.
//
// For example, given the following specification, numbers with
// the same hundreds digit would be part of the same triangle:
//
// 101 301 501
// 102 302 502
// 103 303 503
// 201 401 601
// 202 402 602
// 203 403 603
// In your puzzle input, and instead reading by columns, how many of the listed triangles are possible?
func main() {
	total := totalListedTriangles()
	fmt.Printf("Total: %v \n", total)
}

func totalListedTriangles() int {
	total := 0
	fileScanner := createFileScanner()
	numOfLines := countNumOfLines(fileScanner)
	fileScanner = createFileScanner()
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	for i := 0; i < numOfLines; i++ {
		arr := [][]int{}
		j := i
		for ; j < i+3 && j < numOfLines; j++ {
			currLine := lines[j]
			n := len(currLine)
			numsFromStr := extractNumsFromString(currLine, n)
			arr = append(arr, numsFromStr)
		}
		i = j - 1
		if isSumOfTwoSidesLarger([]int{arr[0][0], arr[1][0], arr[2][0]}) {
			total++
		}
		if isSumOfTwoSidesLarger([]int{arr[0][1], arr[1][1], arr[2][1]}) {
			total++
		}
		if isSumOfTwoSidesLarger([]int{arr[0][2], arr[1][2], arr[2][2]}) {
			total++
		}
	}
	return total
}

func extractNumsFromString(line string, n int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
				i = j
				num := convStrToInt(curr)
				arr = append(arr, num)
				break
			}
			curr += string(line[j])
		}
	}
	return arr
}

func countNumOfLines(fileScanner *bufio.Scanner) int {
	n := 0
	for fileScanner.Scan() {
		n++
	}
	return n
}

func isTriangle(line string, n int) bool {
	arr := []int{}
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
				i = j
				num := convStrToInt(curr)
				arr = append(arr, num)
				break
			}
			curr += string(line[j])
		}
	}
	return isSumOfTwoSidesLarger(arr)
}

func isSumOfTwoSidesLarger(arr []int) bool {
	sort.Ints(arr)
	return arr[0]+arr[1] > arr[2]
}

func convStrToInt(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	return num
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
