package day5

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jlannoo/advent-of-go/year"
)

type Dock struct {
	Stacks       []Stack
	Instructions []Instruction
}

func (dock *Dock) ExecuteInstructionPart1(instruction Instruction) {
	stacks := dock.Stacks

	for i := 0; i < instruction.Quantity; i++ {
		fromStack := &stacks[instruction.From-1]
		toStack := &stacks[instruction.To-1]

		movedCrate := fromStack.Crates[len(fromStack.Crates)-1]
		fromStack.Crates = fromStack.Crates[:len(fromStack.Crates)-1]
		toStack.Crates = append(toStack.Crates, movedCrate)
	}

	dock.Stacks = stacks
}

func (dock *Dock) ExecuteInstructionPart2(instruction Instruction) {
	stacks := dock.Stacks

	fromStack := &stacks[instruction.From-1]
	toStack := &stacks[instruction.To-1]

	movedCrates := fromStack.Crates[len(fromStack.Crates)-instruction.Quantity:]
	fromStack.Crates = fromStack.Crates[:len(fromStack.Crates)-instruction.Quantity]

	toStack.Crates = append(toStack.Crates, movedCrates...)

	dock.Stacks = stacks
}

func (dock *Dock) ExecuteInstructionsPart1() {
	for _, instruction := range dock.Instructions {
		dock.ExecuteInstructionPart1(instruction)
	}
}

func (dock *Dock) ExecuteInstructionsPart2() {
	for _, instruction := range dock.Instructions {
		dock.ExecuteInstructionPart2(instruction)
	}
}

func (dock *Dock) GetTopCrates() []string {
	topCrates := []string{}

	for _, stack := range dock.Stacks {
		topCrates = append(topCrates, stack.Crates[len(stack.Crates)-1])
	}

	return topCrates
}

func (dock *Dock) Print() {
	str, _ := json.MarshalIndent(dock.Stacks, "", "  ")
	fmt.Print(string(str))
	fmt.Println("\n")
}

type Stack struct {
	Crates []string
}

type Instruction struct {
	Quantity int
	From     int
	To       int
}

func Part1() {
	input := readFile()

	dock := parseInput(input)

	dock.ExecuteInstructionsPart1()

	fmt.Println("Part 1:", strings.Join(dock.GetTopCrates(), ""))
}

func Part2() {
	input := readFile()

	dock := parseInput(input)

	dock.ExecuteInstructionsPart2()

	fmt.Println("Part 2:", strings.Join(dock.GetTopCrates(), ""))
}

func parseInput(input string) Dock {
	split := strings.Split(input, "\n\n")

	stacksText := split[0]
	instructionsText := split[1]

	stacks := parseStacks(stacksText)
	instructions := parseInstructions(instructionsText)

	return Dock{stacks, instructions}
}

func parseStacks(stacksText string) []Stack {
	lines := strings.Split(stacksText, "\n")
	reversedLines := reverse(lines)

	stacks := []Stack{}
	for i, char := range reversedLines[0] {
		if char == ' ' {
			continue
		}

		crates := []string{}

		for _, line := range reversedLines[1:] {
			char := string(line[i])
			if char == " " {
				break
			}
			crates = append(crates, string(line[i]))
		}

		stacks = append(stacks, Stack{crates})
	}

	return stacks
}

func parseInstructions(instructionsText string) []Instruction {
	instructions := strings.Split(instructionsText, "\n")

	var parsed []Instruction

	for _, instruction := range instructions {
		split := strings.Split(instruction, " ")

		quantity, err1 := strconv.Atoi(split[1])
		from, err2 := strconv.Atoi(split[3])
		to, err3 := strconv.Atoi(split[5])

		if err1 != nil || err2 != nil || err3 != nil {
			panic("Error parsing instruction")
		}

		parsed = append(parsed, Instruction{quantity, from, to})
	}

	return parsed
}

func reverse(lines []string) []string {
	var reversed []string

	for i := len(lines) - 1; i >= 0; i-- {
		reversed = append(reversed, lines[i])
	}

	return reversed
}

func readFile() string {
	input, err := os.ReadFile("2022/day5/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

func Run() {
	fmt.Println("===== Day 5 =====")
	Part1()
	Part2()

	fmt.Println()
}

var Day = year.Day{
	Number: 5,
	Run:    Run,
}
