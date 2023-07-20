package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
	input := readFile()

	elf_calories := []int{}
	for _, elf_backpack := range getItemsPerElf(input) {
		elf_calories = append(elf_calories, sumCalories(elf_backpack))
	}

	sort.Slice(elf_calories, func(i, j int) bool { return elf_calories[i] < elf_calories[j] })

	fmt.Println("Part 1: ", elf_calories[len(elf_calories)-1])
}

func Part2() {
	input := readFile()

	elf_calories := []int{}
	for _, elf_backpack := range getItemsPerElf(input) {
		elf_calories = append(elf_calories, sumCalories(elf_backpack))
	}

	sort.Slice(elf_calories, func(i, j int) bool { return elf_calories[i] > elf_calories[j] })
	top_three := elf_calories[:3]

	fmt.Println("Part 2: ", top_three[0]+top_three[1]+top_three[2])
}

func readFile() string {
	input, error := os.ReadFile("day1/input.txt")
	if error != nil {
		panic(error)
	}

	return string(input)
}

func getItemsPerElf(input string) [][]string {
	items_per_elf := [][]string{}
	for _, elf_backpack := range strings.Split(input, "\n\n") {
		items_per_elf = append(items_per_elf, strings.Split(elf_backpack, "\n"))
	}

	return items_per_elf
}

func sumCalories(elf_backpack []string) int {
	total_calories := 0
	for _, item := range elf_backpack {
		if item == "" {
			continue
		}
		calories, error := strconv.Atoi(item)
		if error != nil {
			panic(error)
		}

		total_calories += calories
	}

	return total_calories
}

func Run() {
	fmt.Println("===== Day 1 =====")
	Part1()
	Part2()

	fmt.Println()
}
