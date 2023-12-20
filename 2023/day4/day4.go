package day4

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/jlannoo/advent-of-go/year"
)

type Card struct {
	ID             int
	WinningNumbers mapset.Set
	MyNumbers      mapset.Set
	Matches        mapset.Set
}

type CardInventory = map[int]int

func (card *Card) GetMatches() {
	matches := mapset.NewSet()

	for _, number := range card.MyNumbers.ToSlice() {
		if card.WinningNumbers.Contains(number) {
			matches.Add(number)
		}
	}

	card.Matches = matches
}

func (card *Card) GetPoints() int {
	return int(math.Pow(2, float64(card.Matches.Cardinality()-1)))
}

func ParseCard(card string) Card {
	lineSplit := strings.Split(card, ": ")

	idSplit := strings.Split(lineSplit[0], " ")
	id, _ := strconv.Atoi(idSplit[len(idSplit)-1])

	numbersSplit := strings.Split(lineSplit[1], " | ")
	winningNumbers := mapset.NewSet()
	myNumbers := mapset.NewSet()

	for _, number := range strings.Split(numbersSplit[0], " ") {
		n, _ := strconv.Atoi(number)
		if n == 0 {
			continue
		}

		winningNumbers.Add(n)
	}

	for _, number := range strings.Split(numbersSplit[1], " ") {
		n, _ := strconv.Atoi(number)
		if n == 0 {
			continue
		}

		myNumbers.Add(n)
	}

	return Card{
		ID:             id,
		WinningNumbers: winningNumbers,
		MyNumbers:      myNumbers,
	}
}

func Part1() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	var cards []Card
	for _, line := range lines {
		card := ParseCard(line)
		card.GetMatches()

		cards = append(cards, card)
	}

	sum := 0
	for _, card := range cards {
		sum += card.GetPoints()
	}

	fmt.Println("Part 1:", sum)
}

func Part2() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	var cards []Card
	inventory := CardInventory{}

	for _, line := range lines {
		card := ParseCard(line)
		card.GetMatches()

		cards = append(cards, card)
	}

	for _, card := range cards {
		inventory[card.ID] = 1
	}

	sum := 0
	for i := 0; i < len(inventory); i++ {
		card := cards[i]
		amount := inventory[card.ID]
		
		for i := 0; i < card.Matches.Cardinality(); i++ {
			inventory[card.ID+1+i] += amount
		}

		sum += amount
	}

	fmt.Println("Part 2:", sum)
}

func ReadFile() string {
	input, err := os.ReadFile("2023/day4/input.txt")

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

var Day = year.Day{
	Number: 4,
	Run:    Run,
}
