package day3

import (
	"fmt"
	"os"
	"strings"

	"github.com/jlannoo/advent-of-go/year"
)

func Part1() {
	input := readFile()

	sum := 0
	for _, rucksack := range strings.Split(input, "\n") {
		firstCompartment := rucksack[0 : len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		repeatedChars := findRepeated(firstCompartment, secondCompartment)

		prioritySum := 0
		for _, char := range repeatedChars {
			prioritySum += getPriority(char)
		}

		sum += prioritySum
	}

	fmt.Println("Part 1: ", sum)
}

func Part2() {
	input := readFile()

	groups := getGroups(input)

	repeatedItems := []string{}
	for _, group := range groups {
		first := group[0]

		for _, item := range first {
			if strings.Contains(group[1], string(item)) && strings.Contains(group[2], string(item)) {
				repeatedItems = append(repeatedItems, string(item))
				break
			}
		}
	}

	sum := 0
	for _, item := range repeatedItems {
		sum += getPriority(item)
	}

	fmt.Println("Part 2: ", sum)
}

func readFile() string {
	input, err := os.ReadFile("2022/day3/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

func getPriority(char string) int {
	// a-z -> 1-26
	// A-Z -> 27-52

	if isLowerCase(char) {
		return int(char[0]) - 97 + 1
	} else {
		return int(char[0]) - 65 + 27
	}
}

func findRepeated(r1 string, r2 string) []string {
	repeated := []string{}

	for _, char := range r1 {
		char := string(char)
		if strings.Contains(r2, char) {
			repeated = append(repeated, char)
			break
		}
	}

	return repeated
}

func isLowerCase(str string) bool {
	return strings.ToLower(str) == str
}

func getGroups(input string) [][]string {
	groups := [][]string{}

	group := []string{}
	for _, rucksack := range strings.Split(input, "\n") {
		group = append(group, rucksack)

		if len(group) == 3 {
			groups = append(groups, group)
			group = []string{}
		}
	}

	return groups
}

func Run() {
	fmt.Println("===== Day 3 =====")
	Part1()
	Part2()

	fmt.Println()
}

var Day = year.Day{
	Number: 3,
	Run:    Run,
}
