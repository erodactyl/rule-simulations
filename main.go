package main

import (
	"github.com/tfriedel6/canvas/glfwcanvas"
)

var Count = 1000
var MaxSpeed float64 = 10
var MaxDistance float64 = 100
var SlowingCoeffient = 0.9

var Height, Width float64 = 500 * 2, 500 * 2

func main() {
	wnd, cv, err := glfwcanvas.CreateWindow(int(Width/2), int(Height/2), "Rule Simulation")
	if err != nil {
		panic(err)
	}
	defer wnd.Window.Destroy()

	setup()

	wnd.MainLoop(func() {
		iterate()
		render(cv)
	})
}
