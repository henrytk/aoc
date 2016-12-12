package main

import (
	. "../01"
	"testing"
)

func TestBunnyInstructionParser(t *testing.T) {
	bip := NewBunnyInstructionParser()
	if instructions := bip.Parse("R1, L2"); len(instructions) != 2 {
		t.Errorf("Expected 'R1, L2' to yield a list containing two instructions, found %v", len(instructions))
	}
	if instructions := bip.Parse("R1, L2"); instructions[0].Direction != "R" {
		t.Errorf("Expected 'R1, L2' to yield a list in which the first item's Direction is R, found %v", instructions[0].Direction)
	}
	if instructions := bip.Parse("R1, L2"); instructions[0].Steps != int(1) {
		t.Errorf("Expected 'R1, L2' to yield a list in which the first item's Steps is 1, found %v", instructions[0].Steps)
	}
}

func TestTaxi(t *testing.T) {
	taxi := NewTaxi()
	ins := "R1"
	taxi.FollowInstructions(ins)
	if taxi.Position.X != int(1) {
		t.Errorf("Expected R1 to leave taxi in position X=1, X=%v", taxi.Position.X)
	}
	if d := taxi.Distance(); d != float64(1.0) {
		t.Errorf("Expected R1 to yield distance travelled of 1, got %v", d)
	}
}

func TestCase1(t *testing.T) {
	taxi := NewTaxi()
	ins := "R2, L3"
	taxi.FollowInstructions(ins)
	expectedDistance := float64(5)
	if d := taxi.Distance(); d != expectedDistance {
		t.Errorf("Expected %v to yield distance travelled of %v, got %v", ins, expectedDistance, d)
	}
}

func TestCase2(t *testing.T) {
	taxi := NewTaxi()
	ins := "R2, R2, R2"
	taxi.FollowInstructions(ins)
	expectedDistance := float64(2)
	if d := taxi.Distance(); d != expectedDistance {
		t.Errorf("Expected %v to yield distance travelled of %v, got %v", ins, expectedDistance, d)
	}
}

func TestCase3(t *testing.T) {
	taxi := NewTaxi()
	ins := "R5, L5, R5, R3"
	taxi.FollowInstructions(ins)
	expectedDistance := float64(12)
	if d := taxi.Distance(); d != expectedDistance {
		t.Errorf("Expected %v to yield distance travelled of %v, got %v", ins, expectedDistance, d)
	}
}

func TestCase4(t *testing.T) {
	taxi := NewTaxi()
	ins := "R8, R4, R4, R8"
	taxi.FollowInstructions(ins)
	expectedDistance := float64(4)
	if d := taxi.DistanceFromFirstRepeatedPosition(); d != expectedDistance {
		t.Errorf("Expected %v to yield distance travelled of %v, got %v", ins, expectedDistance, d)
	}
}
