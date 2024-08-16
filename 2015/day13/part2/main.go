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
// In all the commotion, you realize that you forgot to seat yourself.
// At this point, you're pretty apathetic toward the whole thing, and
// your happiness wouldn't really go up or down regardless of who you
// sit next to. You assume everyone else would be just as ambivalent
// about sitting next to you, too.
//
// So, add yourself to the list, and give all happiness relationships
// that involve you a score of 0.
//
// What is the total change in happiness for the optimal seating
// arrangement that actually includes yourself?
type Pair struct {
	firstPerson  string
	secondPerson string
}

func main() {
	input := parseInput()
	graph := createGraph(input)
	set := createSet(graph)
	graph = addYouInTheGraph(graph, set)
	set = append(set, "ME")
	allArangements := permute(set)
	optimalHappiness := findOptimalHappiness(graph, allArangements)
	fmt.Printf("Optimal Happiness: %v \n", optimalHappiness)
}

func addYouInTheGraph(graph map[Pair]int, set []string) map[Pair]int {
	for i := 0; i < len(set); i++ {
		curr := set[i]
		// person -> ME
		p := Pair{
			firstPerson:  curr,
			secondPerson: "ME",
		}
		graph[p] = 0
		// ME -> person
		p2 := Pair{
			firstPerson:  "ME",
			secondPerson: curr,
		}
		graph[p2] = 0
	}
	return graph
}

func findOptimalHappiness(graph map[Pair]int, arr [][]string) int {
	h := math.MinInt
	m, n := len(arr), len(arr[0])
	for i := 0; i < m; i++ {
		curr := 0
		// clockwise
		for j := 0; j < n-1; j++ {
			firstPerson := arr[i][j]
			secondPerson := arr[i][j+1]
			p := Pair{
				firstPerson,
				secondPerson,
			}
			curr += graph[p]
		}
		// for the last and the first
		f := arr[i][n-1]
		s := arr[i][0]
		p := Pair{
			firstPerson:  f,
			secondPerson: s,
		}
		curr += graph[p]
		// anti clock wise
		for j := n - 1; j >= 1; j-- {
			firstPerson := arr[i][j]
			secondPerson := arr[i][j-1]
			p := Pair{
				firstPerson,
				secondPerson,
			}
			curr += graph[p]
		}
		// for the first and the last
		f = arr[i][0]
		s = arr[i][n-1]
		p = Pair{
			firstPerson:  f,
			secondPerson: s,
		}
		curr += graph[p]
		h = max(h, curr)
	}
	return h
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

func createSet(graph map[Pair]int) []string {
	mp := map[string]int{}
	for k := range graph {
		mp[k.firstPerson] = 0
	}
	set := []string{}
	for k := range mp {
		set = append(set, k)
	}
	return set
}

func createGraph(input [][]string) map[Pair]int {
	graph := make(map[Pair]int)
	m := len(input)
	for i := 0; i < m; i++ {
		currPerson := string(input[i][0])
		nextPerson := string(input[i][3])
		gainOrLose := string(input[i][1])
		amount := string(input[i][2])
		pair := Pair{
			firstPerson:  currPerson,
			secondPerson: nextPerson,
		}
		a, err := strconv.Atoi(amount)
		if err != nil {
			log.Fatal(err)
		}
		if string(gainOrLose) == "lose" {
			a = -a
		}
		graph[pair] = a
	}
	return graph
}

func parseInput() [][]string {
	input := [][]string{}
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		curr := findHappinessInCurrLine(currLine, n)
		input = append(input, curr)
	}
	return input
}

func findHappinessInCurrLine(line string, n int) []string {
	arr := []string{}
	for i := 0; i < n; i++ {
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " || j == n-1 {
				arr = append(arr, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	newArr := []string{} // stores curr person, gain/loss,amount,  and next person
	newArr = append(newArr, arr[0])
	newArr = append(newArr, arr[2])
	newArr = append(newArr, arr[3])
	newArr = append(newArr, arr[len(arr)-1])
	return newArr
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
