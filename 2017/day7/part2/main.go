package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// The programs explain the situation: they can't get down. Rather,
// they could get down, if they weren't expending all of their energy
// trying to keep the tower balanced. Apparently, one program has the
// wrong weight, and until it's fixed, they're stuck here.
//
// For any program holding a disc, each program standing on that disc
// forms a sub-tower. Each of those sub-towers are supposed to be the
// same weight, or the disc itself isn't balanced. The weight of a
// tower is the sum of the weights of the programs in that tower.
//
// In the example above, this means that for ugml's disc to be balanced,
// gyxo, ebii, and jptl must all have the same weight, and they do: 61.
//
// However, for tknk to be balanced, each of the programs standing on
// its disc and all programs above it must each match. This means that
// the following sums must all be the same:
//
// ugml + (gyxo + ebii + jptl) = 68 + (61 + 61 + 61) = 251
// padx + (pbga + havc + qoyq) = 45 + (66 + 66 + 66) = 243
// fwft + (ktlj + cntj + xhth) = 72 + (57 + 57 + 57) = 243
//
// As you can see, tknk's disc is unbalanced: ugml's stack is heavier
// than the other two. Even though the nodes above ugml are balanced,
// ugml itself is too heavy: it needs to be 8 units lighter for its
// stack to weigh 243 and keep the towers balanced. If this change
// were made, its weight would be 60.
//
// Given that exactly one program is the wrong weight, what would its
// weight need to be to balance the entire tower?
type Disc struct {
	name   string
	weight string
}

func main() {
	fileScanner := createFileScanner()
	mp := make(map[string][]string) // map of the discs and their sub-towers
	weightMp := make(map[string]int)
	for fileScanner.Scan() {
		disc, weight, subTowers := parseInput(fileScanner.Text())
		mp[disc] = subTowers
		weightMp[disc] = weight
	}
	fmt.Printf("Map of discs: %v \n", mp)
	fmt.Printf("Map of weights: %v \n", weightMp)
	bottomDisc := findBottomDisc(mp)
	fmt.Printf("Bottom Disc %v \n", bottomDisc)
	visited := make(map[string]int)
	newWeightMp := make(map[string]int)
	recursivelyCalculateWeight(bottomDisc, mp, &visited, &weightMp, &newWeightMp)
	finalizeNewWeightMp(&newWeightMp, &weightMp)
	fmt.Printf("Map of new weights: %v \n", newWeightMp)
	balancingWeight := findBalancingWeight(newWeightMp, weightMp, bottomDisc)
	fmt.Printf("Balancing weight: %v \n", balancingWeight)
}

func findBalancingWeight(newWeightMp map[string]int, weightMp map[string]int, bottomDisc string) int {
	setOfDiscs := make(map[int]int)
	commonWeight := 0
	for disc, weight := range newWeightMp {
		if disc == bottomDisc {
			continue
		}
		setOfDiscs[weight]++
		if setOfDiscs[weight] > 1 {
			commonWeight = weight
		}
	}

	wrongDisc, wrongWeight := "", 0
	for disc, weight := range newWeightMp {
		if weight != commonWeight {
			wrongWeight = weight
			wrongDisc = disc
		}
	}

	fmt.Printf("w weight: %v \n", wrongWeight)
	fmt.Printf("c weight: %v \n", commonWeight)
	fmt.Printf(" weight mp wrong disc: %v \n", weightMp[wrongDisc])
	balancingWeight := weightMp[wrongDisc] - (wrongWeight - commonWeight)
	return balancingWeight
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

func parseInput(line string) (string, int, []string) {
	n, idx := len(line), 0
	name := ""
	// parse name
	for idx < n {
		if string(line[idx]) == " " {
			idx++
			break
		}
		name += string(line[idx])
		idx++
	}
	// parse weight
	weightStr := ""
	idx++
	for idx < n {
		if string(line[idx]) == ")" {
			idx++
			break
		}
		weightStr += string(line[idx])
		idx++
	}
	weightInt := convStrToInt(weightStr)
	// find subTowers
	idx += 4
	subTowers := []string{}
	for idx < n {
		curr := ""
		for j := idx; j < n; j++ {
			if string(line[j]) == "," {
				subTowers = append(subTowers, curr)
				idx = j + 2
				break
			}
			if j == n-1 {
				curr += string(line[j])
				subTowers = append(subTowers, curr)
				idx = j + 1
				break
			}
			curr += string(line[j])
		}
	}
	return name, weightInt, subTowers
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func findBottomDisc(mp map[string][]string) string {
	for disc := range mp {
		visited := make(map[string]int)
		searchRecusively(disc, mp, &visited)
		if len(visited) == len(mp) {
			return disc
		}
	}
	return "Not found"
}

func searchRecusively(disc string, mp map[string][]string, visited *map[string]int) {
	if _, ok := (*visited)[disc]; ok {
		return
	}
	(*visited)[disc] = 0
	for _, subTower := range mp[disc] {
		fmt.Printf("Sub Tower: %v \n", subTower)
		if _, ok := (*visited)[subTower]; !ok {
			searchRecusively(subTower, mp, visited)
		}
	}
}

func recursivelyCalculateWeight(disc string, mp map[string][]string, visited *map[string]int, weightMp *map[string]int, newWeightMp *map[string]int) {
	if _, ok := (*visited)[disc]; ok {
		return
	}
	(*visited)[disc] = 0
	for _, subTower := range mp[disc] {
		fmt.Printf("Sub Tower: %v \n", subTower)
		if _, ok := (*visited)[subTower]; !ok {
			(*newWeightMp)[disc] += (*weightMp)[subTower]
			recursivelyCalculateWeight(subTower, mp, visited, weightMp, newWeightMp)
		}
	}
}

func finalizeNewWeightMp(newWeightMp *map[string]int, weightMp *map[string]int) {
	for disc := range *newWeightMp {
		(*newWeightMp)[disc] += (*weightMp)[disc]
	}
}
