package main

import (
	"bufio"
	"os"
)

type Pos struct {
	r       int
	c       int
	flashed bool
}

func main() {
	grid := initOctoGrid("day11.txt")
	flashamount := 0
	allFlashedStep := 0

Synced:
	for i := 0; i < 500; i++ {

		//Step 1 Increase Energy
		increaseGridEnergy(grid)

		//Step 2 Check for flashes
		flashedPositions := []Pos{}

		for rowkey, row := range grid {
			for colkey, e := range row {
				if e > 9 {
					//add to flash list
					flashedPositions = append(flashedPositions, Pos{r: rowkey, c: colkey, flashed: false})
				}
			}
		}

		//Flash and increase adjacents
		for i := 0; i < len(flashedPositions); i++ {
			p := flashedPositions[i]
			flashedPositions[i].flashed = true
			adjacents := getAdjacentPos(p)
			for _, v := range adjacents {
				grid[v.r][v.c]++
				if grid[v.r][v.c] > 9 {
					if !checkIfAlreadyFlashed(&flashedPositions, v.r, v.c) {
						flashedPositions = append(flashedPositions, Pos{r: v.r, c: v.c, flashed: false})
					}
				}
			}
		}

		flashamount += len(flashedPositions)
		if len(flashedPositions) == 100 {
			allFlashedStep = i
			break Synced
		}
		//Step 3 Set flashed Octos to 0
		resetFlashedOctos(grid, flashedPositions)
	}
	println("done, amount flashed: ", flashamount)
	println("all synced up at step: ", allFlashedStep+1)
}

func getAdjacentPos(p Pos) (adjacents []Pos) {
	mr, mc := 9, 9
	r, c := p.r, p.c
	if r > 0 && c > 0 {
		adjacents = append(adjacents, Pos{r: r - 1, c: c - 1, flashed: false})
	}
	if r > 0 {
		adjacents = append(adjacents, Pos{r: r - 1, c: c, flashed: false})
	}
	if r > 0 && c < mc {
		adjacents = append(adjacents, Pos{r: r - 1, c: c + 1, flashed: false})
	}
	if c > 0 {
		adjacents = append(adjacents, Pos{r: r, c: c - 1, flashed: false})
	}
	if c < mc {
		adjacents = append(adjacents, Pos{r: r, c: c + 1, flashed: false})
	}
	if r < mr && c > 0 {
		adjacents = append(adjacents, Pos{r: r + 1, c: c - 1, flashed: false})
	}
	if r < mr {
		adjacents = append(adjacents, Pos{r: r + 1, c: c, flashed: false})
	}
	if r < mr && c < mc {
		adjacents = append(adjacents, Pos{r: r + 1, c: c + 1, flashed: false})
	}

	return adjacents
}

func checkIfAlreadyFlashed(fp *[]Pos, r, c int) (exists bool) {
	for _, pos := range *fp {
		if pos.r == r && pos.c == c {
			return true
		}
	}
	return false
}

func increaseGridEnergy(grid [][]int) {
	for rowkey, row := range grid {
		for colkey := range row {
			grid[rowkey][colkey]++
		}
	}
}
func resetFlashedOctos(grid [][]int, fps []Pos) {
	for _, p := range fps {
		grid[p.r][p.c] = 0
	}
}

func initOctoGrid(f string) (octogrid [][]int) {
	scanner := openFile(f)
	for scanner.Scan() {
		octorow := []int{}
		for _, c := range scanner.Text() {
			octorow = append(octorow, int(c-'0'))
		}
		octogrid = append(octogrid, octorow)
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
