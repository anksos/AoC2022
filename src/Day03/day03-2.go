package main

import (
	"fmt"
	"os"
	"strings"
)

func buildItems(file string) int {
	if file >= "a" && file <= "z" {
		return int(file[0]) - 96
	}
	return int(file[0]) - 38
}

func elfGroups(file1 string, file2 string, file3 string) string {
	var elfGroups string
	for i := range file1 {
		for j := range file2 {
			for k := range file3 {
				if file1[i] == file2[j] && file1[i] == file3[k] {
					return string(file1[i])
				}
			}
		}
	}
	return elfGroups
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	badge := strings.Split(string(file), "\n")
	sumItems := 0
	for i := 0; i < len(badge); i += 3 {
		sumItems += buildItems(elfGroups(badge[i], badge[i+1], badge[i+2]))
		//testing: fmt.Println("badge: ", badge[i], "sumItems: ", sumItems)
	}
	fmt.Println("sumItems: ", sumItems)
}
