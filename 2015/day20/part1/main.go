package main

import (
	"fmt"
	"sync"
)

// --- Day 20: Infinite Elves and Infinite Houses ---
// To keep the Elves busy, Santa has them deliver
// some presents by hand, door-to-door. He sends
// them down a street with infinite houses numbered
// sequentially: 1, 2, 3, 4, 5, and so on.
//
// Each Elf is assigned a number, too, and delivers
// presents to houses based on that number:
//
// The first Elf (number 1) delivers presents to
// every house: 1, 2, 3, 4, 5, ....
// The second Elf (number 2) delivers presents to
// every second house: 2, 4, 6, 8, 10, ....
// Elf number 3 delivers presents to every third
// house: 3, 6, 9, 12, 15, ....
// There are infinitely many Elves, numbered starting
// with 1. Each Elf delivers presents equal to ten
// times his or her number at each house.
//
// So, the first nine houses on the street end up like this:
//
// House 1 got 10 presents.
// House 2 got 30 presents.
// House 3 got 40 presents.
// House 4 got 70 presents.
// House 5 got 60 presents.
// House 6 got 120 presents.
// House 7 got 80 presents.
// House 8 got 150 presents.
// House 9 got 130 presents.
// The first house gets 10 presents: it is visited
// only by Elf 1, which delivers 1 * 10 = 10 presents.
// The fourth house gets 70 presents, because it is
// visited by Elves 1, 2, and 4, for a total of
// 10 + 20 + 40 = 70 presents.
//
// What is the lowest house number of the house to get
// at least as many presents as the number in your puzzle input?
func main() {
	input := 1612317
	mpHousePresents := findHouseToPresentsMap(input)
	// mpHousePresents = findHouseToPresentsMapAgain(input) // input as starting point
	// fmt.Printf("mp: %v \n", mpHousePresents)
	lowestHouseNum, presents := findLowestHouseNum((*mpHousePresents).houseMp, input)
	fmt.Printf("Lowest house number: %v \n", lowestHouseNum)
	fmt.Printf("Presents: %v \n", presents)
}

type Pair struct {
	p1 int
	p2 int
}

var wg sync.WaitGroup

type MP struct {
	houseMp     map[int]int
	multiplesMp map[Pair]int
	mu          sync.Mutex
}

func findHouseToPresentsMapAgain(start, num int, mp *MP) *MP {
	for i := start; i <= num; i++ {
		if i > 1612317 {
			fmt.Printf("Curr Num: %v \n", i)
		}
		wg.Add(1)
		go calculateAndMultiply(num, i, mp)
	}
	wg.Done()
	return mp
}

func findHouseToPresentsMap(num int) *MP {
	mp := &MP{
		houseMp:     make(map[int]int),
		multiplesMp: make(map[Pair]int),
	}
	for i := 1; i <= num; i++ {
		if i > 1600000 {
			fmt.Printf("Curr Num: %v \n", i)
		}
		wg.Add(1)
		go calculateAndMultiply(num, i, mp)
	}

	wg.Done()
	start := 1600000 + 1
	return findHouseToPresentsMapAgain(start, 29000000, mp)
}

func calculateAndMultiply(num, i int, mp *MP) {
	defer wg.Done()
	for j := i; j <= num; j++ {
		if j%i == 0 {
			p := Pair{p1: i, p2: 10}
			(*mp).mu.Lock()
			if _, ok := (*mp).multiplesMp[p]; ok {
				(*mp).houseMp[j] += (*mp).multiplesMp[p]
				continue
			}
			(*mp).multiplesMp[p] = i * 10
			(*mp).mu.Unlock()
		}
	}
}

func findLowestHouseNum(mp map[int]int, input int) (int, int) {
	houseNum, numOfPresents := 1, 0
	for currHouseNum, presents := range mp {
		if presents == input {
			if currHouseNum < houseNum {
				houseNum = currHouseNum
				numOfPresents = presents
			}
		}
	}
	return houseNum, numOfPresents
}
