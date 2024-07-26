package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
//
//
// In the above example, card 1 has five winning numbers
// (41, 48, 83, 86, and 17) and eight numbers you have
// (83, 86, 6, 31, 17, 9, 48, and 53). Of the numbers you
// have, four of them (48, 83, 17, and 86) are winning numbers!
// That means card 1 is worth 8 points (1 for the first match,
// then doubled three times for each of the three matches
// after the first).
//
// Card 2 has two winning numbers (32 and 61), so it is worth 2 points.
// Card 3 has two winning numbers (1 and 21), so it is worth 2 points.
// Card 4 has one winning number (84), so it is worth 1 point.
// Card 5 has no winning numbers, so it is worth no points.
// Card 6 has no winning numbers, so it is worth no points.
// So, in this example, the Elf's pile of scratchcards is worth 13 points.
//
//How many points are they worth in total?

func main() {
	points := calculatePoints()
	fmt.Printf("Points Array: %v \n", points)
	totalPoints := sumOfPoints(points)
	fmt.Printf("points: %v \n", totalPoints)
}

func calculatePoints() []int {
	fileScanner := createFileScanner()
	n := findLenOfStrInALine(fileScanner)

	fileScanner = createFileScanner()
	pointsPerEachLine := calculatePointsPerEachLine(fileScanner, n)
	return pointsPerEachLine
}

func calculatePointsPerEachLine(fileScanner *bufio.Scanner, n int) []int {
	pointsArray := []int{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		totalWinningNums := findTotalNumOfWinningNums(currLine, n)
		pointsForThisCard := calculatePointsFromWinningNums(totalWinningNums)
		pointsArray = append(pointsArray, pointsForThisCard)
	}
	return pointsArray
}

func findTotalNumOfWinningNums(currLine string, n int) int {
	mpOfWinningNums := createMapOfWinningNums(currLine, n)
	mpOfNumsYouHave := createMapOfNumsYouHave(currLine, n)
	return findHowManyMatches(mpOfWinningNums, mpOfNumsYouHave)
}

func createMapOfWinningNums(currLine string, n int) map[int]int {
	mp := make(map[int]int)
	for i := 4; string(currLine[i]) != "|" && i < n; i++ {
		if string(currLine[i]) != ":" {
			continue
		}
		for j := i + 2; string(currLine[j]) != "|" && j < n; j++ {
			fmt.Printf("j: %v \n", j)
			currNumInStr := ""
			k := j
			for ; string(currLine[k]) != " " && k < n; k++ {
				fmt.Println("running k loop ")
				currNumInStr += string(currLine[k])
			}
			fmt.Printf("curr num: %v \n", currNumInStr)
			currNum, err := strconv.Atoi(currNumInStr)
			if err != nil {
				log.Fatal(err)
			}
			mp[currNum] = 0
			j = k
		}
	}
	return mp
}

func createMapOfNumsYouHave(currLine string, n int) map[int]int {
	mp := make(map[int]int)
	startingIdxForNumsYouHave := 0
	for i := 4; i < n; i++ {
		if string(currLine[i]) == "|" {
			startingIdxForNumsYouHave = i
			break
		}
	}

	for i := startingIdxForNumsYouHave + 1; i < n; i++ {
		for j := i + 2; j < n; j++ {
			currNumInStr := ""
			k := j
			for ; string(currLine[k]) != "" && k < n; k++ {
				currNumInStr += string(currLine[k])
			}
			currNum, err := strconv.Atoi(currNumInStr)
			if err != nil {
				log.Fatal(err)
			}
			mp[currNum] = 0
			j = k
		}
	}
	return mp
}

func findHowManyMatches(mp1, mp2 map[int]int) int {
	totalMatches := 0
	for k := range mp2 {
		if _, ok := mp1[k]; ok {
			totalMatches++
		}
	}
	return totalMatches
}

func calculatePointsFromWinningNums(lenOfWinningNums int) int {
	curr := 1
	for i := 1; i < lenOfWinningNums; i++ {
		curr *= 2
	}
	return curr
}

func findLenOfStrInALine(fileScanner *bufio.Scanner) int {
	n := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n = len(currLine)
		break
	}
	return n
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

func sumOfPoints(points []int) int {
	total := 0
	for _, p := range points {
		total += p
	}
	return total
}
