package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// Your cookie recipe becomes wildly popular! Someone asks if you can
// make another recipe that has exactly 500 calories per cookie (so
// they can use it as a meal replacement). Keep the rest of your
// award-winning process the same (100 teaspoons, same ingredients,
// same scoring system).
//
// For example, given the ingredients above, if you had instead
// selected 40 teaspoons of butterscotch and 60 teaspoons of cinnamon
// (which still adds to 100), the total calorie count would be 40*8
// + 60*3 = 500. The total score would go down, though: only 57600000,
// the best you can do in such trying circumstances.
//
// Given the ingredients in your kitchen and their properties, what is
// the total score of the highest-scoring cookie you can make with a
// calorie total of 500?

type Properties struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calorie    int
}

func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner)
	n := len(input) // the total number of ingredients
	maxAmount := 100
	arr := make([]int, n)
	highestScore := 0
	generateCombinations(arr, n, 0, maxAmount, &highestScore, input)
	fmt.Printf("Highest Score: %v \n", highestScore)
}

func generateCombinations(arr []int, n, index, remaining int, highestScore *int, input map[string]Properties) {
	if index == n-1 {
		arr[index] = remaining
		fmt.Printf("Arr: %v \n", arr)
		score, cal := calculateScore(input, arr)
		if score > *highestScore && cal == 500 {
			*highestScore = score
		}
		return
	}
	for i := 0; i <= remaining; i++ {
		arr[index] = i
		generateCombinations(arr, n, index+1, remaining-i, highestScore, input)
	}
}

func calculateScore(input map[string]Properties, arr []int) (int, int) {
	c, d, f, t, cal := 0, 0, 0, 0, 0 // capacity, durability, flavor, texture, calorie
	set := []string{}
	for k := range input {
		set = append(set, k)
	}
	for i := 0; i < len(arr); i++ {
		v := input[set[i]]
		c += v.capacity * arr[i]
		d += v.durability * arr[i]
		f += v.flavor * arr[i]
		t += v.texture * arr[i]
		cal += v.calorie * arr[i]
	}
	if c <= 0 || d <= 0 || f <= 0 || t <= 0 || cal == 0 {
		return 0, 0
	}
	totalScore := c * d * f * t
	return totalScore, cal
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

func parseInput(fileScanner *bufio.Scanner) map[string]Properties {
	mp := make(map[string]Properties)
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		n := len(currLine)
		arr := createArray(currLine, n) // create array from the given line
		p := createPropertiesStruct(arr)
		mp[arr[0]] = p
	}
	return mp
}

func createPropertiesStruct(arr []string) Properties {
	capacity := convStrToInt(arr[2])
	durability := convStrToInt(arr[4])
	flavor := convStrToInt(arr[6])
	texture := convStrToInt(arr[8])
	calorie := convStrToInt(arr[10])
	p := Properties{
		capacity,
		durability,
		flavor,
		texture,
		calorie,
	}
	return p
}

func convStrToInt(n string) int {
	length := len(n)
	n = n[0 : length-1]
	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func createArray(line string, n int) []string {
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
