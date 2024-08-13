package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	parseInput()
	// shortestDistance := findShortestDistance()
	// fmt.Printf("shortest distance: %v \n", shortestDistance)
}

// parse the input
// create a map of (d1, d2) = distance and (d2, d1) = distance
// create a map of cities
// create a permutations of all the route
// get distances for all the routes
// the short distance is the answer
//

type Pair struct {
	d1 string
	d2 string
}

func parseInput() {
	graph := make(map[Pair]int)
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		d1, d2, distance := parseCitiesAndDistance(currLine, n)
		p1 := Pair{d1: d1, d2: d2}
		p2 := Pair{d1: d2, d2: d1}
		graph[p1] = distance
		graph[p2] = distance
	}
	cities := createSet(graph)           // creates set of cites
	routes := createPermutations(cities) // create permutations of all the routes
	shortestDistance := findShortestDistance(routes, graph)
	fmt.Printf("Shortest Distance: %v \n", shortestDistance)
}

func findShortestDistance(routes [][]string, graph map[Pair]int) int {
	m, n := len(routes), len(routes[0])
	shortest := math.MaxInt

	for i := 0; i < m; i++ {
		curr := 0
		for j := 0; j < n-1; j++ {
			d1 := routes[i][j]
			d2 := routes[i][j+1]
			currPair := Pair{d1: d1, d2: d2}
			if v, ok := graph[currPair]; ok {
				curr += v
			}
		}
		shortest = min(shortest, curr)
	}
	return shortest
}

func createPermutations(cities map[string]int) [][]string {
	arr := []string{}
	for k := range cities {
		arr = append(arr, k)
	}
	return permute(arr)
}

func permute(nums []string) [][]string {
	n := len(nums)
	ans := make([][]string, 0)
	curr := make([]string, 0, n)
	vis := make(map[int]int)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(curr) == n {
			ans = append(ans, append([]string{}, curr...))
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				vis[i]++
				curr = append(curr, nums[i])
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
				vis[i]--
			}
		}
	}
	backtrack(0)
	return ans
}

func createSet(graph map[Pair]int) map[string]int {
	st := make(map[string]int)
	for k := range graph {
		st[k.d1] = 0
	}
	return st
}

func parseCitiesAndDistance(line string, n int) (string, string, int) {
	input := []string{}
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		j := i
		for ; j < n; j++ {
			if string(line[j]) == " " {
				i = j
				break
			}
			curr += string(line[j])
		}
		if len(curr) > 0 {
			input = append(input, curr)
			i = j
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
