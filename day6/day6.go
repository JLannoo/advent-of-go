package day6

import (
	"fmt"
	"os"
	"strings"
	mapset "github.com/deckarep/golang-set"
)

const SIGNAL_LENGTH_PART_1 = 4
const SIGNAL_LENGTH_PART_2 = 14

func Part1() {
	input := readFile()

	index := findIndex(input, SIGNAL_LENGTH_PART_1)

	fmt.Println("Part 1: ", index)
}

func Part2() {
	input := readFile()

	index := findIndex(input, SIGNAL_LENGTH_PART_2)

	fmt.Println("Part 2: ", index)
}

func findIndex(input string, signalLength int) int {
	var index int
	for i := 0; i < len(input)-signalLength; i++ {
		signal := strings.Split(input[i:i+signalLength], "")

		set := mapset.NewSet()
		for _, s := range signal {
			set.Add(s)
		}

		if set.Cardinality() == signalLength {
			index = i + signalLength
			break
		}		
	}

	return index
}

func readFile() string {
	input, err := os.ReadFile("day6/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

func Run() {
	fmt.Println("===== Day 6 =====")
	Part1()
	Part2()

	fmt.Println()
}
