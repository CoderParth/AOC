package main

import "fmt"

// 199
// 200
// 208
// 210
// 200
// 207
// 240
// 269
// 260
// 263

// 199 (N/A - no previous measurement)
// 200 (increased)
// 208 (increased)
// 210 (increased)
// 200 (decreased)
// 207 (increased)
// 240 (increased)
// 269 (increased)
// 260 (decreased)
// 263 (increased)
// In this example, there are 7 measurements that are larger than the previous measurement.
//
// How many measurements are larger than the previous measurement?

func main() {
	arrOfMeasurements := createArrOfMeasurements()
	numOfMeasurementsLargerThanPrevious := findTotalLargerMeasurements(arrOfMeasurements)
	fmt.Printf("Number of measurements: ", numOfMeasurementsLargerThanPrevious)
}

func findTotalLargerMeasurements(arr []int) int {
	total, prev := 0, arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > prev {
			total++
		}
		prev = arr[i]
	}
	return total
}
