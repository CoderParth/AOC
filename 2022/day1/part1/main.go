package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 1000
// 2000
// 3000
//
// 4000
//
// 5000
// 6000
//
// 7000
// 8000
// 9000
//
// 10000
//
// In case the Elves get hungry and need extra snacks,
// they need to know which Elf to ask: they'd like to
// know how many Calories are being carried by the Elf
// carrying the most Calories. In the example above,
// this is 24000 (carried by the fourth Elf).
//
// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := findTheListOfCalories() // list of calories stored in maxHeap
	highestCal := calculatehighestCal(h)
	fmt.Printf("highest calorie: %v \n", highestCal)
}

func findTheListOfCalories() *maxHeap {
	fileScanner := createFileScanner()
	h := &maxHeap{} // list of calories stored in maxheap
	heap.Init(h)
	totalCalsPerElf := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		if n == 0 {
			heap.Push(h, int(totalCalsPerElf))
			totalCalsPerElf = 0
			continue
		}
		calorie := findCalorieForThisLine(currLine, n)
		totalCalsPerElf += calorie
	}
	return h
}

func findCalorieForThisLine(currLine string, n int) int {
	currNumInStr := ""
	for i := 0; i < n; i++ {
		currNumInStr += string(currLine[i])
	}
	currNum, err := strconv.Atoi(currNumInStr)
	if err != nil {
		log.Fatal(err)
	}
	return currNum
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

func calculatehighestCal(h *maxHeap) int {
	x := heap.Pop(h).(int)
	return x
}
