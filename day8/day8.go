package day8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Day 8: Treetop Tree House
	https://adventofcode.com/2022/day/8

	- Input is a grid of numbers, 0-9
	- A tree is hidden if it is blocked by another tree from the edge

	Part 1:
	How many trees are visible?

	Part 2:
	- Calculate orthogonal viewing distances
		- Viewing distance is calculated by going in every direction until you hit a tree or the edge
	- Calculate the `scenic score` for each tree
		- The scenic score is the product of the viewing distances in each direction

	- Find the tree with the highest scenic score
*/

type Tree struct {
	Height    int
	X         int
	Y         int
	IsVisible bool
}

func parseTrees(input string) [][]Tree {
	lines := strings.Split(input, "\n")

	grid := make([][]Tree, len(lines))

	for y, line := range lines {
		grid[y] = make([]Tree, len(line))

		for x, char := range line {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			grid[y][x] = Tree{Height: number, X: x, Y: y}
		}
	}

	return grid
}

type Grid struct {
	Size  int
	Trees [][]Tree
}

func (grid *Grid) CalculateVisibility() {
	for i, row := range grid.Trees {
		for j, tree := range row {
			tree.IsVisible = grid.isVisible(tree)
			grid.Trees[i][j] = tree
		}
	}
}

func (grid *Grid) isVisible(tree Tree) bool {
	if tree.X == 0 || tree.X == grid.Size-1 {
		return true
	}
	if tree.Y == 0 || tree.Y == grid.Size-1 {
		return true
	}

	// Don't like this solutions very much, but it works
	coveredLeft := false
	coveredRight := false
	coveredUp := false
	coveredDown := false

	for i := 0; i < tree.X; i++ {
		if grid.Trees[tree.Y][i].Height >= tree.Height {
			coveredLeft = true
			break
		}
	}

	for i := tree.X + 1; i < grid.Size; i++ {
		if grid.Trees[tree.Y][i].Height >= tree.Height {
			coveredRight = true
			break
		}
	}

	for i := 0; i < tree.Y; i++ {
		if grid.Trees[i][tree.X].Height >= tree.Height {
			coveredUp = true
			break
		}
	}

	for i := tree.Y + 1; i < grid.Size; i++ {
		if grid.Trees[i][tree.X].Height >= tree.Height {
			coveredDown = true
			break
		}
	}

	return !(coveredLeft && coveredRight && coveredUp && coveredDown)
}

func (grid *Grid) CalculateScenicScore(tree Tree) int {
	scenicScore := 0

	// Don't like this solutions very much, but it works
	treeCountUp := 0
	treeCountDown := 0
	treeCountLeft := 0
	treeCountRight := 0

	for x := tree.X - 1; x >= 0; x-- {
		treeCountLeft++
		if grid.Trees[tree.Y][x].Height >= tree.Height {
			break
		}
	}

	for x := tree.X + 1; x < grid.Size; x++ {
		treeCountRight++
		if grid.Trees[tree.Y][x].Height >= tree.Height {
			break
		}
	}

	for y := tree.Y - 1; y >= 0; y-- {
		treeCountUp++
		if grid.Trees[y][tree.X].Height >= tree.Height {
			break
		}
	}

	for y := tree.Y + 1; y < grid.Size; y++ {
		treeCountDown++
		if grid.Trees[y][tree.X].Height >= tree.Height {
			break
		}
	}

	scenicScore = treeCountUp * treeCountDown * treeCountLeft * treeCountRight
	return scenicScore
}

func (grid *Grid) Print(highlightX int, highlightY int) {
	for _, row := range grid.Trees {
		for _, tree := range row {
			isHighlighted := tree.X == highlightX && tree.Y == highlightY

			if isHighlighted {
				fmt.Print("\033[44m")
			}

			if tree.IsVisible {
				fmt.Print("X")
			} else {
				fmt.Print("O")
			}

			if isHighlighted {
				fmt.Print("\033[0m")
			}
		}
		fmt.Println()
	}
	
	fmt.Println()
}

func (grid *Grid) Count() (int, int) {
	visible := 0
	hidden := 0

	for _, row := range grid.Trees {
		for _, tree := range row {
			if tree.IsVisible {
				visible++
			} else {
				hidden++
			}
		}
	}

	return visible, hidden
}

func Part1() {
	input := ReadFile()

	trees := parseTrees(input)
	grid := Grid{Trees: trees, Size: len(trees)}

	grid.CalculateVisibility()

	visible, _ := grid.Count()

	fmt.Println("Part 1:", visible)
}

func Part2() {
	input := ReadFile()

	trees := parseTrees(input)
	grid := Grid{Trees: trees, Size: len(trees)}

	max := 0
	for _, row := range grid.Trees {
		for _, tree := range row {
			scenicScore := grid.CalculateScenicScore(tree)
			if scenicScore > max {
				max = scenicScore
			}
		}
	}

	fmt.Println("Part 2:", max)
}

func ReadFile() string {
	input, err := os.ReadFile("day8/input.txt")
	if err != nil {
		panic(err)
	}

	return string(input)
}

func Run() {
	fmt.Println("==== Day 8 ====")
	Part1()
	Part2()

	fmt.Println()
}
