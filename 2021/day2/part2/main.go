package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// Based on your calculations, the planned course doesn't seem to make any sense. You find the submarine manual and discover that the process is actually slightly more complicated.
//
// In addition to horizontal position and depth, you'll also need to track a third value, aim, which also starts at 0. The commands also mean something entirely different than you first thought:
//
// down X increases your aim by X units.
// up X decreases your aim by X units.
// forward X does two things:
// It increases your horizontal position by X units.
// It increases your depth by your aim multiplied by X.
// Again note that since you're on a submarine, down and up do the opposite of what you might expect: "down" means aiming in the positive direction.
//
// Now, the above example does something different:
//
// forward 5 adds 5 to your horizontal position, a total of 5. Because your aim is 0, your depth does not change.
// down 5 adds 5 to your aim, resulting in a value of 5.
// forward 8 adds 8 to your horizontal position, a total of 13. Because your aim is 5, your depth increases by 8*5=40.
// up 3 decreases your aim by 3, resulting in a value of 2.
// down 8 adds 8 to your aim, resulting in a value of 10.
// forward 2 adds 2 to your horizontal position, a total of 15. Because your aim is 10, your depth increases by 2*10=20 to a total of 60.
// After following these new instructions, you would have a horizontal position of 15 and a depth of 60. (Multiplying these produces 900.)
//
// Using this new interpretation of the commands, calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?
func main() {
	fileScanner := createFileScanner()
	horPos, depth := calculatePos(fileScanner)
	fmt.Printf("horizontal Position: %v \n", horPos)
	fmt.Printf("Depth: %v \n", depth)
	product := horPos * depth
	fmt.Printf("Product: %v \n", product)
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

func calculatePos(fileScanner *bufio.Scanner) (int, int) {
	horPos, depth, aim := 0, 0, 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		command, units := parseLine(currLine)
		switch command {
		case "forward":
			horPos += units
			depth += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	return horPos, depth
}

func parseLine(line string) (string, int) {
	n := len(line)
	i := 0
	command := ""
	for ; i < n; i++ {
		if string(line[i]) == " " {
			i++
			break
		}
		command += string(line[i])
	}
	unitsInStr := ""
	for ; i < n; i++ {
		unitsInStr += string(line[i])
	}
	units := convStrToInt(unitsInStr)
	return command, units
}

func convStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
