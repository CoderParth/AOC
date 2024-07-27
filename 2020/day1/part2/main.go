package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 1721
// 979
// 366
// 299
// 675
// 1456
//
//Using the above example again, the three entries
//that sum to 2020 are 979, 366, and 675. Multiplying
//them together produces the answer, 241861950.
//
// In your expense report, what is the product of
// the three entries that sum to 2020?

func main() {
	arrOfNums := createMapFromInput()
	threeNums := findThreeNumsThatSumTo2020(arrOfNums)
	multipleOfThreeNums := multiplyThreeNums(threeNums)
	fmt.Printf("multiple of three nums: %v \n", multipleOfThreeNums)
}

func createMapFromInput() []int {
	arr := []int{}
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		num := getNumFromALine(currLine, n)
		arr = append(arr, num)
	}
	return arr
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

func getNumFromALine(currLine string, n int) int {
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

func findThreeNumsThatSumTo2020(arr []int) []int {
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 2020 {
					return []int{arr[i], arr[j], arr[k]}
				}
			}
		}
	}
	return []int{}
}

func multiplyThreeNums(arr []int) int {
	fmt.Printf("f: %v, s: %v, t: %v \n", arr[0], arr[1], arr[2])
	return arr[0] * arr[1] * arr[2]
}
