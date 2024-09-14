package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Day 3: Crossed Wires ---
// The gravity assist was successful, and you're well on your way to the
// Venus refuelling station. During the rush back on Earth, the fuel
// management system wasn't completely installed, so that's next on the
// priority list.
//
// Opening the front panel reveals a jumble of wires. Specifically, two
// wires are connected to a central port and extend outward on a grid.
// You trace the path each wire takes as it leaves the central port, one
// wire per line of text (your puzzle input).
//
// The wires twist and turn, but the two wires occasionally cross paths. To
// fix the circuit, you need to find the intersection point closest to the
// central port. Because the wires are on a grid, use the Manhattan distance
// for this measurement. While the wires do technically cross right at the
// central port where they both start, this point does not count, nor does
// a wire count as crossing with itself.
//
// For example, if the first wire's path is R8,U5,L5,D3, then starting from
// the central port (o), it goes right 8, up 5, left 5, and finally down 3:
//
// ...........
// ...........
// ...........
// ....+----+.
// ....|....|.
// ....|....|.
// ....|....|.
// .........|.
// .o-------+.
// ...........
// Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:
//
// ...........
// .+-----+...
// .|.....|...
// .|..+--X-+.
// .|..|..|.|.
// .|.-X--+.|.
// .|..|....|.
// .|.......|.
// .o-------+.
// ...........
// These wires cross at two locations (marked X), but the lower-left
// one is closer to the central port: its distance is 3 + 3 = 6.
//
// Here are a few more examples:
//
// R75,D30,R83,U83,L12,D49,R71,U7,L72
// U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
// R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
// U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135
// What is the Manhattan distance from the central port to the closest intersection?

type Coords struct {
	x int
	y int
}

type FloatCoords struct {
	x float64
	y float64
}

func main() {
	distance := findClosestIntersection()
	fmt.Printf("Distance: %v \n", distance)
}

func findClosestIntersection() int {
	fileScanner := createFileScanner()
	currLine := ""
	for fileScanner.Scan() {
		currLine += fileScanner.Text()
	}
	input := parse(currLine)
	// fmt.Printf("Input: %v \n", input)
	mp := moveAndMark(input)
	// fmt.Printf("MP: %v \n", mp)
	arr := createArrOfIntersection(mp)
	// fmt.Printf("Arr: %v \n", arr)
	intersectionPoints := findIntersectionPoints(arr)
	fmt.Printf("intersection points: %v \n", intersectionPoints)
	closestDist := findClosestDist()
	return closestDist
}

func createArrOfIntersection(mp map[Coords]int) [][]int {
	arr := [][]int{}
	for c := range mp {
		arr = append(arr, []int{c.x, c.y})
	}
	return arr
}

func findIntersectionPoints(arr [][]int) map[FloatCoords]int {
	intersections := make(map[FloatCoords]int)
	n := len(arr)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				for l := k + 1; l < n; l++ {
					x1, y1 := arr[i][0], arr[i][1]
					x2, y2 := arr[j][0], arr[j][1]
					x3, y3 := arr[k][0], arr[k][1]
					x4, y4 := arr[l][0], arr[l][1]
					m1, m2 := calcSlope(float64(x1), float64(x2), float64(y1), float64(y2)), calcSlope(float64(x3), float64(x4), float64(y3), float64(y4))
					b1, b2 := findYIntercept(float64(x1), float64(y1), float64(m1)), findYIntercept(float64(x1), float64(y1), float64(m2))
					// y = m1x + b1
					// y = m2x + b2
					// m1x + b1 = m2x + b2
					// m1x - m2x = b2 - b1
					// x = b2 - b1 / m1 - m2
					x := (b2 - b1) / (m1 - m2)
					y := m1*x + b1
					c := FloatCoords{
						x,
						y,
					}
					intersections[c] = 0
				}
			}
		}
	}

	return intersections
}

func calcSlope(xa, xb, ya, yb float64) float64 {
	return (yb - ya) / (xb - xa)
}

func findYIntercept(xa, ya, ma float64) float64 {
	ba := ya - (ma * xa)
	return ba
}

func moveAndMark(input []string) map[Coords]int {
	mp := make(map[Coords]int)
	x, y, n := 0, 0, len(input)
	for i := 0; i < n; i++ {
		if string(input[i][0]) == "X" {
			x = 0
			y = 0
			continue
		}
		currDir := string(input[i][0])
		currDist := convStrToInt(string(input[i][1:len(input[i])]))
		switch currDir {
		case "U":
			y += currDist
			c := Coords{
				x,
				y,
			}
			mp[c] = 0
		case "D":
			y -= currDist
			c := Coords{
				x,
				y,
			}
			mp[c] = 0
		case "L":
			x -= currDist
			c := Coords{
				x,
				y,
			}
			mp[c] = 0
		case "R":
			x += currDist
			c := Coords{
				x,
				y,
			}
			mp[c] = 0
		}
	}
	return mp
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func findClosestDist() int {
	return 0
}

func parse(line string) []string {
	n := len(line)
	arr := []string{}
	for i := 0; i < n; i++ {
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
