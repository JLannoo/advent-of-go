package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jlannoo/advent-of-go/day1"
	"github.com/jlannoo/advent-of-go/day2"
	"github.com/jlannoo/advent-of-go/day3"
	"github.com/jlannoo/advent-of-go/day4"
	"github.com/jlannoo/advent-of-go/day5"
	"github.com/jlannoo/advent-of-go/day6"
	"github.com/jlannoo/advent-of-go/day7"
	"github.com/jlannoo/advent-of-go/day8"
)

func main() {
	days := []func(){
		day1.Run,
		day2.Run,
		day3.Run,
		day4.Run,
		day5.Run,
		day6.Run,
		day7.Run,
		day8.Run,
	}

	fmt.Println("===== Advent of Go =====")

	args := ""

	if len(os.Args) == 1 {
		fmt.Println("Please provide a day or range (n-m) or 'all'")
		fmt.Scanln(&args)
	} else {
		args = os.Args[1]
	}

	isAll := args == "all" || args == "All" || args == "ALL"
	if isAll {
		fmt.Println("Running all days")
		runAll(days)
		os.Exit(0)
	}

	split := strings.Split(args, "-")
	isRange := len(split) == 2
	if isRange {
		fmt.Println("Running range", split[0], "-", split[1])
		start := toInt(split[0])
		end := toInt(split[1])

		runRange(days, start, end)
		os.Exit(0)
	}

	day, err := strconv.Atoi(args)
	isDay := err == nil && day > 0 && day < 26
	if isDay {
		fmt.Println("Running day", day)
		days[day-1]()
		os.Exit(0)
	}

	fmt.Println("Invalid input")
}

func toInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}

	return i
}

func runAll(days []func()) {
	for _, day := range days {
		day()
	}
}

func runRange(days []func(), start int, end int) {
	for i := start; i <= end; i++ {
		days[i-1]()
	}
}
