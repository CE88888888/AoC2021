package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	X1, Y1 int
	X2, Y2 int
}

func (l *Line) SortX() {
	if l.X1 > l.X2 {
		x := l.X1
		y := l.Y1
		l.X1, l.Y1 = l.X2, l.Y2
		l.X2, l.Y2 = x, y

	}
}

type Area struct {
	A            [][]int
	maxCollision int
}

func newArea(size int) *Area {
	slice := make([][]int, size)

	for i := 0; size > i; i++ {
		slice[i] = make([]int, size)
	}

	a := Area{
		A:            slice,
		maxCollision: 0,
	}
	return &a
}

func newLine(start string, end string) *Line {
	var coordinates []int
	split := strings.Split(start, ",")
	split = append(split, strings.Split(end, ",")...)

	for _, value := range split {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal()
		}
		coordinates = append(coordinates, i)
	}

	l := Line{
		X1: coordinates[0],
		Y1: coordinates[1],
		X2: coordinates[2],
		Y2: coordinates[3],
	}
	return &l
}

func drawLines(l Line, a Area) *Area {

	if l.X1 == l.X2 {
		ylow := l.Y1
		ymax := l.Y2
		if l.Y1 > l.Y2 {
			ylow = l.Y2
			ymax = l.Y1
		}
		for y := ylow; y < ymax+1; y++ {
			a.A[y][l.X1]++
			if a.maxCollision < a.A[y][l.X1] {
				a.maxCollision = a.A[y][l.X1]
			}
		}
	}
	if l.Y1 == l.Y2 {
		l.SortX()
		for x := l.X1; x < l.X2+1; x++ {
			a.A[l.Y1][x]++
			if a.maxCollision < a.A[l.Y1][x] {
				a.maxCollision = a.A[l.Y1][x]
			}
		}
	}

	if (l.X1 != l.X2) && (l.Y1 != l.Y2) {
		l.SortX()
		if l.Y1 < l.Y2 {
			for x, y := l.X1, l.Y1; x < l.X2+1; x, y = x+1, y+1 {
				a.A[y][x]++
				if a.maxCollision < a.A[y][x] {
					a.maxCollision = a.A[y][x]
				}
			}

		} else {
			for x, y := l.X1, l.Y1; x < l.X2+1; x, y = x+1, y-1 {
				a.A[y][x]++
				if a.maxCollision < a.A[y][x] {
					a.maxCollision = a.A[y][x]
				}
			}
		}
	}
	return &a
}

func main() {

	var allLines []Line
	area := newArea(1000)

	scanner := openFile()
	for scanner.Scan() {
		strline := ""
		strline = scanner.Text()
		pairs := strings.Fields(strline)
		firstPair := pairs[0]
		secondPair := pairs[2]

		nl := newLine(firstPair, secondPair)
		allLines = append(allLines, *nl)

	}

	for _, value := range allLines {
		area = drawLines(value, *area)
	}

	amountofMax := 0

	for _, row := range area.A {
		for _, value := range row {
			if value > 1 {
				amountofMax++
			}
		}
	}

	println(amountofMax)

}

func openFile() (scanner *bufio.Scanner) {
	if file, err := os.Open("day5.txt"); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		log.Fatal(err)
	}
	return
}
