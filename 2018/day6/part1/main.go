package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// --- Day 6: Chronal Coordinates ---
// The device on your wrist beeps several times, and once again you feel
// like you're falling.
//
// "Situation critical," the device announces. "Destination indeterminate.
// Chronal interference detected. Please specify new target coordinates."
//
// The device then produces a list of coordinates (your puzzle input). Are
// they places it thinks are safe or dangerous? It recommends you check
// manual page 729. The Elves did not give you a manual.
//
// If they're dangerous, maybe you can minimize the danger by finding the
// coordinate that gives the largest distance from the other points.
//
// Using only the Manhattan distance, determine the area around each
// coordinate by counting the number of integer X,Y locations that are
// closest to that coordinate (and aren't tied in distance to any other
// coordinate).
//
// Your goal is to find the size of the largest area that isn't infinite.
// For example, consider the following list of coordinates:
//
// 1, 1
// 1, 6
// 8, 3
// 3, 4
// 5, 5
// 8, 9
// If we name these coordinates A through F, we can draw them on a grid,
// putting 0,0 at the top left:
//
// ..........
// .A........
// ..........
// ........C.
// ...D......
// .....E....
// .B........
// ..........
// ..........
// ........F.
// This view is partial - the actual grid extends infinitely in all directions.
// Using the Manhattan distance, each location's closest coordinate can be
// determined, shown here in lowercase:
//
// aaaaa.cccc
// aAaaa.cccc
// aaaddecccc
// aadddeccCc
// ..dDdeeccc
// bb.deEeecc
// bBb.eeee..
// bbb.eeefff
// bbb.eeffff
// bbb.ffffFf
// Locations shown as . are equally far from two or more coordinates, and
// so they don't count as being closest to any.
//
// In this example, the areas of coordinates A, B, C, and F are infinite -
// while not shown here, their areas extend forever outside the visible grid.
// However, the areas of coordinates D and E are finite: D is closest to 9
// locations, and E is closest to 17 (both including the coordinate's location
// itself). Therefore, in this example, the size of the largest area is 17.
//
// What is the size of the largest area that isn't infinite?
func main() {
	fileScanner := createFileScanner()
	input := parseInput(fileScanner)
	largestCol, largestRow := findLargestColAndRow(input)
	fmt.Printf("Col: %v \n Row: %v \n", largestCol, largestRow)
	grid := initializeGrid(largestCol, largestRow)
	calculateDistanceAndMarkCord(&grid, input)
	for _, i := range grid {
		fmt.Printf("- :%v \n", i)
	}
	largestArea := findLargestArea(grid)
	fmt.Printf("Largest Area: %v \n", largestArea)
}

type Pair struct {
	c1 int
	c2 int
}

func findLargestArea(grid [][][3]int) int {
	mp := make(map[Pair]int)
	gridRowLen := len(grid)
	gridColLen := len(grid[0])
	for i := 0; i < gridRowLen; i++ {
		for j := 0; j < gridColLen; j++ {
			c1 := grid[i][j][1]
			c2 := grid[i][j][2]
			p := Pair{
				c1,
				c2,
			}
			mp[p]++
		}
	}

	largest := math.MinInt
	for _, v := range mp {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func calculateDistanceAndMarkCord(grid *[][][3]int, input [][]int) {
	m := len(input)
	gridRowLen := len(*grid)
	gridColLen := len((*grid)[0])
	for i := 0; i < m; i++ {
		currCol, currRow := input[i][0], input[i][1]
		for x := 0; x < gridRowLen; x++ {
			for y := 0; y < gridColLen; y++ {
				if x == currRow && y == currCol {
					continue
				}
				currDist := findManhattanDist(currCol, currRow, x, y)
				if currDist < (*grid)[x][y][0] {
					(*grid)[x][y][0] = currDist
					(*grid)[x][y][1] = currCol
					(*grid)[x][y][2] = currRow
				}
			}
		}
	}
}

func findManhattanDist(x1, y1, x2, y2 int) int {
	dist := abs(x1-x2) + abs(y1-y2)
	return dist
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func initializeGrid(largestCol, largestRow int) [][][3]int {
	grid := make([][][3]int, largestRow+1)
	for i := 0; i < largestRow+1; i++ {
		grid[i] = make([][3]int, largestCol+1)
	}
	for i := 0; i < largestRow+1; i++ {
		for j := 0; j < largestCol+1; j++ {
			grid[i][j][0] = math.MaxInt
		}
	}
	return grid
}

func findLargestColAndRow(input [][]int) (int, int) {
	largestCol, largestRow := math.MinInt, math.MinInt
	n := len(input)
	for i := 0; i < n; i++ {
		if input[i][0] > largestCol {
			largestCol = input[i][0]
		}
		if input[i][1] > largestRow {
			largestRow = input[i][1]
		}
	}
	return largestCol, largestRow
}

func parseInput(fileScanner *bufio.Scanner) [][]int {
	rowsAndCols := [][]int{}
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		col, row := parse(currLine)
		rowsAndCols = append(rowsAndCols, []int{col, row})
	}
	return rowsAndCols
}

func parse(line string) (int, int) {
	arr := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
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
	col, row := convStrToInt(arr[0]), convStrToInt(arr[1])
	return col, row
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
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
