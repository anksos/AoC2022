/*
Q: In how many assignment pairs does one range fully contain the other?
We know assignments have equal length for each line also assignments are separated by coma character
build strings into arrays of assignments and compare the overlaps
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	countContained := 0
	allOveralps := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		left, right := split[0], split[1]
		leftSlice := intSlice(strings.Split(left, "-"))
		rightSlice := intSlice(strings.Split(right, "-"))

		if leftSlice[0] >= rightSlice[0] && leftSlice[1] <= rightSlice[1] || rightSlice[0] >= leftSlice[0] && rightSlice[1] <= leftSlice[1] {
			countContained++
		}
		if !(leftSlice[0] > rightSlice[1] || leftSlice[1] < rightSlice[0]) {
			allOveralps++
		}
	}
	fmt.Println(countContained)
	fmt.Println(allOveralps)
}

func intSlice(slice []string) []int {
	var intSlice []int
	for _, s := range slice {
		i, _ := strconv.Atoi(s)
		intSlice = append(intSlice, i)
	}
	return intSlice
}
