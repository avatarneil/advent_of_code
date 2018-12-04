package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// It would be better for such a function to return error, instead of handling
// it on their own.
func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func sum(input []int) int {
	freq := 0
	for _, item := range input {
		freq += item
	}
	return freq
}

func rollingSum(start int, input []int, seen map[int]bool) int {
	freq := start
	for _, item := range input {
		freq += item
		if seen[freq] {
			return freq
		} else {
			seen[freq] = true
		}
	}
	return rollingSum(freq, input, seen)
}

func main() {
	nums, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}
	// Part a
	fmt.Println(sum(nums))

	// Part b (probably could've done this so much better)
	seen := make(map[int]bool)
	fmt.Println(rollingSum(0, nums, seen))
}
