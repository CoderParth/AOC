package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// There are more programs than just the ones in the group containing program
// ID 0. The rest of them have no way of reaching that group, and still might
// have no way of reaching each other.
//
// A group is a collection of programs that can all communicate via pipes
// either directly or indirectly. The programs you identified just a moment
// ago are all part of the same group. Now, they would like you to determine
// the total number of groups.
//
// In the example above, there were 2 groups: one consisting of programs
// 0,2,3,4,5,6, and the other consisting solely of program 1.
//
// How many groups are there in total?
func main() {
	numOfGroups := findNumOfGroups() // connected to zero
	fmt.Printf("Number of Groups: %v \n", numOfGroups)
}

func findNumOfGroups() int {
	programsMap := make(map[string][]string) // Acts as a graph
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		currProgram, idx := findCurrProgram(currLine)
		programsInGroup := findProgramsInGroup(currLine, idx)
		programsMap[currProgram] = programsInGroup
	}
	fmt.Printf("Maps: %v \n", programsMap)
	total := findTotalGroups(programsMap)
	return total
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

func findCurrProgram(line string) (string, int) {
	n := len(line)
	for i := 0; i < n; i++ {
		curr := ""
		for j := 0; j < n; j++ {
			if string(line[j]) == " " {
				return curr, j
			}
			curr += string(line[j])
		}
	}
	return "", 0
}

func findProgramsInGroup(line string, idx int) []string {
	group := []string{}
	idx += 5
	n := len(line)
	for i := idx; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == "," {
				group = append(group, curr)
				i = j + 1
				break
			}
			if j == n-1 {
				curr += string(line[j])
				group = append(group, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return group
}

func findTotalGroups(programsMap map[string][]string) int {
	visited := make(map[string]int)
	numOfGroups := 0
	for program := range programsMap {
		if _, ok := visited[program]; !ok {
			dfs(program, programsMap, &visited)
			numOfGroups++
		}
	}
	return numOfGroups
}

func dfs(program string, programsMap map[string][]string, visited *map[string]int) {
	if _, ok := (*visited)[program]; ok {
		return
	}
	(*visited)[program] = 0
	for _, groupProgam := range programsMap[program] {
		if _, ok := (*visited)[groupProgam]; !ok {
			dfs(groupProgam, programsMap, visited)
		}
	}
}
