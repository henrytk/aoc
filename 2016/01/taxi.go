package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Steps     int
}

type InstructionParser interface {
	Parse(string) []Instruction
}

type BunnyInstructionParser struct{}

func NewBunnyInstructionParser() *BunnyInstructionParser {
	return &BunnyInstructionParser{}
}

func (bip *BunnyInstructionParser) Parse(in string) []Instruction {
	var instructions []Instruction
	split := strings.Split(in, ", ")
	for _, v := range split {
		steps, _ := strconv.ParseInt(v[1:], 10, 64)
		instructions = append(instructions, Instruction{Direction: v[:1], Steps: int(steps)})
	}
	return instructions
}

type Compass interface {
	Turn(string) Compass
}

type North struct{}
type East struct{}
type South struct{}
type West struct{}

func (n *North) Turn(dir string) Compass {
	if dir == "R" {
		return &East{}
	}
	return &West{}
}

func (e *East) Turn(dir string) Compass {
	if dir == "R" {
		return &South{}
	}
	return &North{}
}

func (s *South) Turn(dir string) Compass {
	if dir == "R" {
		return &West{}
	}
	return &East{}
}

func (w *West) Turn(dir string) Compass {
	if dir == "R" {
		return &North{}
	}
	return &South{}
}

type Point struct {
	X int
	Y int
}

type Taxi struct {
	Position  Point
	Direction Compass
	BeenTo    []Point
}

func NewTaxi() Taxi {
	return Taxi{
		Position:  Point{X: 0, Y: 0},
		Direction: &North{},
	}
}

func (t *Taxi) FollowInstructions(instructions string) {
	bip := NewBunnyInstructionParser()
	in := bip.Parse(instructions)
	for _, i := range in {
		t.Direction = t.Direction.Turn(i.Direction)
		for x := 0; x < i.Steps; x++ {
			t.Drive()
		}
	}

}

func (t *Taxi) Drive() {
	switch t.Direction.(type) {
	case *North:
		t.Position.Y = t.Position.Y + 1
	case *East:
		t.Position.X = t.Position.X + 1
	case *South:
		t.Position.Y = t.Position.Y - 1
	default:
		t.Position.X = t.Position.X - 1
	}
	t.BeenTo = append(t.BeenTo, t.Position)
}

func (t *Taxi) Distance() float64 {
	return distanceFromPoint(t.Position)
}

func distanceFromPoint(p Point) float64 {
	x := math.Abs(float64(p.X))
	y := math.Abs(float64(p.Y))
	return x + y
}

func (t *Taxi) DistanceFromFirstRepeatedPosition() float64 {
	positions := make(map[Point]bool)
	for _, location := range t.BeenTo {
		if _, ok := positions[location]; ok {
			return distanceFromPoint(location)
		}
		positions[location] = true
	}
	return float64(0)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "You must pass instructions as the first and only argument, you passed %v\n", len(os.Args))
		os.Exit(1)
	}
	taxi := NewTaxi()
	taxi.FollowInstructions(os.Args[1])
	fmt.Fprintf(os.Stdout, "You have travelled a distance of %v\n", taxi.Distance())
	fmt.Fprintf(os.Stdout, "The Easter Bunny HQ is %v blocks away\n", taxi.DistanceFromFirstRepeatedPosition())
}
