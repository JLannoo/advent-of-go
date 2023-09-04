package day9

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

const (
	UP    = "U"
	DOWN  = "D"
	LEFT  = "L"
	RIGHT = "R"
)

var DIRECTIONS = map[string]Coordinate{
	UP:    {X: 0, Y: 1},
	DOWN:  {X: 0, Y: -1},
	LEFT:  {X: -1, Y: 0},
	RIGHT: {X: 1, Y: 0},
}

var SQRT_2 = math.Sqrt(2)

var DEBUG = false

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Distance(other Coordinate) float64 {
	return math.Sqrt(math.Pow(float64(c.X-other.X), 2) + math.Pow(float64(c.Y-other.Y), 2))
}

func (c *Coordinate) Diff(other Coordinate) Coordinate {
	return Coordinate{
		X: other.X - c.X,
		Y: other.Y - c.Y,
	}
}

type Knot struct {
	Position         Coordinate
	PreviousPosition Coordinate
	VisitedPositions mapset.Set
}

type Rope struct {
	Knots []Knot

	Head *Knot
	Tail *Knot

	maxX int
	maxY int
	minX int
	minY int
}

func (rope *Rope) Print() {
	for y := rope.maxY; y >= rope.minY; y-- {
		for x := rope.minX; x <= rope.maxX; x++ {
			isKnot := -1
			for i, knot := range rope.Knots {
				if knot.Position.X == x && knot.Position.Y == y {
					isKnot = i
					break
				}
			}

			if isKnot == -1 {
				if x == 0 && y == 0 {
					fmt.Print("s")
				} else {
					fmt.Print(".")
				}
			} else {
				if isKnot == 0 {
					fmt.Print("H")
				} else if isKnot == len(rope.Knots)-1 {
					fmt.Print("T")
				} else {
					fmt.Print(isKnot)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (rope *Rope) Move(direction string, distance string) {
	amount, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}

	for i := 0; i < amount; i++ {
		direction := DIRECTIONS[direction]
		rope.MoveTail(direction)

		if DEBUG {
			rope.UpdateSize()
			fmt.Println("direction:", direction)
			rope.Print()
		}
	}

}

func (rope *Rope) UpdateSize() {
	if rope.maxX < rope.Head.Position.X {
		rope.maxX = rope.Head.Position.X
	}
	if rope.maxY < rope.Head.Position.Y {
		rope.maxY = rope.Head.Position.Y
	}
	if rope.minX > rope.Head.Position.X {
		rope.minX = rope.Head.Position.X
	}
	if rope.minY > rope.Head.Position.Y {
		rope.minY = rope.Head.Position.Y
	}
}

func (rope *Rope) MoveTail(direction Coordinate) {
	for i := 0; i < len(rope.Knots); i++ {
		if i == 0 {
			rope.Knots[i] = *rope.Head
			rope.Head.PreviousPosition = rope.Head.Position
			rope.Head.Position.X += direction.X
			rope.Head.Position.Y += direction.Y
			rope.Head.VisitedPositions.Add(rope.Head.Position)
		} else {
			previousKnot := rope.Knots[i-1]
			interKnotDistance := previousKnot.Position.Distance(rope.Knots[i].Position)

			if DEBUG {
				fmt.Println("interKnotDistance:", interKnotDistance)
			}

			if interKnotDistance > SQRT_2 {
				diffVector := rope.Knots[i].Position.Diff(previousKnot.Position)
				// There's probably a better way to clamp these values...
				diffVector.X = clamp(diffVector.X, -1, 1)
				diffVector.Y = clamp(diffVector.Y, -1, 1)

				rope.Knots[i].PreviousPosition = rope.Knots[i].Position
				rope.Knots[i].Position = Coordinate{
					X: rope.Knots[i].Position.X + diffVector.X,
					Y: rope.Knots[i].Position.Y + diffVector.Y,
				}
				rope.Knots[i].VisitedPositions.Add(rope.Knots[i].Position)
			}
		}
	}
}

func Part1() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	knots := createKnots(2)

	rope := Rope{
		Knots: knots,
		Head:  &knots[0],
		Tail:  &knots[len(knots)-1],
		maxX:  0,
		maxY:  0,
	}

	for _, instruction := range lines {
		split := strings.Split(instruction, " ")
		direction := split[0]
		distance := split[1]

		rope.Move(direction, distance)
	}

	fmt.Println("Part 1:", rope.Tail.VisitedPositions.Cardinality())
}

func Part2() {
	input := ReadFile()
	lines := strings.Split(input, "\n")

	knots := createKnots(10)

	rope := Rope{
		Knots: knots,
		Head:  &knots[0],
		Tail:  &knots[len(knots)-1],
		maxX:  0,
		maxY:  0,
	}

	for _, instruction := range lines {
		split := strings.Split(instruction, " ")
		direction := split[0]
		distance := split[1]

		rope.Move(direction, distance)
	}

	fmt.Println("Part 2:", rope.Tail.VisitedPositions.Cardinality())
}

func ReadFile() string {
	input, err := os.ReadFile("day9/input.txt")

	if err != nil {
		panic(err)
	}

	return string(input)
}

func Run() {
	fmt.Println("===== Day 9 =====")
	Part1()
	Part2()

	fmt.Println()
}

func createKnots(n int) []Knot {
	knots := make([]Knot, n)
	for i := 0; i < len(knots); i++ {
		knots[i] = Knot{
			Position:         Coordinate{X: 0, Y: 0},
			PreviousPosition: Coordinate{X: 0, Y: 0},
			VisitedPositions: mapset.NewSet(),
		}
		knots[i].VisitedPositions.Add(knots[i].Position)
	}

	return knots
}

func clamp(n int, min int, max int) int {
	if n < min {
		return min
	}

	if n > max {
		return max
	}

	return n
}
