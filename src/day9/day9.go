package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

type Bassin struct {
	coord coord
	size  int
}
type coord struct {
	row int
	col int
}

var maxRow, maxColumn int

func main() {
	scanner := openFile("day9.txt")
	//scanner := openFile("ex9.txt")
	matrix := initMatrix(scanner)

	lowestpoints := []int{}
	bassins := []Bassin{}
	risksum := 0
	bassinSizes := []int{}

	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxColumn; c++ {
			p := pickLowest(r, c, matrix)
			if p < 9 {
				lowestpoints = append(lowestpoints, p)
				risksum = risksum + p + 1
				c := coord{
					row: r,
					col: c,
				}
				bassins = append(bassins, Bassin{
					coord: c,
					size:  0,
				})
			}
		}
	}

	for b, value := range bassins {
		checkedCoords := []coord{}
		determineBassinSize(value.coord, &bassins[b], matrix, &checkedCoords)
		bassinSizes = append(bassinSizes, bassins[b].size)
	}

	sort.Ints(bassinSizes)
	a, b, c := bassinSizes[len(bassinSizes)-1], bassinSizes[len(bassinSizes)-2], bassinSizes[len(bassinSizes)-3]
	println("Exercise A:", risksum)
	println("Exercise B:", a*b*c)

}

func determineBassinSize(co coord, b *Bassin, matrix [][]int, cc *[]coord) {

	r, c := co.row, co.col
	depth := matrix[r][c]

	//ignore Nines
	if depth > 8 {
		return
	}

	//check if already visited
	for _, value := range *cc {
		if value == co {
			return
		}
	}

	//Add 1 to size, add coord to list of checked coords
	b.size++
	*cc = append(*cc, co)

	if r > 0 {
		determineBassinSize(coord{
			row: r - 1,
			col: c,
		}, b, matrix, cc)
	}
	if r < maxRow-1 {
		determineBassinSize(coord{
			row: r + 1,
			col: c,
		}, b, matrix, cc)
	}
	if c > 0 {
		determineBassinSize(coord{
			row: r,
			col: c - 1,
		}, b, matrix, cc)
	}
	if c < maxColumn-1 {
		determineBassinSize(coord{
			row: r,
			col: c + 1,
		}, b, matrix, cc)
	}
}

func pickLowest(r, c int, matrix [][]int) int {
	sample := []int{}
	if r > 0 {
		sample = append(sample, matrix[r-1][c])
	}
	if r < maxRow-1 {
		sample = append(sample, matrix[r+1][c])
	}
	if c > 0 {
		sample = append(sample, matrix[r][c-1])
	}
	if c < maxColumn-1 {
		sample = append(sample, matrix[r][c+1])
	}

	sort.Ints(sample)
	if matrix[r][c] < sample[0] {
		return matrix[r][c]
	} else {
		return 9
	}
}

func initMatrix(scanner *bufio.Scanner) [][]int {
	var matrix [][]int
	//maxRow, maxColumn := 0, 0

	for scanner.Scan() {
		maxRow++
		if maxColumn == 0 {
			maxColumn = len(scanner.Text())
		}
		textrow := scanner.Text()
		row := []int{}
		for _, c := range textrow {
			row = append(row, int(c-'0'))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		log.Fatal(err)
	}
	return
}
