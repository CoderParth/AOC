package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// --- Part Two ---
// The engineers are surprised by the low number of safe reports until
// they realize they forgot to tell you about the Problem Dampener.
//
// The Problem Dampener is a reactor-mounted module that lets the
// reactor safety systems tolerate a single bad level in what would
// otherwise be a safe report. It's like the bad level never happened!
//
// Now, the same rules apply as before, except if removing a single
// level from an unsafe report would make it safe, the report instead
// counts as safe.
//
// More of the above example's reports are now safe:
//
// 7 6 4 2 1: Safe without removing any level.
// 1 2 7 8 9: Unsafe regardless of which level is removed.
// 9 7 6 2 1: Unsafe regardless of which level is removed.
// 1 3 2 4 5: Safe by removing the second level, 3.
// 8 6 4 4 1: Safe by removing the third level, 4.
// 1 3 6 7 9: Safe without removing any level.
// Thanks to the Problem Dampener, 4 reports are actually safe!
//
// Update your analysis by handling situations where the Problem Dampener
// can remove a single level from unsafe reports. How many reports are
// now safe?

func main() {
	fs := fileScanner()
	safeReports := 0
	for fs.Scan() {
		nums := parse(fs.Text())
		if isSafe(nums) {
			safeReports++
			continue
		}
	}
	fmt.Printf("Total safe Reports: %d \n", safeReports)
}

func fileScanner() *bufio.Scanner {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func parse(line string) []int {
	arr := []int{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if j == n-1 {
				curr += string(line[j])
				num := convStrToInt(curr)
				arr = append(arr, num)
				i = j
				break
			}
			if string(line[j]) == " " {
				num := convStrToInt(curr)
				arr = append(arr, num)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return arr
}

func convStrToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func isSafe(nums []int) bool {
	return isIncreasing(nums) || isDecreasing(nums)
}

func isIncreasing(nums []int) bool {
	if nums[0] > nums[1] {
		return false
	}
	if nums[0] == nums[1] {
		if nums[2]-nums[1] >= 3 {
			return false
		}
	}
	curr := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		diff := nums[i] - curr
		if curr < nums[i] {
			if diff <= 3 {
				curr = nums[i]
				continue
			}
		}
		if abs(diff) <= 1 {
			curr = nums[i]
			continue
		}
		return false // nums are not increasing
	}
	return true
}

func isDecreasing(nums []int) bool {
	curr := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		diff := curr - nums[i]
		if curr > nums[i] {
			if diff <= 3 {
				curr = nums[i]
				continue
			}
		}
		if abs(diff) <= 1 {
			curr = nums[i]
			continue
		}
		return false // nums are not decreasing
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
