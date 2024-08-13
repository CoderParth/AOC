package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// --- Part Two ---
// The next year, just to show off, Santa decides to take the
// route with the longest distance instead.
//
// He can still start and end at any two (different) locations
// he wants, and he still must visit each location exactly once.
//
// For example, given the distances above, the longest route
// would be 982 via (for example) Dublin -> London -> Belfast.
//
// What is the distance of the longest route?

type Pair struct {
	d1 string
	d2 string
}

func main() {
	routes, graph := parseInput()
	longestDistance := findLongestDistance(routes, graph)
	fmt.Printf("longest distance: %v \n", longestDistance)
}

func parseInput() ([][]string, map[Pair]int) {
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
	return routes, graph
}

func findLongestDistance(routes [][]string, graph map[Pair]int) int {
	m, n := len(routes), len(routes[0])
	longest := math.MinInt

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
		longest = max(longest, curr)
	}
	return longest
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
