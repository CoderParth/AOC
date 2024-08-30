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
		wordsArr := createWordsArr(fileScanner.Text())
		if isValidPassphrase(wordsArr) {
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

func createWordsArr(line string) []string {
	wordsArr := []string{}
	n := len(line)
	for i := 0; i < n; i++ {
		if string(line[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
				wordsArr = append(wordsArr, curr)
				i = j
				break
			}
			if j == n-1 {
				curr += string(line[j])
				wordsArr = append(wordsArr, curr)
				i = j
				break
			}
			curr += string(line[j])
		}
	}
	return wordsArr
}

func isValidPassphrase(wordsArr []string) bool {
	n := len(wordsArr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if areAnagrams(wordsArr[i], wordsArr[j]) {
				return false
			}
		}
	}
	return true
}

func areAnagrams(word1, word2 string) bool {
	m, n := len(word1), len(word2)
	if m != n {
		return false
	}
	map1, map2 := make(map[rune]int), make(map[rune]int)
	for i := 0; i < m; i++ {
		map1[rune(word1[i])]++
		map2[rune(word2[i])]++
	}
	for k1, v1 := range map1 {
		if v1 != map2[k1] {
			return false
		}
	}
	return true
}
