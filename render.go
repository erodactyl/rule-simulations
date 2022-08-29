package main

import "github.com/tfriedel6/canvas"

func render(cv *canvas.Canvas) {
	w, h := float64(cv.Width()), float64(cv.Height())
	cv.SetFillStyle("#fff")
	cv.FillRect(0, 0, w, h)

	for _, point := range points {
		cv.SetFillStyle(string(point.color))
		cv.FillRect(point.x, point.y, 8, 8)
	}
}
