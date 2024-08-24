package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// Now that the machine is calibrated, you're ready to begin molecule
// fabrication.
//
// Molecule fabrication always begins with just a single electron, e,
// and applying replacements one at a time, just like the ones
// during calibration.
//
// For example, suppose you have the following replacements:
//
// e => H
// e => O
// H => HO
// H => OH
// O => HH
// If you'd like to make HOH, you start with e, and then make the
// following replacements:
//
// e => O to get O
// O => HH to get HH
// H => OH (on the second H) to get HOH
// So, you could make HOH after 3 steps. Santa's favorite molecule,
// HOHOHO, can be made in 6 steps.
//
// How long will it take to make the medicine? Given the available
// replacements and the medicine molecule in your puzzle input,
// what is the fewest number of steps to go from e to the medicine molecule?
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner) // create array of all the input
	rMap := createReplacementMap(input)
	fmt.Printf("Replacement map: %v \n", rMap)
	medicineMolecule := input[len(input)-1][0]
	fmt.Printf("Medicine Molecule: %v \n", medicineMolecule)
	fewestSteps := fewestStepsFromE(rMap, medicineMolecule) // from e to medicine molecule
	fmt.Printf("Fewest Steps: %v \n", fewestSteps)
}

func fewestStepsFromE(rMap map[string][]string, target string) int {
	queue := []string{"e"}
	visited := map[string]int{"e": 0}

	for len(queue) > 0 {
		fmt.Printf("Q len: %v \n", len(queue))
		current := queue[0]
		queue = queue[1:]
		currentSteps := visited[current]

		if current == target {
			return currentSteps
		}

		for i := 0; i < len(current); i++ {
			singleChar := string(current[i])
			if replacements, ok := rMap[singleChar]; ok {
				for _, replacement := range replacements {
					newMolecule := current[:i] + replacement + current[i+1:]
					if _, seen := visited[newMolecule]; !seen {
						visited[newMolecule] = currentSteps + 1
						queue = append(queue, newMolecule)
					}
				}
			}

			if i < len(current)-1 {
				doubleChar := current[i : i+2]
				if replacements, ok := rMap[doubleChar]; ok {
					for _, replacement := range replacements {
						newMolecule := current[:i] + replacement + current[i+2:]
						if _, seen := visited[newMolecule]; !seen {
							visited[newMolecule] = currentSteps + 1
							queue = append(queue, newMolecule)
						}
					}
				}
			}
		}
	}
	return -1
}

func findDistinctMolecules(rMap map[string][]string, m string) int {
	dMap := make(map[string]int) // map of distinct molecules
	for i := 0; i < len(m); i++ {
		singleChar := string(m[i])
		if _, ok := rMap[singleChar]; ok {
			for _, v := range rMap[singleChar] {
				rep := m[0:i]
				rep += v
				rep += m[i+1:]
				dMap[rep]++
			}
		}
		if i < len(m)-1 {
			dubChar := string(m[i]) + string(m[i+1])
			if _, ok := rMap[dubChar]; ok {
				for _, v := range rMap[dubChar] {
					rep := m[0:i]
					rep += v
					if i < len(m)-2 {
						rep += m[i+2:]
					}
					dMap[rep]++
				}
			}
		}
	}
	return len(dMap)
}

func findNumOfSteps(rep string, rMap map[string][]string, med string) int {
	steps := 0
	for rep != med {
		for i := 0; i < len(rep); i++ {
			singleChar := string(rep[i])
			if _, ok := rMap[singleChar]; ok {
				for _, v := range rMap[singleChar] {
					rep := rep[0:i]
					rep += v
					rep += rep[i+1:]
					steps++
				}
			}
			if i < len(rep)-1 {
				dubChar := string(rep[i]) + string(rep[i+1])
				if _, ok := rMap[dubChar]; ok {
					for _, v := range rMap[dubChar] {
						rep := rep[0:i]
						rep += v
						if i < len(rep)-2 {
							rep += rep[i+2:]
						}
						steps++
					}
				}
			}
		}
	}
	return steps
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

func parseInput(fileScanner *bufio.Scanner) [][]string {
	input := [][]string{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		arr := createArrFromCurrLine(currLine)
		input = append(input, arr)
	}
	return input
}

func createArrFromCurrLine(line string) []string {
	arr := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
				arr = append(arr, curr)
				i = j
				break
			}
			if j == n-1 {
				curr += string(line[j])
				arr = append(arr, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return arr
}

func createReplacementMap(input [][]string) map[string][]string {
	rMp := make(map[string][]string)
	m := len(input)
	for i := 0; i < m-1; i++ {
		rMp[input[i][0]] = append(rMp[input[i][0]], input[i][2])
	}
	return rMp
}
