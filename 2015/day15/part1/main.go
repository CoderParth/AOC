package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// For instance, suppose you have these two ingredients:
//
// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
// Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
// Then, choosing to use 44 teaspoons of butterscotch and 56 teaspoons of
// cinnamon (because the amounts of each ingredient must add up to 100)
// would result in a cookie with the following properties:
//
// A capacity of 44*-1 + 56*2 = 68
// A durability of 44*-2 + 56*3 = 80
// A flavor of 44*6 + 56*-2 = 152
// A texture of 44*3 + 56*-1 = 76
// Multiplying these together (68 * 80 * 152 * 76, ignoring calories for now)
// results in a total score of 62842880, which happens to be the best score
// possible given these ingredients. If any properties had produced a negative
// total, it would have instead become zero, causing the whole score to multiply
// to zero.
//
// Given the ingredients in your kitchen and their properties, what is the total
// score of the highest-scoring cookie you can make?

type Properties struct {
	capacity   int
	durability int
	flavor     int
	texture    int
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
		score := calculateScore(input, arr)
		if score > *highestScore {
			*highestScore = score
		}
		return
	}
	for i := 0; i <= remaining; i++ {
		arr[index] = i
		generateCombinations(arr, n, index+1, remaining-i, highestScore, input)
	}
}

func calculateScore(input map[string]Properties, arr []int) int {
	c, d, f, t := 0, 0, 0, 0 // capacity, durability, flavor, texture
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

	}
	if c <= 0 || d <= 0 || f <= 0 || t <= 0 {
		return 0
	}
	totalScore := c * d * f * t
	return totalScore
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
	p := Properties{
		capacity,
		durability,
		flavor,
		texture,
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
