package main

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
//
// Copies of scratchcards are scored like normal
// scratchcards and have the same card number as the
// card they copied. So, if you win a copy of card 10
// and it has 5 matching numbers, it would then win a
// copy of the same cards that the original card 10 won:
// cards 11, 12, 13, 14, and 15. This process repeats
// until none of the copies cause you to win any more
// cards. (Cards will never make you copy a card past the end of the table.)
//
// Once all of the originals and copies have been processed,
// you end up with 1 instance of card 1, 2 instances of card 2,
// 4 instances of card 3, 8 instances of card 4, 14 instances
// of card 5, and 1 instance of card 6. In total, this example
// pile of scratchcards causes you to ultimately have 30 scratchcards!
//
// Process all of the original and copied scratchcards until
// no more scratchcards are won. Including the original set
// of scratchcards, how many total scratchcards do you end up with?

func main() {
}
