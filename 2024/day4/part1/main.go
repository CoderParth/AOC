package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Day 4: Ceres Search ---
// "Looks like the Chief's not here. Next!" One of The Historians pulls
// out a device and pushes the only button on it. After a brief flash,
// you recognize the interior of the Ceres monitoring station!
//
// As the search for the Chief continues, a small Elf who lives on the
// station tugs on your shirt; she'd like to know if you could help
// her with her word search (your puzzle input). She only has to
// find one word: XMAS.
//
// This word search allows words to be horizontal, vertical, diagonal,
// written backwards, or even overlapping other words. It's a little
// unusual, though, as you don't merely need to find one instance of
// XMAS - you need to find all of them. Here are a few ways XMAS
// might appear, where irrelevant characters have been replaced with .:
//
// ..X...
// .SAMX.
// .A..A.
// XMAS.S
// .X....
// The actual word search will be full of letters instead. For example:
//
// MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX
// In this word search, XMAS occurs a total of 18 times; here's the
// same word search again, but where letters not involved in any
// XMAS have been replaced with .:
//
// ....XXMAS.
// .SAMXMS...
// ...S..A...
// ..A.A.MS.X
// XMASAMX.MM
// X.....XA.A
// S.S.S.S.SS
// .A.A.A.A.A
// ..M.M.M.MM
// .X.X.XMASX
// Take a look at the little Elf's word search. How many times does XMAS appear?

func main() {
	// NOTE: I tried using DFS, it did not work for some reason.
	// Perhaps, my implementation was wrong.
	// So, instead, I applied brute-force kind of solution, which gives correct answer.
	fs := fileScanner()
	grid := [][]byte{}
	for fs.Scan() {
		currArr := parseLine(fs.Text())
		grid = append(grid, currArr)
	}
	total := search(grid)
	fmt.Printf("Total Occurence: %d \n", total)
}

func search(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	word := "XMAS"
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if string(grid[i][j]) == "X" {
				total += searchHorizontally(&grid, word, i, j, n)
				total += searchVertically(&grid, word, i, j, m)
				total += searchTopDiagonals(&grid, word, i, j, m)
				total += searchBottomDiagonals(&grid, word, i, j, m, n)
			}
		}
	}
	return total
}

func searchHorizontally(grid *[][]byte, word string, i, j, n int) int {
	total := 0
	curr := "X"
	// horizontally right
	for k := j + 1; k < j+4; k++ {
		if k >= n {
			break
		}
		curr += string((*grid)[i][k])
	}
	if curr == word {
		total++
	}

	curr = "X"
	// horizontally left
	for k := j - 1; k > j-4; k-- {
		if k < 0 {
			break
		}
		curr += string((*grid)[i][k])
	}
	if curr == word {
		total++
	}
	return total
}

func searchVertically(grid *[][]byte, word string, i, j, m int) int {
	total := 0
	curr := "X"
	// Vertically down
	for k := i + 1; k < i+4; k++ {
		if k >= m {
			break
		}
		curr += string((*grid)[k][j])
	}
	if curr == word {
		total++
	}
	// Vertically Upwards
	curr = "X"
	for k := i - 1; k > i-4; k-- {
		if k < 0 {
			break
		}
		curr += string((*grid)[k][j])
	}
	if curr == word {
		total++
	}

	return total
}

func searchTopDiagonals(grid *[][]byte, word string, i, j, m int) int {
	total := 0
	curr := "X"
	// top left diagonal
	k, l := j-1, i-1
	times := 0
	for k >= 0 && l >= 0 && times < 3 {
		curr += string((*grid)[l][k])
		l--
		k--
		times++
	}
	if curr == word {
		total++
	}

	// top right diagonal
	curr = "X"
	k, l = j+1, i-1
	times = 0
	for k < m && l >= 0 && times < 3 {
		curr += string((*grid)[l][k])
		l--
		k++
		times++
	}
	if curr == word {
		total++
	}
	return total
}

func searchBottomDiagonals(grid *[][]byte, word string, i, j, m, n int) int {
	total := 0
	curr := "X"
	// bottom left diagonal
	k, l := j-1, i+1
	times := 0
	for k >= 0 && l < m && times < 3 {
		curr += string((*grid)[l][k])
		l++
		k--
		times++
	}
	if curr == word {
		total++
	}

	// bottom right diagonal
	curr = "X"
	k, l = j+1, i+1
	times = 0
	for k < n && l < m && times < 3 {
		curr += string((*grid)[l][k])
		l++
		k++
		times++
	}
	if curr == word {
		total++
	}
	return total
}

func parseLine(line string) []byte {
	n := len(line)
	currArr := []byte{}
	for i := 0; i < n; i++ {
		currArr = append(currArr, line[i])
	}
	return currArr
}

func fileScanner() *bufio.Scanner {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)
	return fs
}
