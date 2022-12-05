package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	countContained := 0 // part one variable
	allOveralps := 0    // part two variable
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		left, right := split[0], split[1]
		leftSlice := intSlice(strings.Split(left, "-"))
		rightSlice := intSlice(strings.Split(right, "-"))

		// part one, checking if the left slice is contained in the right slice or vice versa
		if leftSlice[0] >= rightSlice[0] && leftSlice[1] <= rightSlice[1] || rightSlice[0] >= leftSlice[0] && rightSlice[1] <= leftSlice[1] {
			countContained++
		}
		// part two, checking if the left slice overlaps with the right slice or vice versa and how many times
		if !(leftSlice[0] > rightSlice[1] || leftSlice[1] < rightSlice[0]) {
			allOveralps++
		}
	}
	fmt.Println(countContained)
	fmt.Println(allOveralps)
}

// creating an array of ints from an array of strings (for the input) using slice as is more flexible than arrays
// and you don't need to specify the size.
func intSlice(slice []string) []int {
	var intSlice []int
	for _, s := range slice {
		i, _ := strconv.Atoi(s)
		intSlice = append(intSlice, i)
	}
	return intSlice
}
