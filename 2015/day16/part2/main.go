package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// As you're about to send the thank you note, something in the MFCSAM's
// instructions catches your eye. Apparently, it has an outdated
// retroencabulator, and so the output from the machine isn't exact
// values - some of them indicate ranges.
//
// In particular, the cats and trees readings indicates that there are
// greater than that many (due to the unpredictable nuclear decay of
// cat dander and tree pollen), while the pomeranians and goldfish
// readings indicate that there are fewer than that many (due to
// the modial interaction of magnetoreluctance).
//
// What is the number of the real Aunt Sue?
func main() {
	detectedCompounds := createDetectedComopoundsMap()
	fileScanner := createFileScanner()
	input := parseInput(fileScanner)
	correctAuntSue := findCorrectSue(detectedCompounds, input)
	fmt.Printf("Correct Aunt Sue: %v \n", correctAuntSue)
}

func findCorrectSue(c map[string]string, input map[string]map[string]string) string {
	for currSue, v := range input {
		fmt.Printf("curr Sue: %v \n", currSue)
		numOfSimilarVals := 0
		for compound, compoundVal := range v {
			if compound == "cats" || compound == "trees" {
				if c[compound] < compoundVal {
					numOfSimilarVals++
				}
				continue
			}
			if compound == "pomeranians" || compound == "goldfish" {
				if c[compound] > compoundVal {
					numOfSimilarVals++
				}
				continue
			}
			if c[compound] == compoundVal {
				numOfSimilarVals++
			}
		}
		if numOfSimilarVals == 3 {
			return currSue
		}
	}
	return "Not Found"
}

func parseInput(fileScanner *bufio.Scanner) map[string]map[string]string {
	mp := make(map[string]map[string]string)
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		arr := createArr(currLine, n)          // create array out of data from current line
		currNum, mpOfThings := createList(arr) // list of things for current Aunt sue
		mp[currNum] = mpOfThings
	}
	return mp
}

func createList(arr []string) (string, map[string]string) {
	mp := make(map[string]string)
	currNum := arr[1]
	currNum = currNum[0 : len(currNum)-1]
	first := arr[2]
	first = first[0 : len(first)-1]
	firstValue := arr[3]
	firstValue = firstValue[0 : len(firstValue)-1]
	second := arr[4]
	second = second[0 : len(second)-1]
	secondValue := arr[5]
	secondValue = secondValue[0 : len(secondValue)-1]
	third := arr[6]
	third = third[0 : len(third)-1]
	thirdValue := arr[7]
	thirdValue = thirdValue[0 : len(thirdValue)-1]
	mp[first] = firstValue
	mp[second] = secondValue
	mp[third] = thirdValue
	return currNum, mp
}

func createArr(line string, n int) []string {
	arr := []string{}
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if j == n-1 {
				curr += string(line[j])
				arr = append(arr, curr)
				i = j
				break
			}
			if string(line[j]) == " " {
				arr = append(arr, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return arr
}

func createDetectedComopoundsMap() map[string]string {
	mp := map[string]string{
		"children":    "3",
		"cats":        "7",
		"samoyeds":    "2",
		"pomeranians": "3",
		"akitas":      "0",
		"vizslas:":    "0",
		"goldfish":    "5",
		"trees":       "3",
		"cars":        "2",
		"perfumes":    "1",
	}
	return mp
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
