package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	id        int
	x, y, z   int
	beacons   []Beacon
	distances []Distance
}

func (s *Scanner) calculateDistances() {
	s.distances = []Distance{}
	b := s.beacons
	for i := 0; i < len(b); i++ {
		for j := i; j < len(b); j++ {
			if i != j {
				d := Distance{
					source: &b[i],
					dest:   &b[j],
					x:      absDiffInt(b[i][0], b[j][0]),
					y:      absDiffInt(b[i][1], b[j][1]),
					z:      absDiffInt(b[i][2], b[j][2]),
				}
				d.product = d.x*d.x + d.y*d.y + d.z*d.z
				s.distances = append(s.distances, d)
			}
		}
	}
}

func (s *Scanner) compareBeacons(s1 Scanner, s1T *Scanner) {
	sBeacons := []Beacon{}
	s1Beacons := []Beacon{}
	//Find common beacons
	for i := 0; i < len(s.distances); i++ {
		for j := 0; j < len(s1.distances); j++ {
			d1 := s.distances[i]
			d2 := s1.distances[j]
			if d1.product == d2.product {
				sBeacons = addUniqueBeacon(sBeacons, *d1.source)
				sBeacons = addUniqueBeacon(sBeacons, *d1.dest)
				s1Beacons = addUniqueBeacon(s1Beacons, *d2.source)
				s1Beacons = addUniqueBeacon(s1Beacons, *d2.dest)
			}
		}
	}

	// Treshold 12 gives us enough confidence to try and find matches.
	if len(sBeacons) >= 12 {
	Find:
		for _, rt := range rotArray {
			for _, org := range s.beacons {
				for _, targetB := range s1.beacons {
					tryTransform := targetB.rotate(rt).minus(org)
					cnt := 0
					newBeacons := make([]Beacon, 0, len(s1.beacons)-12)
					for _, i := range s1.beacons {
						hit := false
						itrf := i.rotate(rt).add(tryTransform)
						for _, j := range sBeacons {
							if j.equals(itrf) {
								cnt += 1
								hit = true
							}
						}
						if !hit {
							newBeacons = append(newBeacons, itrf)
						}
					}
					if cnt > 11 {
						s1T.x = s.x + tryTransform[0]
						s1T.y = s.y + tryTransform[1]
						s1T.z = s.z + tryTransform[2]
						s.beacons = append(s.beacons, newBeacons...)
						break Find
					}
				}
			}
		}
	}
}

func addUniqueBeacon(sBeacons []Beacon, b Beacon) []Beacon {
	for _, v := range sBeacons {
		if v[0] == b[0] && v[1] == b[1] && v[2] == b[2] {
			return sBeacons
		}
	}
	sBeacons = append(sBeacons, b)
	return sBeacons
}

type Beacon []int

type rotation struct {
	mp  Beacon
	inv Beacon
}

func (b Beacon) add(t Beacon) (nv Beacon) {
	return Beacon{b[0] + t[0], b[1] + t[1], b[2] + t[2]}
}

func (b Beacon) minus(target Beacon) Beacon {
	return Beacon{target[0] - b[0], target[1] - b[1], target[2] - b[2]}
}

func (b1 Beacon) equals(b2 Beacon) bool {
	return b2[0] == b1[0] && b2[1] == b1[1] && b2[2] == b1[2]
}

func (b Beacon) rotate(r rotation) (result Beacon) {
	result = Beacon{0, 0, 0}
	for i := 0; i < 3; i++ {
		result[i] = r.inv[i] * b[r.mp[i]]
	}
	return
}

type Distance struct {
	source  *Beacon
	dest    *Beacon
	x       int
	y       int
	z       int
	product int
}

var s0Adress *Scanner
var rotArray []rotation

func main() {
	//filescanner := openFile("ex19.txt")
	filescanner := openFile("day19.txt")
	scanners := []Scanner{}
	id := 0
	for filescanner.Scan() {
		line := filescanner.Text()
		if len(line) > 0 {
			if line[0:3] == "---" {
				//create scanners
				s := Scanner{id: id, beacons: []Beacon{}}
				scanners = append(scanners, s)
				id++
			} else {
				//create beacon
				points := strings.Split(line, ",")
				x, _ := strconv.Atoi(points[0])
				y, _ := strconv.Atoi(points[1])
				z, _ := strconv.Atoi(points[2])
				scanners[len(scanners)-1].beacons = append(scanners[len(scanners)-1].beacons, Beacon{x, y, z})
			}
		}
	}

	undiscovereds := []*Scanner{}
	for key := range scanners {
		scanners[key].calculateDistances()
		undiscovereds = append(undiscovereds, &scanners[key])
	}
	undiscovereds = undiscovereds[1:]
	scanners[0].x = 0
	scanners[0].y = 0
	scanners[0].z = 0
	s0Adress = &scanners[0]

	rotArray = initRotation()

	//	countleft := 0
	for len(undiscovereds) > 0 {
		for j := 0; j < len(undiscovereds); j++ {
			if undiscovereds[j].x == 0 {
				s0Adress.compareBeacons(*undiscovereds[j], undiscovereds[j])
				s0Adress.calculateDistances()
			}
		}

		temp := undiscovereds[:0]
		for i := 0; i < len(undiscovereds); i++ {
			if undiscovereds[i].x == 0 {
				temp = append(temp, undiscovereds[i])
			}
		}
		undiscovereds = temp
	}

	//Exercise A
	println(len(scanners[0].beacons))

	//Exercise B
	x := 0
	for i := 0; i < len(scanners); i++ {
		for j := 0; j < len(scanners); j++ {
			if i != j {
				mh := absDiffInt(scanners[i].x, scanners[j].x) + absDiffInt(scanners[i].y, scanners[j].y) + absDiffInt(scanners[i].z, scanners[j].z)
				if x < mh {
					x = mh
				}
			}
		}
	}
	println(x)

}

func initRotation() (r []rotation) {

	r = []rotation{}
	inv := []Beacon{
		{1, 1, 1},
		{1, 1, -1},
		{1, -1, 1},
		{-1, 1, 1},
		{-1, -1, 1},
		{-1, 1, -1},
		{1, -1, -1},
		{-1, -1, -1}}

	mp := []Beacon{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{2, 0, 1},
		{1, 2, 0},
		{2, 1, 0}}

	for _, i := range inv {
		for _, m := range mp {
			r = append(r, rotation{inv: i, mp: m})
		}
	}

	return
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		println("Err opening file")
	}
	return
}
