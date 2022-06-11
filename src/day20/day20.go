package main

import (
	aocreader "AoC2021/src/AoCReader"
	"math"
)

type byte9 [9]bool

type matrix [][]bool

type view struct {
	v  [3][3]bool
	mr int
	mc int
}

func (m matrix) countLit() int {
	r := 0
	for _, row := range m {
		println()
		for _, c := range row {
			if c {
				r++
				print("#")
			} else {
				print(".")
			}
		}
	}

	return r
}

func (v view) hasValue() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if v.v[i][j] {
				return true
			}
		}
	}
	return false
}

func getbyte9(v view) (r byte9) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r[i*3+j] = v.v[i][j]
		}
	}
	return
}

func (b byte9) toDecimal() int {
	r := 0
	for i := 0; i < 9; i++ {
		if b[i] {
			r += int(math.Round(math.Pow(2, float64(8-i))))
		}
	}
	return r
}

func NewView(row, col int, value bool, source matrix) *view {
	v := new(view)
	v.mr = row
	v.mc = col

	v.v[1][1] = value
	if row > 0 && col > 0 {
		v.v[0][0] = source[row-1][col-1]
	}
	if row > 0 {
		v.v[0][1] = source[row-1][col]
	}
	if row > 0 && col < len(source[0])-1 {
		v.v[0][2] = source[row-1][col+1]
	}
	if col > 0 {
		v.v[1][0] = source[row][col-1]
	}
	if col < len(source[0])-1 {
		v.v[1][2] = source[row][col]
	}
	if row < len(source)-1 && col > 0 {
		v.v[2][0] = source[row+1][col-1]
	}
	if row < len(source)-1 {
		v.v[2][1] = source[row+1][col]
	}
	if row < len(source)-1 && col < len(source[0])-1 {
		v.v[2][2] = source[row+1][col+1]
	}

	return v
}

func main() {
	in := aocreader.OpenFile("ex20.txt")

	in.Scan()
	imageEA := in.Text()
	grid := matrix{}
	for i := 0; i < 5; i++ {
		row := make([]bool, 15)
		grid = append(grid, row)
	}
	for in.Scan() {
		line := in.Text()
		if line != "" {
			row := make([]bool, 5)
			for _, r := range line {
				if r == '.' {
					row = append(row, false)
				} else {
					row = append(row, true)
				}
			}
			rowend := make([]bool, 5)
			row = append(row, rowend...)
			grid = append(grid, row)
		}
	}
	for i := 0; i < 5; i++ {
		row := make([]bool, 15)
		grid = append(grid, row)
	}
	println(grid.countLit())

	banaan := view{
		v:  [3][3]bool{{false, false, false}, {true, false, false}, {false, true, false}},
		mr: 0,
		mc: 0,
	}
	println(getbyte9(banaan).toDecimal())

	image := []view{}
	//find all points and neighbours
	for rk, row := range grid {
		for ck, col := range row {
			v := NewView(rk, ck, col, grid)
			if v.hasValue() {
				image = append(image, *v)
			}
		}
	}

	for _, view := range image {
		x := getbyte9(view).toDecimal()
		if imageEA[x:x+1] == "#" {
			grid[view.mr][view.mc] = true
		}
		if imageEA[x:x+1] == "." {
			grid[view.mr][view.mc] = false
		}

	}
	println(grid.countLit())

	image = []view{}
	//find all points and neighbours
	for rk, row := range grid {
		for ck, col := range row {
			v := NewView(rk, ck, col, grid)
			if v.hasValue() {
				image = append(image, *v)
			}
		}
	}

	for _, view := range image {
		x := getbyte9(view).toDecimal()
		if imageEA[x:x+1] == "#" {
			grid[view.mr][view.mc] = true
		}
	}
	println(grid.countLit())

	print(imageEA, in)
}
