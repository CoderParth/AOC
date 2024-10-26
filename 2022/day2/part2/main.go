package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// The Elf finishes helping with the tent and sneaks back over to you.
// "Anyway, the second column says how the round needs to end: X means
// you need to lose, Y means you need to end the round in a draw, and
// Z means you need to win. Good luck!"
//
// The total score is still calculated in the same way, but now you
// need to figure out what shape to choose so the round ends as
// indicated. The example above now goes like this:
//
// In the first round, your opponent will choose Rock (A), and you
// need the round to end in a draw (Y), so you also choose Rock.
// This gives you a score of 1 + 3 = 4.
// In the second round, your opponent will choose Paper (B), and
// you choose Rock so you lose (X) with a score of 1 + 0 = 1.
// In the third round, you will defeat your opponent's Scissors
// with Rock for a score of 1 + 6 = 7.
// Now that you're correctly decrypting the ultra top secret
// strategy guide, you would get a total score of 12.
//
// Following the Elf's instructions for the second column,
// what would your total score be if everything goes exactly
// according to your strategy guide?
func main() {
	guideMap, pointsMap := initialize()
	fileScanner := createFileScanner()
	total := 0
	for fileScanner.Scan() {
		opponentChoice, yourChoice := parseLine(fileScanner.Text(), guideMap)
		total += pointsMap[yourChoice] // points for the shape you selected
		total += matchOutcome(opponentChoice, yourChoice)
	}
	fmt.Printf("Total Score: %v \n", total)
}

func parseLine(line string, guideMap map[string]string) (string, string) {
	n := len(line)
	arr := []string{}
	for i := 0; i < n; i++ {
		curr := ""
		for j := i; j < n; j++ {
			if string(line[j]) == " " {
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
	opponentChoice := guideMap[arr[0]]
	yourChoice := guideMap[arr[1]]
	return opponentChoice, yourChoice
}

func matchOutcome(opponentChoice, yourChoice string) int {
	// opponent chooses rock
	if opponentChoice == "Rock" {
		if yourChoice == "Paper" {
			return 6 // won
		}
		if yourChoice == "Scissor" {
			return 0 // lost
		}
	}
	// opponent chooses paper
	if opponentChoice == "Paper" {
		if yourChoice == "Rock" {
			return 0 // lost
		}
		if yourChoice == "Scissor" {
			return 6 // won
		}
	}
	// opponent chooses scissors
	if opponentChoice == "Scissor" {
		if yourChoice == "Rock" {
			return 6 // won
		}
		if yourChoice == "Paper" {
			return 0 // lost
		}
	}
	return 3 // draw - you chose same as your opponent
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

func initialize() (map[string]string, map[string]int) {
	guideMap := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissor",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissor",
	}

	pointsMap := map[string]int{
		"Rock":    1,
		"Paper":   2,
		"Scissor": 3,
	}
	return guideMap, pointsMap
}
