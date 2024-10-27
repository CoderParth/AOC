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
	guideMap := initialize()
	fileScanner := createFileScanner()
	total := 0
	for fileScanner.Scan() {
		opponentChoice, wayRoundNeedstoEnd := parseLine(fileScanner.Text(), guideMap)
		total += matchOutcome(opponentChoice, wayRoundNeedstoEnd)
	}
	fmt.Printf("Total Score: %v \n", total)
}

func matchOutcome(opponentChoice, wayRoundNeedstoEnd string) int {
	// opponent chooses rock
	if opponentChoice == "Rock" {
		if wayRoundNeedstoEnd == "Draw" {
			// Choose rock , get 1 point
			// Get 3 points for draw
			return 4 // 1 + 3
		}
		if wayRoundNeedstoEnd == "Lose" {
			// Choose Scissors, get 3 points
			// No points for Losing - 0
			return 3 // 0 + 3
		}
		if wayRoundNeedstoEnd == "Win" {
			// Choose Paper, Get 2 points
			// Get 6 points fo winning
			return 8
		}
	}

	// opponent chooses Scissors
	if opponentChoice == "Scissors" {
		if wayRoundNeedstoEnd == "Draw" {
			// Choose Scissors, get 3 points
			// Get 3 points for draw
			return 6 // 3 + 3
		}
		if wayRoundNeedstoEnd == "Lose" {
			// Choose Paper, get 2 points
			// Get 0 points for loss
			return 2 // 0 + 2
		}
		if wayRoundNeedstoEnd == "Win" {
			// Choose Rock, get 1 point
			// Get 6 points for a win
			return 7 // 6 + 1
		}
	}

	// opponent chooses Paper
	if opponentChoice == "Paper" {
		if wayRoundNeedstoEnd == "Draw" {
			// Choose Paper, Get 2 points
			// Get 3 points for draw
			return 5 // 2 + 3
		}
		if wayRoundNeedstoEnd == "Lose" {
			// Choose Rock, get 1 point
			// Get 0 points for loss
			return 1 // 0 + 1
		}
		if wayRoundNeedstoEnd == "Win" {
			// Choose Scissors, get 3 points
			// Get 6 points for a win
			return 9 // 6 + 3
		}
	}

	return 0
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

func createFileScanner() *bufio.Scanner {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func initialize() map[string]string {
	guideMap := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}
	return guideMap
}
