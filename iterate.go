package main

import (
	"math"
	"sync"
)

func iterate() {
	wg := sync.WaitGroup{}

	// Update locations

	wg.Add(len(points))

	for _, p := range points {
		go func(p *Point) {
			move(p)
			wg.Done()
		}(p)
	}

	wg.Wait()

	// Update speeds

	wg.Add(len(points))

	for _, p1 := range points {
		go func(p *Point) {
			accelerate(p)
			wg.Done()
		}(p1)
	}

	wg.Wait()
}

func move(point *Point) {
	point.x = point.x + point.vx
	point.y = point.y + point.vy
	if point.x > Width || point.x < 0 {
		point.vx = -point.vx
	}
	if point.y > Height || point.y < 0 {
		point.vy = -point.vy
	}
}

func accelerate(p1 *Point) {
	fx, fy := 0.0, 0.0
	for _, p2 := range points {
		if p1 != p2 {
			cfx, cfy := interact(p1, p2)
			fx += cfx
			fy += cfy
		}
	}
	p1.vx = (p1.vx + fx) * SlowingCoeffient
	p1.vy = (p1.vy + fy) * SlowingCoeffient

	// Max speed
	if p1.vx > MaxSpeed {
		p1.vx = MaxSpeed
	} else if p1.vx < -MaxSpeed {
		p1.vx = -MaxSpeed
	}
	if p1.vy > MaxSpeed {
		p1.vy = MaxSpeed
	} else if p1.vy < -MaxSpeed {
		p1.vy = -MaxSpeed
	}
}

func interact(p1, p2 *Point) (float64, float64) {
	var fx, fy float64

	g, exists := rules[p1.color][p2.color]
	if !exists {
		return 0, 0
	}

	dx := p1.x - p2.x
	dy := p1.y - p2.y
	d := math.Sqrt(dx*dx + dy*dy)

	if d > 0 && d < MaxDistance {
		F := g / d
		fx = dx * F
		fy = dy * F
	}

	return fx, fy
}
