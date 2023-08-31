package day7

import (
	"fmt"
	"os"
	"strings"
)

const (
	CD = "cd"
	LS = "ls"
)

type Command struct {
	Operation string
	Argument  string
	Output    []string
}

/*
	Part 1

	Find the sum of the sizes of all directories that are smaller than 100000
*/
func Part1() {
	input := readFile()

	root := TreeNode{IsFolder: true, Name: "root"}

	tree := Tree{
		Root:          root,
		CurrentFolder: &root,
	}

	lines := strings.Split(input, "\n")
	instructions := getInstructions(lines)

	for _, instruction := range instructions {
		switch instruction.Operation {
		case CD: // Change directory
			tree.ChangeDirectory(instruction.Argument)

		case LS: // Create directory contents
			tree.CreateDirectoryContents(instruction)
		}
	}

	directories := tree.GetDirectoriesSmallerThan(100000)

	sum := 0
	for _, dir := range directories {
		sum += dir.Size
	}

	fmt.Println("Part 1:", sum)
}

/*
	Part 2

	Choose the smallest directory to delete that would leave enough space
*/
func Part2() {
	input := readFile()

	const TOTAL_FS_SIZE = 70000000
	const REQUIRED_SIZE = 30000000

	root := TreeNode{IsFolder: true, Name: "root"}

	tree := Tree{
		Root:          root,
		CurrentFolder: &root,
	}

	lines := strings.Split(input, "\n")
	instructions := getInstructions(lines)

	for _, instruction := range instructions {
		switch instruction.Operation {
		case CD: // Change directory
			tree.ChangeDirectory(instruction.Argument)

		case LS: // Create directory contents
			tree.CreateDirectoryContents(instruction)
		}
	}

	usedSpace := tree.Root.GetTotalSize()
	freeSpace := TOTAL_FS_SIZE - usedSpace
	spaceToFree := REQUIRED_SIZE - freeSpace

	directories := tree.GetDirectoriesWithSize()
	smallestBigEnough := directories[0]
	for _, dir := range directories {
		if dir.Size >= spaceToFree && dir.Size < smallestBigEnough.Size {
			smallestBigEnough = dir
		}
	}

	fmt.Printf("Part 2: %s (%d)\n", smallestBigEnough.GetFullPath(), smallestBigEnough.Size)
}

func getInstructions(input []string) []Command {
	var instructions []Command

	output := []string{}

	// Traverse backwards to get the output for each command
	for i := len(input) - 1; i >= 0; i-- {
		line := input[i]
		if line[0] == '$' {
			split := strings.Split(line, " ")

			operation := split[1]
			var argument string

			if len(split) >= 3 {
				argument = split[2]
			}

			instructions = append(instructions, Command{
				Operation: operation,
				Argument:  argument,
				Output:    output,
			})

			output = []string{}
		} else {
			output = append(output, line)
		}
	}

	// Reverse the instructions so they're in the correct order
	for i, j := 0, len(instructions)-1; i < j; i, j = i+1, j-1 {
		instructions[i], instructions[j] = instructions[j], instructions[i]
	}

	return instructions
}

func readFile() string {
	input, err := os.ReadFile("day7/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

func Run() {
	fmt.Println("===== Day 7 =====")
	Part1()
	Part2()

	fmt.Println()
}
