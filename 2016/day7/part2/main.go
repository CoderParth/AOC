package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// You would also like to know which IPs support SSL (super-secret
// listening).
//
// An IP supports SSL if it has an Area-Broadcast Accessor, or ABA,
// anywhere in the supernet sequences (outside any square bracketed
// sections), and a corresponding Byte Allocation Block, or BAB,
// anywhere in the hypernet sequences. An ABA is any three-character
// sequence which consists of the same character twice with a different
// character between them, such as xyx or aba. A corresponding BAB is
// the same characters but in reversed positions: yxy and bab, respectively.
//
// For example:
//
// aba[bab]xyz supports SSL (aba outside square brackets with
// corresponding bab within square brackets).
// xyx[xyx]xyx does not support SSL (xyx, but no corresponding yxy).
// aaa[kek]eke supports SSL (eke in supernet with corresponding kek
// in hypernet; the aaa sequence is not related, because the interior
// character must be different).
// zazbz[bzb]cdb supports SSL (zaz has no corresponding aza, but zbz
// has a corresponding bzb, even though zaz and zbz overlap).
// How many IPs in your puzzle input support SSL?
func main() {
	total := findTlsSupportingIps()
	fmt.Printf("Total: %v \n", total)
}

func findTlsSupportingIps() int {
	fileScanner := createFileScanner()
	total := 0
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		input := parseInput(currLine)
		if supportsTls(input) {
			total++
		}
	}
	return total
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

func parseInput(line string) []string {
	n := len(line)
	arr := []string{}
	for i := 0; i < n; i++ {
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == "[" || string(line[j]) == "]" {
				if string(line[j]) == "]" {
					temp := "["
					temp += curr
					curr = temp
				}
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

func supportsTls(input []string) bool {
	supports := false
	for _, seq := range input {
		if string(seq[0]) == "[" {
			if hasAbba(seq[1:]) { // for seq inside square brackets
				return false
			}
			continue
		}
		if hasAbba(seq) { // for seq outside square brackets
			supports = true
		}
	}
	return supports
}

func hasAbba(seq string) bool {
	n := len(seq)
	for i := 0; i < n-3; i++ {
		leftSeq := string(seq[i]) + string(seq[i+1])
		rightSeq := string(seq[i+2]) + string(seq[i+3])
		if leftSeq == rightSeq {
			continue
		}
		reversedSeq := string(seq[i+3]) + string(seq[i+2])
		if leftSeq == reversedSeq {
			return true
		}
	}
	return false
}
