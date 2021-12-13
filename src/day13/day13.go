package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type foldinstruction struct {
	axis  string
	value int
}

func main() {
	//scanner := openFile("ex13.txt")
	scanner := openFile("day13.txt")
	coords := [][]int{}
	xMax, yMax := 0, 0
	fis := []foldinstruction{}

	//Read file; get points and folding instructions
	for scanner.Scan() {
		xMax, yMax, coords, fis = getInput(scanner, xMax, yMax, coords, fis)
	}

	//Create the grid
	grid := newGrid(yMax, xMax, coords)

	//Start Folding
	for _, instruction := range fis {
		grid = fold(instruction, grid)
	}
	//Print Result
	printGrid(grid)
	//Print number of marked points
	println(countMarked(grid))

}

func countMarked(grid [][]int) (marked int) {
	for _, row := range grid {
		for _, v := range row {
			if v > 0 {
				marked++
			}
		}
	}
	return marked
}

func fold(fis foldinstruction, grid [][]int) (returngrid [][]int) {
	if fis.axis == "y" {
		foldA := grid[:fis.value]
		foldT := grid[fis.value+1:]
		foldB := [][]int{}
		for i := len(foldT) - 1; i >= 0; i-- {
			foldB = append(foldB, foldT[i])
		}

		for i := 0; i < len(foldB); i++ {
			for j := 0; j < len(foldB[i]); j++ {
				if foldA[i][j] != foldB[i][j] {
					foldA[i][j] = 1
				}
			}
		}
		returngrid = foldA

	}
	if fis.axis == "x" {
		foldA := [][]int{}
		foldT := [][]int{}
		foldB := [][]int{}
		for _, row := range grid {
			foldA = append(foldA, row[:fis.value])
			foldT = append(foldT, row[fis.value+1:])
		}

		for _, row := range foldT {
			b := []int{}
			for i := len(row) - 1; i >= 0; i-- {
				b = append(b, row[i])
			}
			foldB = append(foldB, b)
		}

		for i := 0; i < len(foldB); i++ {
			for j := 0; j < len(foldB[i]); j++ {
				if foldA[i][j] != foldB[i][j] {
					foldA[i][j] = 1
				}
			}
		}
		returngrid = foldA
	}
	return returngrid
}

func getInput(scanner *bufio.Scanner, xMax int, yMax int, coords [][]int, fis []foldinstruction) (int, int, [][]int, []foldinstruction) {
	line := scanner.Text()
	xyAsStringo := strings.Split(line, ",")

	if len(xyAsStringo) > 1 {

		x, _ := strconv.Atoi(xyAsStringo[0])
		y, _ := strconv.Atoi(xyAsStringo[1])

		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}

		coords = append(coords, []int{x, y})
	} else {
		if line != "" {
			words := strings.Fields(line)
			foldinst := words[2]
			a, _ := strconv.Atoi(foldinst[2:])
			fis = append(fis, foldinstruction{
				axis:  string(foldinst[0]),
				value: a,
			})
		}
	}
	return xMax, yMax, coords, fis
}

func newGrid(yMax int, xMax int, coords [][]int) [][]int {
	var grid [][]int
	for i := 0; i < yMax+1; i++ {
		grid = append(grid, make([]int, xMax+1))
	}

	for _, row := range coords {
		grid[row[1]][row[0]] = 1
	}
	return grid
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, v := range row {
			if v > 0 {
				print("#")
			} else {
				print(" ")
			}
		}
		print("\n")
	}

}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		println("Err opening file")
	}
	return
}
