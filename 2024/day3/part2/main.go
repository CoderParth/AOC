package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --- Part Two ---
// As you scan through the corrupted memory, you notice that some of the
// conditional statements are also still intact. If you handle some of
// the uncorrupted conditional statements in the program, you might be
// able to get an even more accurate result.
//
// There are two new instructions you'll need to handle:
//
// The do() instruction enables future mul instructions.
// The don't() instruction disables future mul instructions.
// Only the most recent do() or don't() instruction applies. At the
// beginning of the program, mul instructions are enabled.
//
// For example:
//
// xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
// This corrupted memory is similar to the example from before, but this
// time the mul(5,5) and mul(11,8) instructions are disabled because
// there is a don't() instruction before them. The other mul
// instructions function normally, including the one at the end that gets
// re-enabled by a do() instruction.
//
// This time, the sum of the results is 48 (2*4 + 8*5).
//
// Handle the new instructions; what do you get if you add up all of
// the results of just the enabled multiplications?

func main() {
	fs := fileScanner()
	total := 0
	pastIns := "do()" // past instruction - either do() or don't()
	for fs.Scan() {
		currLine := fs.Text()
		total += parseAndEval(currLine, &pastIns)
	}
	fmt.Printf("Answer is: %d \n", total)
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

func parseAndEval(line string, pastIns *string) int {
	n, total := len(line), 0
	for i := 0; i < n; i++ {
		if string(line[i]) != "m" && string(line[i]) != "d" {
			continue
		}

		if string(line[i]) == "d" {
			if doExists(&i, &line) {
				*pastIns = "do()"
				continue
			}

			if dontExists(&i, &line) {
				*pastIns = "don't()"
			}
			continue
		}

		if string(line[i]) == "m" && (*pastIns == "do()") {
			curr := ""
			j := i
			for ; j < i+4; j++ {
				curr += string(line[j])
			}
			if curr != "mul(" {
				continue
			}

			k := j
			firstStr := ""
			for ; k < j+6; k++ { // max digit of a num is 3 in input
				if string(line[k]) == "," {
					break
				}
				firstStr += string(line[k])
			}
			firstNum, err := strconv.Atoi(firstStr)
			if err != nil {
				i = j
				continue
			}

			l := k + 1
			secondStr := ""
			for ; l < k+6; l++ { // max digit of a num is 3 in input
				if string(line[l]) == ")" {
					break
				}
				secondStr += string(line[l])
			}
			secondNum, err := strconv.Atoi(secondStr)
			if err != nil {
				i = j
				continue
			}

			total += (firstNum * secondNum)
			i = l
		}
	}
	return total
}

func doExists(i *int, line *string) bool {
	curr := ""
	for j := *i; j < (*i)+4; j++ {
		curr += string((*line)[j])
	}
	return curr == "do()"
}

func dontExists(i *int, line *string) bool {
	curr := ""
	for k := *i; k < (*i)+7; k++ {
		curr += string((*line)[k])
	}
	return curr == "don't()"
}
