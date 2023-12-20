package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jlannoo/advent-of-go/year"
)

type Reveal struct {
	Quantity int
	Color    string
}

type Game struct {
	ID  int
	Sets [][]Reveal
}

type Configuration = map[string]int

func (game *Game) CheckIfPossible(initial Configuration) bool {
	for _, set := range game.Sets {
		for _, reveal := range set {
			if initial[reveal.Color] < reveal.Quantity {
				return false
			}
		}
	}

	return true
}

func (game *Game) SmallestPossible() Configuration {
	config := make(Configuration)

	for _, set := range game.Sets {
		for _, reveal := range set {
			if config[reveal.Color] < reveal.Quantity {
				config[reveal.Color] = reveal.Quantity
			}
		}
	}

	return config
}

func NewGame(game string) *Game {
	gameStringSplit := strings.Split(game, ": ")

	nameSplit := strings.Split(gameStringSplit[0], " ")
	gameId, err := strconv.Atoi(nameSplit[1])
	if err != nil {
		panic("Invalid game id")
	}

	setsStrings := strings.Split(gameStringSplit[1], "; ")

	var sets [][]Reveal
	for _, set := range setsStrings {
		pulls := strings.Split(set, ", ")

		var set []Reveal
		for _, pull := range pulls {
			splitPull := strings.Split(pull, " ")

			quantity, _ := strconv.Atoi(splitPull[0])
			color := splitPull[1]

			reveal := Reveal{
				Quantity: quantity,
				Color:    color,
			}

			set = append(set, reveal)
		}

		sets = append(sets, set)
	}

	return &Game{
		ID:  gameId,
		Sets: sets,
	}
}

func Part1() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	config := Configuration{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var games []*Game
	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	var possible []*Game
	for _, game := range games {
		if game.CheckIfPossible(config) {
			possible = append(possible, game)
		}
	}

	sum := 0
	for _, game := range possible {
		sum += game.ID
	}

	fmt.Println("Part 1:", sum)
}

func Part2() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	var games []*Game
	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	var smallestConfigs []Configuration
	for _, game := range games {
		smallestConfigs = append(smallestConfigs, game.SmallestPossible())
	}

	sum := 0
	for _, config := range smallestConfigs {
		sum += config["red"] * config["green"] * config["blue"]
	}

	fmt.Println("Part 2:", sum)
}

func Run() {
	fmt.Println("===== Day 2 =====")
	Part1()
	Part2()
}

func ReadFile() string {
	input, err := os.ReadFile("2023/day2/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

var Day = year.Day{
	Number: 2,
	Run:    Run,
}
