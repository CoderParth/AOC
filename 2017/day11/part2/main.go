package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// --- Day 11: Hex Ed ---
// Crossing the bridge, you've barely reached the other side of the stream
// when a program comes up to you, clearly in distress. "It's my
// child process," she says, "he's gotten lost in an infinite grid!"
//
// Fortunately for her, you have plenty of experience with infinite grids.
//
// Unfortunately for you, it's a hex grid.
//
// The hexagons ("hexes") in this grid are aligned such that adjacent
// hexes can be found to the north, northeast, southeast, south, southwest, and northwest:
//
//	\ n  /
//
// nw +--+ ne
//
//	/    \
//
// -+      +-
//
//	\    /
//
// sw +--+ se
//
//	/ s  \
//
// You have the path the child process took. Starting where he started,
// you need to determine the fewest number of steps required to reach
// him. (A "step" means to move from the hex you are in to any adjacent hex.)
//
// For example:
//
// ne,ne,ne is 3 steps away.
// ne,ne,sw,sw is 0 steps away (back where you started).
// ne,ne,s,s is 2 steps away (se,se).
// se,sw,se,sw,sw is 3 `steps away (s,s,sw).

// •  North (n):
// (x,y,z)→(x,y+1,z−1)
// •  Northeast (ne):
// (x,y,z)→(x+1,y,z−1)
// •  Southeast (se):
// (x,y,z)→(x+1,y−1,z)
// •  South (s):
// (x,y,z)→(x,y−1,z+1)
// •  Southwest (sw):
// (x,y,z)→(x−1,y,z+1)
// •  Northwest (nw):
// (x,y,z)→(x−1,y+1,z)
func main() {
	furthestDistanceAway := findFurthestDistance()
	fmt.Printf("Furthest Distance: %v \n", furthestDistanceAway)
}

func findFurthestDistance() int {
	hexes := getHexesFromInput()
	x, y, z := 0, 0, 0
	furthestDistance := math.MinInt
	for _, hex := range hexes {
		switch hex {
		case "n":
			y++
			z--
			currDist := calculateDistance(x, y, z)
			furthestDistance = max(furthestDistance, currDist)
		case "ne":
			x++
			z--
			currDist := calculateDistance(x, y, z)
			furthestDistance = max(furthestDistance, currDist)
		case "se":
			x++
			y--
			currDist := calculateDistance(x, y, z)
			furthestDistance = max(furthestDistance, currDist)
		case "s":
			y--
			z++
			currDist := calculateDistance(x, y, z)
			furthestDistance = max(furthestDistance, currDist)
		case "sw":
			x--
			z++
			currDist := calculateDistance(x, y, z)
			furthestDistance = max(furthestDistance, currDist)
		case "nw":
			x--
			y++
			currDist := calculateDistance(x, y, z)
			furthestDistance = max(furthestDistance, currDist)
		}
	}
	return furthestDistance
}

func calculateDistance(x, y, z int) int {
	d := (abs(x) + abs(y) + abs(z)) / 2
	return d
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getHexesFromInput() []string {
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		arr := createArr(line)
		return arr
	}
	return []string{}
}

func createArr(line string) []string {
	arr := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == "," {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == "," {
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

func createFileScanner() *bufio.Scanner {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}
