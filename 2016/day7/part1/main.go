package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Day 7: Internet Protocol Version 7 ---
// While snooping around the local network of EBHQ, you compile a
// list of IP addresses (they're IPv7, of course; IPv6 is much
// too limited). You'd like to figure out which IPs support TLS
// (transport-layer snooping).
//
// An IP supports TLS if it has an Autonomous Bridge Bypass Annotation,
// or ABBA. An ABBA is any four-character sequence which consists of a
// pair of two different characters followed by the reverse of that pair,
// such as xyyx or abba. However, the IP also must not have an ABBA
// within any hypernet sequences, which are contained by square brackets.
//
// For example:
//
// abba[mnop]qrst supports TLS (abba outside square brackets).
// abcd[bddb]xyyx does not support TLS (bddb is within square brackets,
// even though xyyx is outside square brackets).
// aaaa[qwer]tyui does not support TLS (aaaa is invalid; the interior
// characters must be different).
// ioxxoj[asdfgh]zxcvbn supports TLS (oxxo is outside square brackets,
// even though it's within a larger string).
// How many IPs in your puzzle input support TLS?
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
		fmt.Printf("Input: %v \n", input)
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
	for _, seq := range input {
		fmt.Printf("curr seq: %v \n", seq)
		if string(seq[0]) == "[" {
			if hasAbba(seq[1:]) { // for seq inside square brackets
				return false
			}
			continue
		}
		if !hasAbba(seq) { // for seq outside square brackets
			return false
		}
	}
	fmt.Printf("Returning true")
	return true
}

func hasAbba(seq string) bool {
	fmt.Printf("curr seq: %v \n", seq)
	n := len(seq)
	for i := 0; i < n-3; i++ {
		leftSeq := string(seq[i]) + string(seq[i+1])
		rightSeq := string(seq[i+2]) + string(seq[i+3])
		fmt.Printf("left seq: %v \n", leftSeq)
		fmt.Printf("right seq: %v \n", rightSeq)
		if leftSeq == rightSeq {
			return false
		}
		reversedSeq := string(seq[i+3]) + string(seq[i+2])
		if leftSeq == reversedSeq {
			return true
		}
	}
	return false
}
