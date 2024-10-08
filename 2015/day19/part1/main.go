package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Day 19: Medicine for Rudolph ---
// Rudolph the Red-Nosed Reindeer is sick! His
// nose isn't shining very brightly, and he
// needs medicine.
//
// Red-Nosed Reindeer biology isn't similar to
// regular reindeer biology; Rudolph is going
// to need custom-made medicine. Unfortunately,
// Red-Nosed Reindeer chemistry isn't similar
// to regular reindeer chemistry, either.
//
// The North Pole is equipped with a Red-Nosed
// Reindeer nuclear fusion/fission plant, capable
// of constructing any Red-Nosed Reindeer molecule
// you need. It works by starting with some input
// molecule and then doing a series of replacements,
// one per step, until it has the right molecule.
//
// However, the machine has to be calibrated before it
// can be used. Calibration involves determining the
// number of molecules that can be generated in one
// step from a given starting point.
//
// For example, imagine a simpler machine that supports
// only the following replacements:
//
// H => HO
// H => OH
// O => HH
// Given the replacements above and starting with HOH,
// the following molecules could be generated:
//
// HOOH (via H => HO on the first H).
// HOHO (via H => HO on the second H).
// OHOH (via H => OH on the first H).
// HOOH (via H => OH on the second H).
// HHHH (via O => HH).
//
// So, in the example above, there are 4 distinct molecules
// (not five, because HOOH appears twice) after one
// replacement from HOH. Santa's favorite molecule, HOHOHO,
// can become 7 distinct molecules (over nine replacements:
// six from H, and three from O).
//
// The machine replaces without regard for the surrounding
// characters. For example, given the string H2O, the
// transition H => OO would result in OO2O.
//
// Your puzzle input describes all of the possible replacements
// and, at the bottom, the medicine molecule for which you need
// to calibrate the machine. How many distinct molecules can be
// created after all the different ways you can do one
// replacement on the medicine molecule?
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner) // create array of all the input
	fmt.Printf("Input: %v \n", input)
	rMap := createReplacementMap(input)
	fmt.Printf("Replacement map: %v \n", rMap)
	medicineMolecule := input[len(input)-1][0]
	fmt.Printf("Medicine Molecule: %v \n", medicineMolecule)
	numOfDistinctMolecules := findDistinctMolecules(rMap, medicineMolecule)
	fmt.Printf("Num: %v \n", numOfDistinctMolecules)
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
