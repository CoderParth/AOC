package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Day 7: Recursive Circus ---
// Wandering further through the circuits of the computer, you come
// upon a tower of programs that have gotten themselves into a bit
// of trouble. A recursive algorithm has gotten out of hand, and now
// they're balanced precariously in a large tower.
//
// One program at the bottom supports the entire tower. It's holding a
// large disc, and on the disc are balanced several more sub-towers.
// At the bottom of these sub-towers, standing on the bottom disc, are
// other programs, each holding their own disc, and so on. At the very
// tops of these sub-sub-sub-...-towers, many programs stand simply
// keeping the disc below them balanced but with no disc of their own.
//
// You offer to help, but first you need to understand the structure of
// these towers. You ask each program to yell out their name, their weight,
// and (if they're holding a disc) the names of the programs immediately
// above them balancing on that disc. You write this information down
// (your puzzle input). Unfortunately, in their panic, they don't do
// this in an orderly fashion; by the time you're done, you're not sure
// which program gave which information.
//
// For example, if your list is the following:
//
// pbga (66)
// xhth (57)
// ebii (61)
// havc (66)
// ktlj (57)
// fwft (72) -> ktlj, cntj, xhth
// qoyq (66)
// padx (45) -> pbga, havc, qoyq
// tknk (41) -> ugml, padx, fwft
// jptl (61)
// ugml (68) -> gyxo, ebii, jptl
// gyxo (61)
// cntj (57)
// ...then you would be able to recreate the structure of the towers that looks like this:
//
//	           gyxo
//	         /
//	    ugml - ebii
//	  /      \
//	 |         jptl
//	 |
//	 |         pbga
//	/        /
//
// tknk --- padx - havc
//
//	\        \
//	 |         qoyq
//	 |
//	 |         ktlj
//	  \      /
//	    fwft - cntj
//	         \
//	           xhth
//
// In this example, tknk is at the bottom of the tower (the bottom program),
// and is holding up ugml, padx, and fwft. Those programs are, in turn,
// holding up other programs; in this example, none of those programs are
// holding up any other programs, and are all the tops of their own towers.
// (The actual tower balancing in front of you is much larger.)
//
// Before you're ready to help them, you need to make sure your information
// is correct. What is the name of the bottom program?

type Disc struct {
	name   string
	weight string
}

func main() {
	fileScanner := createFileScanner()
	mp := make(map[string][]string) // map of the discs and their sub-towers
	for fileScanner.Scan() {
		disc, subTowers := parseInput(fileScanner.Text())
		mp[disc] = subTowers
	}
	fmt.Printf("Map of discs: %v \n", mp)
	bottomDisc := findBottomDisc(mp)
	fmt.Printf("Bottom Disc %v \n", bottomDisc)
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

func parseInput(line string) (string, []string) {
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
	weight := ""
	idx++
	for idx < n {
		if string(line[idx]) == ")" {
			idx++
			break
		}
		weight += string(line[idx])
		idx++
	}
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
	return name, subTowers
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
