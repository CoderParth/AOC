package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Alice would gain 54 happiness units by sitting next to Bob.
// Alice would lose 79 happiness units by sitting next to Carol.
// Alice would lose 2 happiness units by sitting next to David.
// Bob would gain 83 happiness units by sitting next to Alice.
// Bob would lose 7 happiness units by sitting next to Carol.
// Bob would lose 63 happiness units by sitting next to David.
// Carol would lose 62 happiness units by sitting next to Alice.
// Carol would gain 60 happiness units by sitting next to Bob.
// Carol would gain 55 happiness units by sitting next to David.
// David would gain 46 happiness units by sitting next to Alice.
// David would lose 7 happiness units by sitting next to Bob.
// David would gain 41 happiness units by sitting next to Carol.
//
//	+41 +46
//
// +55   David    -2
// Carol       Alice
// +60    Bob    +54
//
//	-7  +83
//
// After trying every other seating arrangement in this hypothetical
// scenario, you find that this one is the most optimal, with a total
// change in happiness of 330.
//
// What is the total change in happiness for the optimal seating
// arrangement of the actual guest list?
type Pair struct {
	firstPerson  string
	secondPerson string
}

func main() {
	input := parseInput()
	graph := createGraph(input)
	set := createSet(graph)
	allArangements := permute(set)
	fmt.Printf("Graph: %v \n", graph)
	optimalHappiness := findOptimalHappiness(graph, allArangements)
	fmt.Printf("Optimal Happiness: %v \n", optimalHappiness)
}

func findOptimalHappiness(graph map[Pair]int, arr [][]string) int {
	h := math.MinInt
	m, n := len(arr), len(arr[0])
	for i := 0; i < m; i++ {
		curr := 0
		for j := 0; j < n-1; j++ {
			firstPerson := arr[i][j]
			secondPerson := arr[i][j+1]
			p := Pair{
				firstPerson,
				secondPerson,
			}
			curr += graph[p]
			p = Pair{
				secondPerson,
				firstPerson,
			}
			curr += graph[p]

		}
		fmt.Printf("Curr Arrangement: %v \n", arr[i])
		fmt.Printf("curr amount: %v \n", curr)
		h = max(h, curr)
	}
	return h
}

// find the max

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
		fmt.Printf("curr input: %v\n", input[i])
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
