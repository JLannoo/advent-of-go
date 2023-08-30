package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r1 Range) Contains(r2 Range) bool {
	return r2.Start >= r1.Start && r2.End <= r1.End
}

func (r1 Range) IsContained(r2 Range) bool {
	return r2.Start <= r1.Start && r2.End >= r1.End
}

func (r1 Range) Overlaps(r2 Range) bool {
	return r1.Start <= r2.End && r1.End >= r2.Start
}

func Part1() {
	input := readFile()

	pairs := strings.Split(input, "\n")

	containedPairs := make([]string, 0)
	for _, pair := range pairs {
		ranges := strings.Split(pair, ",")

		firstRange := toRange(ranges[0])
		secondRange := toRange(ranges[1])

		if firstRange.Contains(secondRange) || 
			firstRange.IsContained(secondRange) {
			containedPairs = append(containedPairs, pair)
		}
	}

	fmt.Println("Part 1:", len(containedPairs))
}

func Part2() {
	input := readFile()

	pairs := strings.Split(input, "\n")

	overlappingPairs := make([]string, 0)
	for _, pair := range pairs {
		ranges := strings.Split(pair, ",")

		firstRange := toRange(ranges[0])
		secondRange := toRange(ranges[1])

		if firstRange.Overlaps(secondRange) {
			overlappingPairs = append(overlappingPairs, pair)
		}
	}

	fmt.Println("Part 2:", len(overlappingPairs))
}

func readFile() string {
	input, err := os.ReadFile("day4/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

func Run() {
	fmt.Println("===== Day 4 =====")
	Part1()
	Part2()
}

func toRange(r string) Range {
	numbers := strings.Split(r, "-")
	ints := make([]int, len(numbers))

	for i, n := range numbers {
		number, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		ints[i] = number
	}

	return Range{Start: ints[0], End: ints[1]}
}