package main

import (
	"bufio"
	"os"
)

type Point struct {
	r     int
	c     int
	label string
	cost  int
}

var maxrow, maxcol int
var grid [][]int

func main() {
	// Read input file, collect grid
	scanner := openFile("day15.txt")
	//scanner := openFile("ex15.txt")
	for scanner.Scan() {
		line := scanner.Text()
		newrow := []int{}
		for _, v := range line {
			newrow = append(newrow, int(v-'0'))
		}
		grid = append(grid, newrow)
	}

	//Initialise
	start := Point{r: 0, c: 0, label: "next", cost: 0}
	paths := []Point{start}
	costtoend := 0

	//Exercise B init
	grid = enlargeGrid(&grid, 5)
	maxrow = len(grid) - 1
	maxcol = len(grid[0]) - 1

Find:
	for getNextPoint(paths).label != "" {
		current := getNextPoint(paths)

		if current.r == maxrow && current.c == maxcol {
			costtoend = current.cost
			break Find
		}

		current.label = "done"
		nbs := getNeighbours(*current)
		for _, nb := range nbs {
			if !checkIfExists(nb, paths) {
				nb.cost += current.cost
				paths = append(paths, nb)
			}
		}
		setNext(&paths)
	}
	print(costtoend)
}

func setNext(paths *[]Point) {
	lcPoint := &(Point{r: 0, c: 0, label: "", cost: 0})
	for k, v := range *paths {
		if v.label != "done" {
			if lcPoint.cost == 0 {
				lcPoint = &(*paths)[k]
			} else if v.cost < lcPoint.cost {
				lcPoint = &(*paths)[k]
			}
		}
	}
	lcPoint.label = "next"
}

func enlargeGrid(grid *[][]int, size int) [][]int {
	//Make a list of grids where index is the shift amount vs the original
	grids := [][][]int{}
	for i := 0; i < 10; i++ {
		if i == 0 {
			grids = append(grids, *grid)
		} else {
			grids = append(grids, incrementGrid(grids[i-1]))
		}
	}

	//Pick the correct grids and build the large grid for the first 'tile-row'
	gridLarge := [][]int{}
	for i := 0; i < size; i++ {
		for key := range grids[i] {
			for j := 1; j < size; j++ {
				grids[i][key] = append(grids[i][key], grids[i+j][key]...)
			}
		}
	}

	//Build the rest of the grid by using the enlarged picking list
	for i := 0; i < size; i++ {
		gridLarge = append(gridLarge, grids[i]...)
	}

	return gridLarge
}

func incrementGrid(grid [][]int) [][]int {
	gridB := [][]int{}
	for _, row := range grid {
		nr := []int{}
		for _, val := range row {
			if val < 9 {
				nr = append(nr, val+1)
			} else {
				nr = append(nr, 1)
			}
		}
		gridB = append(gridB, nr)
	}
	return gridB
}

func getNextPoint(pointlist []Point) *Point {
	for k, v := range pointlist {
		if v.label == "next" {
			return &pointlist[k]
		}
	}
	return &Point{}
}

func checkIfExists(nb Point, paths []Point) bool {
	for _, value := range paths {
		if value.r == nb.r && value.c == nb.c {
			return true
		}
	}
	return false
}

func getNeighbours(start Point) (nbs []Point) {
	nbs = []Point{}

	if start.c < maxcol {
		nbs = append(nbs, Point{r: start.r, c: start.c + 1, label: "new", cost: grid[start.r][start.c+1]})
	}
	if start.c > 0 {
		nbs = append(nbs, Point{r: start.r, c: start.c - 1, label: "new", cost: grid[start.r][start.c-1]})
	}
	if start.r < maxrow {
		nbs = append(nbs, Point{r: start.r + 1, c: start.c, label: "new", cost: grid[start.r+1][start.c]})
	}
	if start.r > 0 {
		nbs = append(nbs, Point{r: start.r - 1, c: start.c, label: "new", cost: grid[start.r-1][start.c]})
	}
	return nbs
}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		println("Err opening file")
	}
	return
}
