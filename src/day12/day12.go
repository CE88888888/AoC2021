package main

import (
	"bufio"
	"os"
	"strings"
)

type Point struct {
	p string
	v int
}

type Path struct {
	p1 string
	p2 string
}

func main() {
	//scanner := openFile("ex12.txt")
	scanner := openFile("day12.txt")
	paths := initPaths(scanner)

	start := Point{p: "start", v: 8}

	routes := [][]Point{}
	route := []Point{start}
	routes = append(routes, route)

	for i := 0; i < len(routes); i++ {
		notEnd := true
		for notEnd {
			p1 := routes[i]
			current := p1[len(p1)-1]

			if current.p == strings.ToLower(current.p) && current.p != "start" {
				p1[len(p1)-1].v = 1
			}
			if current.p == "end" {
				notEnd = false
			}

			possibles := getNextPoints(current, paths, routes[i])

			if len(possibles) == 0 {
				notEnd = false
			} else {
				for j := 1; j < len(possibles); j++ {
					newPad := make([]Point, len(p1))
					copy(newPad, p1)
					newPad = append(newPad, possibles[j])
					routes = append(routes, newPad)
				}
				routes[i] = append(routes[i], possibles[0])
			}
		}
	}

	endingPaths := [][]Point{}
	for i := range routes {
		lastpoint := routes[i][len(routes[i])-1]
		if lastpoint.p == "end" {
			endingPaths = append(endingPaths, routes[i])
		}
	}

	println(len(endingPaths))

}

func initPaths(scanner *bufio.Scanner) (paths []Path) {
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), "-")
		p := Path{p1: points[0], p2: points[1]}

		if p.p1 == "end" {
			pr := Path{p1: p.p2, p2: p.p1}
			paths = append(paths, pr)
		} else if p.p2 == "end" {
			paths = append(paths, p)
		} else {
			paths = append(paths, p)
			pr := Path{p1: p.p2, p2: p.p1}
			paths = append(paths, pr)
		}
	}
	return paths
}

func checkTwiceVisetedYet(p1 []Point, c string) bool {
	s := ""
	for _, p := range p1 {
		if !(p.p == "start" || p.p == "end") && p.p == strings.ToLower(p.p) {
			s += p.p
		}
	}
	for _, p := range p1 {
		if strings.Count(s, p.p) > 1 {
			return true
		}
	}
	if c == "start" {
		return c == "start"
	}
	return false
}

func getNextPoints(current Point, paths []Path, pad []Point) (nextPoints []Point) {
	for _, value := range paths {
		if current.p == value.p1 {
			valid := true
			for _, point := range pad {
				if point.p == value.p2 && point.v > 0 {
					valid = false
					if !checkTwiceVisetedYet(pad, value.p2) {
						valid = true
					}
				}
			}
			if valid {
				nextPoints = append(nextPoints, Point{p: value.p2, v: 0})
			}
		}
	}
	return
}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		println("Err opening file")
	}
	return
}
