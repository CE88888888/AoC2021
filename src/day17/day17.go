package main

import "strconv"

type Probe struct {
	xpos, ypos int
	vx, vy     int
	steps      int
	maxheight  int
}

func (p *Probe) isInRange(t TargetArea) bool {
	if p.ypos >= t.ymin && p.ypos <= t.ymax {
		if p.xpos >= t.xmin && p.xpos <= t.xmax {
			return true
		}
	}
	return false
}

func (p *Probe) isAboveMin(t TargetArea) bool {

	return p.ypos >= t.ymin
}

func (p *Probe) Step() {
	p.xpos += p.vx
	p.ypos += p.vy

	if p.ypos > p.maxheight {
		p.maxheight = p.ypos
	}

	if p.vx > 0 {
		p.vx--
	} else if p.vx < 0 {
		p.vx++
	}
	p.vy--
	p.steps++
}

type TargetArea struct {
	xmin, xmax int
	ymin, ymax int
}

func NewTargetArea(x1, x2, y1, y2 int) *TargetArea {
	t := new(TargetArea)
	t.xmin = x1
	t.xmax = x2
	t.ymin = y1
	t.ymax = y2
	return t
}

func main() {

	ta := NewTargetArea(138, 184, -125, -71)

	println("Exercise A:", getMaxHeigtSpeed(ta))
	println("Exercise B:", len(ta.countLaunchValues()))

}

func (ta *TargetArea) countLaunchValues() []string {
	values := []string{}
	for xspeed := 1; xspeed <= ta.xmax; xspeed++ {
		for yspeed := ta.ymin; yspeed < (-1 * ta.ymin); yspeed++ {
			p := Probe{xpos: 0, ypos: 0, vx: xspeed, vy: yspeed, steps: 0, maxheight: -71}
			for !p.isInRange(*ta) && p.isAboveMin(*ta) {
				p.Step()
			}
			if p.isInRange(*ta) {
				value := strconv.Itoa(xspeed) + "," + strconv.Itoa(yspeed)
				values = append(values, value)
			}
		}
	}
	return values
}

func getMaxHeigtSpeed(ta *TargetArea) int {
	maxheight := 0
	busy := true
	yspeed := (ta.ymax * -1)
	xspeed := 17 // slowest speed (largest steps s) for which x = v * (v+1)/2 with x between xmin and xmax
	for busy {
		p := Probe{xpos: 0, ypos: 0, vx: xspeed, vy: yspeed, steps: 0, maxheight: -71}
		for !p.isInRange(*ta) && p.isAboveMin(*ta) {
			p.Step()
		}
		if p.isAboveMin(*ta) && p.maxheight > maxheight {
			maxheight = p.maxheight

		} else if p.isAboveMin(*ta) && p.maxheight < maxheight {
			busy = false
		}
		if !p.isAboveMin(*ta) {
			busy = false
		}
		yspeed++
	}
	return maxheight
}
