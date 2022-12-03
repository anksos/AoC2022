/*
a-z: 1-26
A-Z: 27-52
Part 1: find the sum of the items that exist on both compartments compA and compB
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

// was used only for testing
//const items string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// build array with the items
/*
func buildItems(file string) int {
	if file >= "a" && file <= "z" {
		return int(file[0]) - 96
	} else if file >= "A" && file <= "Z" {
		return int(file[0]) - 38
	}
	return 0
}
*/
// don't like the below solution, but it works
func buildItems(file string) int {
	if file >= "a" && file <= "z" {
		return int(file[0]) - 96
	}
	return int(file[0]) - 38
}

func duplicates(file1 string, file2 string) string {
	var duplicates string
	for i := range file1 {
		for j := range file2 {
			if file1[i] == file2[j] {
				return string(file1[i])
			}
		}
	}
	return duplicates
}

func main() {
	// read input from file
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	// compartments sum
	compAB := strings.Split(string(file), "\n")
	sumItems := 0
	// search in the compartments for common items
	for i := range compAB {
		compA := compAB[i][0 : len(compAB[i])/2]
		compB := compAB[i][len(compAB[i])/2 : len(compAB[i])]
		sumItems += buildItems(duplicates(compA, compB))
		// testing: fmt.Println("compA: ", compA, "sumItems: ", sumItems)
		// testing: fmt.Println("compB: ", compB, "sumItems: ", sumItems)
	}
	fmt.Println(sumItems)
}
