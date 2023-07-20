package day2

import (
	"fmt"
	"os"
	"strings"
)

type POINTS struct {
	WIN  int
	DRAW int
	LOSS int
}

var RESULT_POINTS_MAP = POINTS{
	WIN:  6,
	DRAW: 3,
	LOSS: 0,
}

var SHAPE_POINTS_MAP = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var SHAPES_WIN_MAP = map[string]string{
	"rock":     "scissors",
	"paper":    "rock",
	"scissors": "paper",
}

func Part1() {
	input := readFile()

	var MY_SHAPES = map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	var OPPONENT_SHAPES = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	instruction_pairs := getInstructionParis(input)

	total_score := 0
	for _, pair := range instruction_pairs {
		opponent := OPPONENT_SHAPES[pair[0]]
		me := MY_SHAPES[pair[1]]

		total_score += SHAPE_POINTS_MAP[me]

		switch me {
		// LOSS
		case SHAPES_WIN_MAP[opponent]:
			total_score += RESULT_POINTS_MAP.LOSS
		// DRAW
		case opponent:
			total_score += RESULT_POINTS_MAP.DRAW
		// WIN
		default:
			total_score += RESULT_POINTS_MAP.WIN
		}
	}

	println("Part 1: ", total_score)
}

func Part2() {
	input := readFile()

	var OPPONENT_SHAPES = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	var RESULT_MAP = map[string]string{
		"X": "LOSS",
		"Y": "DRAW",
		"Z": "WIN",
	}

	instruction_pairs := getInstructionParis(input)

	total_score := 0
	for _, pair := range instruction_pairs {
		opponent := OPPONENT_SHAPES[pair[0]]
		expected_result := RESULT_MAP[pair[1]]

		switch expected_result {
		case "WIN":
			total_score += RESULT_POINTS_MAP.WIN

			for winner, loser := range SHAPES_WIN_MAP {
				if loser == opponent {
					total_score += SHAPE_POINTS_MAP[winner]
				}
			}
		case "DRAW":
			total_score += RESULT_POINTS_MAP.DRAW
			total_score += SHAPE_POINTS_MAP[opponent]
		case "LOSS":
			total_score += RESULT_POINTS_MAP.LOSS

			for winner, loser := range SHAPES_WIN_MAP {
				if winner == opponent {
					total_score += SHAPE_POINTS_MAP[loser]
				}
			}
		default:
			panic("Uh oh")
		}
	}

	println("Part 2: ", total_score)
}

func readFile() string {
	input, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	return string(input)
}

func getInstructionParis(input string) [][]string {
	instruction_pairs := [][]string{}
	for _, instruction_pair := range strings.Split(input, "\n") {
		instruction_pairs = append(instruction_pairs, strings.Split(instruction_pair, " "))
	}

	return instruction_pairs
}

func Run() {
	fmt.Println("===== Day 2 =====")
	Part1()
	Part2()

	fmt.Println()
}
