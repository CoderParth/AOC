package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// You have 500 Aunts named "Sue".
//
// So, to avoid sending the card to the wrong person, you need to figure
// out which Aunt Sue (which you conveniently number 1 to 500, for sanity)
// gave you the gift. You open the present and, as luck would have it,
// good ol' Aunt Sue got you a My First Crime Scene Analysis Machine!
// Just what you wanted. Or needed, as the case may be.
//
// The My First Crime Scene Analysis Machine (MFCSAM for short) can detect a
// few specific compounds in a given sample, as well as how many distinct
// kinds of those compounds there are. According to the instructions,
// these are what the MFCSAM can detect:
//
// children, by human DNA age analysis.
// cats. It doesn't differentiate individual breeds.
// Several seemingly random breeds of dog: samoyeds, pomeranians, akitas,
// and vizslas.
// goldfish. No other kinds of fish.
// trees, all in one group.
// cars, presumably by exhaust or gasoline or something.
// perfumes, which is handy, since many of your Aunts Sue wear a few kinds.
// In fact, many of your Aunts Sue have many of these. You put the wrapping
// from the gift into the MFCSAM. It beeps inquisitively at you a few times
// and then prints out a message on ticker tape:
//
// children: 3
// cats: 7
// samoyeds: 2
// pomeranians: 3
// akitas: 0
// vizslas: 0
// goldfish: 5
// trees: 3
// cars: 2
// perfumes: 1
//
// You make a list of the things you can remember about each Aunt Sue.
// Things missing from your list aren't zero - you simply don't remember the value.
//
// What is the number of the Sue that got you the gift?

//	type Compounds struct {
//		children    string
//		cats        string
//		samoyeds    string
//		pomeranians string
//		akitas      string
//		vizslas     string
//		goldfish    string
//		trees       string
//		cars        string
//		perfumes    string
//	}
func main() {
	detectedCompounds := createDetectedComopoundsMap()
	fileScanner := createFileScanner()
	input := parseInput(fileScanner)
	correctAuntSue := findCorrectSue(detectedCompounds, input)
	fmt.Printf("Input: %v \n", input)
	fmt.Printf("Correct Aunt Sue: %v \n", correctAuntSue)
}

func findCorrectSue(c map[string]string, input map[string]map[string]string) string {
	for currSue, v := range input {
		numOfSimilarVals := 0
		for compound, compoundVal := range v {
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
		"	children": "3",
		"cats:   ":  "7",
		"samoyeds":  "2",
		"pomerani":  "3",
		"akitas: ":  "0",
		"vizslas:":  "0",
		"goldfish":  "5",
		"trees:  ":  "3",
		"cars:   ":  "2",
		"perfumes":  "1",
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
