package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Day 4: Passport Processing ---
// You arrive at the airport only to realize that you grabbed your North
// Pole Credentials instead of your passport. While these documents are
// extremely similar, North Pole Credentials aren't issued by a country
// and therefore aren't actually valid documentation for travel in most
// of the world.
//
// It seems like you're not the only one having problems, though; a very
// long line has formed for the automatic passport scanners, and the
// delay could upset your travel itinerary.
//
// Due to some questionable network security, you realize you might be
// able to solve both of these problems at the same time.
//
// The automatic passport scanners are slow because they're having trouble
// detecting which passports have all required fields. The expected fields
// are as follows:
//
// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)
// Passport data is validated in batch files (your puzzle input). Each passport
// is represented as a sequence of key:value pairs separated by spaces or
// newlines. Passports are separated by blank lines.
//
// Here is an example batch file containing four passports:
//
// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// byr:1937 iyr:2017 cid:147 hgt:183cm
//
// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
// hcl:#cfa07d byr:1929
//
// hcl:#ae17e1 iyr:2013
// eyr:2024
// ecl:brn pid:760753108 byr:1931
// hgt:179cm
//
// hcl:#cfa07d eyr:2025 pid:166559648
// iyr:2011 ecl:brn hgt:59in
// The first passport is valid - all eight fields are present. The second
// passport is invalid - it is missing hgt (the Height field).
//
// The third passport is interesting; the only missing field is cid, so it
// looks like data from North Pole Credentials, not a passport at all!
// Surely, nobody would mind if you made the system temporarily ignore
// missing cid fields. Treat this "passport" as valid.
//
// The fourth passport is missing two fields, cid and byr. Missing cid is
// fine, but missing any other field is not, so this passport is invalid.
//
// According to the above rules, your improved system would report 2 valid
// passports.
//
// Count the number of valid passports - those that have all required fields.
// Treat cid as optional. In your batch file, how many passports are valid?
func main() {
	fileScanner := createFileScanner()
	requireFields := initializeRequiredFiels()
	passportData := ""
	totalValids := 0
	for fileScanner.Scan() {
		currText := fileScanner.Text()
		if len(currText) == 0 {
			fmt.Printf("Passport Data: %v \n", passportData)
			passportMap := createMap(passportData) // Creates a hashmap out of the key of the passport fields.
			fmt.Printf("Passport Map: %v \n", passportMap)
			if isValid(requireFields, passportMap) {
				totalValids++
			}
			passportData = ""
		}
		passportData += currText + " "
	}
	fmt.Printf("Total Valids: %v \n", totalValids)
}

func createMap(passportData string) map[string]int {
	n := len(passportData)
	mp := make(map[string]int)
	for i := 0; i < n; i++ {
		if string(passportData[i]) == " " {
			continue
		}
		curr := ""
		for j := i; j < n; j++ {
			if string(passportData[j]) == ":" {
				mp[curr] = 0
				i = j
				break
			}
			if string(passportData[j]) == " " {
				i = j
				break
			}
			if j == n-1 {
				i = j
				break
			}
			curr += string(passportData[j])
		}
	}
	return mp
}

func isValid(requiredFields map[string]int, passportMap map[string]int) bool {
	for field := range requiredFields {
		_, ok := passportMap[field]
		if !ok {
			return false
		}
	}
	return true
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

func initializeRequiredFiels() map[string]int {
	requiredFieldsMap := map[string]int{
		"byr": 0,
		"iyr": 0,
		"eyr": 0,
		"hgt": 0,
		"hcl": 0,
		"ecl": 0,
		"pid": 0,
		// "cid": 0, // **Ignore cid field
	}
	return requiredFieldsMap
}
