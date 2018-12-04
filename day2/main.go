package main

import (
	"io/ioutil"
	"strings"
)

func computeChecksum(data []string) int {
	twos := 0
	threes := 0
	for _, item := range data {
		matches := make(map[rune]int)
		for _, c := range item {
			matches[c]++
		}
		if containsValue(matches, 2) {
			twos++
		}
		if containsValue(matches, 3) {
			threes++
		}
	}
	return twos * threes
}

func containsValue(m map[rune]int, v int) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}
	return false
}

func findMatches(data []string) string {
	for _, itema := range data {
		for _, itemb := range data {
			if itema != itemb {
				if difference(itema, itemb) == 1 {
					return returnMatches(itema, itemb)
				}
			}
		}
	}
	return "Sad Parrot"
}

func difference(a string, b string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func returnMatches(a string, b string) string {
	output := ""
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			output += string(a[i])
		}
	}
	return output
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		//Do something
	}

	lines := strings.Split(string(content), "\n")

	// Part a
	println(computeChecksum(lines))

	// Part b
	println(findMatches(lines))
}
