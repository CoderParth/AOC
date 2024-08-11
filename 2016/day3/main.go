package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// --- Day 3: Squares With Three Sides ---
// Now that you can think clearly, you move deeper into the labyrinth
// of hallways and office furniture that makes up this part of Easter
// Bunny HQ. This must be a graphic design department; the walls are
// covered in specifications for triangles.
//
// Or are they?
//
// The design document gives the side lengths of each triangle it describes,
// but... 5 10 25? Some of these aren't triangles. You can't help but mark
// the impossible ones.
//
// In a valid triangle, the sum of any two sides must be larger than the
// remaining side. For example, the "triangle" given above is impossible,
// because 5 + 10 is not larger than 25.
//
// In your puzzle input, how many of the listed triangles are possible?
func main() {
	total := totalListedTriangles()
	fmt.Printf("Total: %v \n", total)
}

func totalListedTriangles() int {
	total := 0
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		if isTriangle(currLine, n) {
			total++
		}
	}
	return total
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
