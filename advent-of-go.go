package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	
	"2022"
	"2023"
	"github.com/jlannoo/advent-of-go/year"
)

func main() {
	years := []year.Year{
		year_2022.Year,
		year_2023.Year,
	}

	yearPtr := flag.String("year", "", "Year to run (all, 2022)")
	dayPtr := flag.String("day", "", "Day to run (all, 1, 1-5)")

	flag.Parse()
	fmt.Print("\033[H\033[2J")

	fmt.Println("===== Advent of Go =====")


	if *yearPtr == "" {
		fmt.Println("Which year do you want to run? (all, 2022)")
		fmt.Println("Available years: ")
		for _, y := range years {
			fmt.Printf("%d ", y.Number)
		}
		fmt.Println()
		fmt.Scanln(yearPtr)
	}

	if *yearPtr == "all" {
		for _, y := range years {
			y.RunAllDays()
		}
		return
	}

	parsedInput, err := strconv.Atoi(*yearPtr)
	if err != nil {
		fmt.Println("Invalid year")
		return
	}

	selectedYear := year.Year{Number: 0}
	for _, y := range years {
		if y.Number == parsedInput {
			selectedYear = y
		}
	}

	if selectedYear.Number == 0 {
		fmt.Println("Year not found")
		return
	}

	fmt.Print("\033[H\033[2J")

	if *dayPtr == "" {
		fmt.Println("Which day do you want to run? (all, 1, 1-5)")
		fmt.Println("Available days: ")
		for _, d := range selectedYear.Days {
			fmt.Printf("%d ", d.Number)
		}
		fmt.Println()
		fmt.Scanln(dayPtr)
	}

	fmt.Print("\033[H\033[2J")

	if *dayPtr == "all" {
		selectedYear.RunAllDays()
		return
	}

	if strings.Contains(*dayPtr, "-") {
		var start, end int
		fmt.Sscanf(*dayPtr, "%d-%d", &start, &end)
		selectedYear.RunRange(start, end)
		return
	}

	var day int
	fmt.Sscanf(*dayPtr, "%d", &day)
	selectedYear.RunDay(day)
}
