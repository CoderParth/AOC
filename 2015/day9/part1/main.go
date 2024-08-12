package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// This year, however, he has some new locations to visit; his elves
// have provided him the distances between every pair of locations.
// He can start and end at any two (different) locations he wants, but
// he must visit each location exactly once. What is the shortest distance
// he can travel to achieve this?
//
// For example, given the following distances:
//
// London to Dublin = 464
// London to Belfast = 518
// Dublin to Belfast = 141
// The possible routes are therefore:
//
// Dublin -> London -> Belfast = 982
// London -> Dublin -> Belfast = 605
// London -> Belfast -> Dublin = 659
// Dublin -> Belfast -> London = 659
// Belfast -> Dublin -> London = 605
// Belfast -> London -> Dublin = 982
// The shortest of these is London -> Dublin -> Belfast = 605, and so
// the answer is 605 in this example.
//
// What is the distance of the shortest route?
func main() {
	input := parseInput()
	shortestDistance := findShortestDistance()
	fmt.Printf("shortest distance: %v \n", shortestDistance)
}

// parse the input
// create a map of (d1, d2) = distance and (d2, d1) = distance
// create a map of cities
// create a permutations of all the route
// get distances for all the routes
// the short distance is the answer
//

func parseInput() {
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		d1, d2, distance := parseCitiesAndDistance(currLine, n)
	}
}

func parseCitiesAndDistance(line string, n int) (string, string, int) {
	input := []string{}
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
				break
			}
			curr += string(line[j])
		}
		if len(curr) > 0 {
			input = append(input, curr)
		}
	}
	distance := convertStrToInt(input[len(input)-1])
	return input[0], input[2], distance
}

func convertStrToInt(n string) int {
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
