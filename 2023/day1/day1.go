package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jlannoo/advent-of-go/year"
)

type Line struct {
	First  rune
	Last   rune
	Number int
}

var NUMBER_WORD_MAP = map[int]string{
	49: "one",
	50: "two",
	51: "three",
	52: "four",
	53: "five",
	54: "six",
	55: "seven",
	56: "eight",
	57: "nine",
}

func ParseLine(line string) Line {
	first := ' '
	last := ' '

	for _, char := range line {
		_, err := strconv.Atoi(string(char))
		if err == nil {
			if first == ' ' {
				first = char
			}
			last = char
		}
	}

	return NewLine(first, last)
}

func ParseLine2(line string) Line {
	first := ' '
	last := ' '

	for i := 0; i < len(line); i++ {
		char := line[i]
		_, err := strconv.Atoi(string(char))

		// Is a number
		if err == nil {
			if first == ' ' {
				first = rune(char)
			}
			last = rune(char)
			continue
		}

		// Is a word
		for number, word := range NUMBER_WORD_MAP {
			if strings.HasPrefix(line[i:], word) {
				if first == ' ' {
					first = rune(number)
				}
				last = rune(number)
				break
			}
		}
	}

	return NewLine(first, last)
}

func NewLine(first rune, last rune) Line {
	number, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		panic(err)
	}

	return Line{
		First:  first,
		Last:   last,
		Number: number,
	}
}

func SumLines(lines []Line) int {
	sum := 0

	for _, line := range lines {
		sum += line.Number
	}

	return sum
}

func Part1() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	var linesParsed []Line
	for _, line := range lines {
		linesParsed = append(linesParsed, ParseLine(line))
	}

	fmt.Println("Part 1:", SumLines(linesParsed))
}

func Part2() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	var linesParsed []Line
	for _, line := range lines {
		linesParsed = append(linesParsed, ParseLine2(line))
	}

	fmt.Println("Part 2:", SumLines(linesParsed))
}

func Run() {
	fmt.Println("===== Day 1 =====")
	Part1()
	Part2()

	fmt.Println()
}

func ReadFile() string {
	input, err := os.ReadFile("2023/day1/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

var Day = year.Day{
	Number: 1,
	Run:    Run,
}
