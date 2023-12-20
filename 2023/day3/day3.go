package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jlannoo/advent-of-go/year"
)

type Coordinate struct {
	X int
	Y int
}

type PartNumber struct {
	Value int
	Start Coordinate
	End   Coordinate
}

type Gear struct {
	Coord       Coordinate
	PartNumbers []PartNumber
	Ratio       int
}

func (number *PartNumber) IsAdjacent(coord Coordinate) bool {
	return ((number.Start.X-1 <= coord.X && coord.X <= number.End.X+1) &&
		(number.Start.Y-1 <= coord.Y && coord.Y <= number.End.Y+1))
}

func (number *PartNumber) HasSymbolAdjacent(lines []string) bool {
	for y := number.Start.Y - 1; y <= number.End.Y+1; y++ {
		for x := number.Start.X - 1; x <= number.End.X+1; x++ {
			if x >= 0 && y >= 0 && x < len(lines[0]) && y < len(lines) {
				char := lines[y][x]
				isDigit := isDigit(string(char))
				if !isDigit && char != '.' {
					return true
				}
			}
		}
	}

	return false
}

func ParseGears(lines []string, partNumbers []PartNumber) []Gear {
	var gears []Gear

	for y, line := range lines {
		for x, char := range line {
			if char == '*' {
				coord := Coordinate{X: x, Y: y}
				gearParts := []PartNumber{}

				for _, number := range partNumbers {
					if number.IsAdjacent(coord) {
						gearParts = append(gearParts, number)
					}
				}

				if len(gearParts) == 2 {
					gears = append(gears, Gear{
						Coord: coord, 
						PartNumbers: gearParts,
						Ratio: gearParts[0].Value * gearParts[1].Value,
					})
				}
			}
		}
	}

	return gears
}

func ParsePartNumbers(lines []string) []PartNumber {
	var numbers []PartNumber

	for y, line := range lines {
		collect := ""
		var start Coordinate

		for x, char := range line {
			if isDigit(string(char)) {
				if collect == "" {
					start = Coordinate{X: x, Y: y}
				}
				collect += string(char)

				if x == len(line)-1 {
					value, _ := strconv.Atoi(collect)

					end := Coordinate{X: x, Y: y}
					numbers = append(numbers, PartNumber{Value: value, Start: start, End: end})
				}
			} else {
				if collect != "" {
					value, _ := strconv.Atoi(collect)

					end := Coordinate{X: x - 1, Y: y}
					numbers = append(numbers, PartNumber{Value: value, Start: start, End: end})
					collect = ""
				}
			}
		}
	}

	return numbers
}

func Part1() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	numbers := ParsePartNumbers(lines)

	var partNumbers []int
	for _, number := range numbers {
		if number.HasSymbolAdjacent(lines) {
			partNumbers = append(partNumbers, number.Value)
		}
	}

	sum := 0
	for _, number := range partNumbers {
		sum += number
		// fmt.Println(number)
	}

	fmt.Println("Part 1:", sum)
}

func Part2() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	numbers := ParsePartNumbers(lines)
	gears := ParseGears(lines, numbers)

	sum := 0
	for _, gear := range gears {
		sum += gear.Ratio
	}

	fmt.Println("Part 2:", sum)
}

func ReadFile() string {
	input, err := os.ReadFile("2023/day3/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func Run() {
	fmt.Println("===== Day 3 =====")
	Part1()
	Part2()
}

var Day = year.Day{
	Number: 3,
	Run:    Run,
}
