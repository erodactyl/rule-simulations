package main

import (
	"fmt"
	"math/rand"
	"time"
)

var points = []*Point{}

type Atom string

const Yellow, Red, Blue, Green Atom = "#FFEB3B", "#F44336", "#3F51B5", "#009688"

type Point struct {
	x, y, vx, vy float64
	color        Atom
}

var rules = map[Atom]map[Atom]float64{} // rules[Yellow][Red] = 3 means Yellow attracts Red with a force of 3

func setup() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < Count; i++ {
		x, y := rand.Float64()*Width, rand.Float64()*Height
		point := &Point{x, y, 0.0, 0.0, randColor()}
		points = append(points, point)
	}

	addRule(Green, Green, randG())
	addRule(Green, Red, randG())
	addRule(Green, Yellow, randG())
	addRule(Green, Blue, randG())

	addRule(Red, Red, randG())
	addRule(Red, Green, randG())
	addRule(Red, Yellow, randG())
	addRule(Red, Blue, randG())

	addRule(Yellow, Yellow, randG())
	addRule(Yellow, Green, randG())
	addRule(Yellow, Red, randG())
	addRule(Yellow, Blue, randG())

	addRule(Blue, Blue, randG())
	addRule(Blue, Green, randG())
	addRule(Blue, Red, randG())
	addRule(Blue, Yellow, randG())
}

func randG() float64 {
	return 0.5 - rand.Float64()
}

func addRule(atom1, atom2 Atom, force float64) {
	if _, exists := rules[atom1]; !exists {
		rules[atom1] = map[Atom]float64{}
	}
	rules[atom1][atom2] = force
}

func randColor() Atom {
	idx := rand.Intn(4)
	switch idx {
	case 0:
		return Yellow
	case 1:
		return Red
	case 2:
		return Green
	case 3:
		return Blue
	default:
		panic(fmt.Sprintf("color for index %d does not exist", idx))
	}
}
