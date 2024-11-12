package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// --- Part Two ---
// The line is moving more quickly now, but you overhear airport
// security talking about how passports with invalid data are
// getting through. Better add some data validation, quick!
//
// You can continue to ignore the cid field, but each other field
// has strict rules about what values are valid for automatic validation:
//
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
// Your job is to count the passports where all required fields are
// both present and valid according to the above rules.
//
// Count the number of valid passports - those that have all required
// fields and valid values. Continue to treat cid as optional. In
// your batch file, how many passports are valid?
func main() {
	fileScanner := createFileScanner()
	// requireFields := initializeRequiredFiels()
	passportData := ""
	totalValids := 0
	for fileScanner.Scan() {
		currText := fileScanner.Text()
		if len(currText) == 0 {
			fmt.Printf("Passport Data: %v \n", passportData)
			passportMap := createMap(passportData) // Creates a hashmap out of the key of the passport fields.
			fmt.Printf("Passport Map: %v \n", passportMap)
			// if isValid(requireFields, passportMap) {
			// 	totalValids++
			// }
			passportData = ""
		}
		passportData += currText + " "
	}
	fmt.Printf("Total Valids: %v \n", totalValids)
}

func createMap(passportData string) map[string]string {
	n := len(passportData)
	mp := make(map[string]string)
	for i := 0; i < n; i++ {
		if string(passportData[i]) == " " {
			continue
		}
		field, isFieldParsed := "", false
		value, readyToParseValue := "", false
		for j := i; j < n; j++ {
			if string(passportData[j]) == ":" {
				mp[field] = ""
				isFieldParsed = true
				readyToParseValue = true
				i = j
				break
			}
			if string(passportData[j]) == " " {
				mp[field] = value
				i = j
				break
			}
			if !isFieldParsed {
				field += string(passportData[j])
			}
			if readyToParseValue {
				value += string(passportData[j])
			}
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
