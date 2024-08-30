package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// For added security, yet another system policy has been put in place.
// Now, a valid passphrase must contain no two words that are anagrams
// of each other - that is, a passphrase is invalid if any word's
// letters can be rearranged to form any other word in the passphrase.
//
// For example:
//
// abcde fghij is a valid passphrase.
// abcde xyz ecdab is not valid - the letters from the third word
// can be rearranged to form the first word.
// a ab abc abd abf abj is a valid passphrase, because all letters
// need to be used when forming another word.
// iiii oiii ooii oooi oooo is valid.
// oiii ioii iioi iiio is not valid - any of these words can be
// rearranged to form any other word.
// Under this new system policy, how many passphrases are valid?
func main() {
	totalValid := 0
	fileScanner := createFileScanner()
	for fileScanner.Scan() {
		wordsMap := createWordsMap(fileScanner.Text())
		fmt.Printf("words map: %v \n", wordsMap)
		if isCurrLineValid(wordsMap) {
			totalValid++
		}
	}
	fmt.Printf("Total Valid: %v \n", totalValid)
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

func createWordsMap(line string) map[string]int {
	wordsMap := make(map[string]int)
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
				wordsMap[curr]++
				i = j
				break
			}
			if j == n-1 {
				curr += string(line[j])
				wordsMap[curr]++
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return wordsMap
}

func isCurrLineValid(wordsMap map[string]int) bool {
	for _, wordCount := range wordsMap {
		if wordCount > 1 {
			return false
		}
	}
	return true
}
