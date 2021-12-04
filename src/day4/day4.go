package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoval struct {
	number int
	hit    bool
}

type board struct {
	bord  [][]bingoval
	bingo bool
}

func main() {
	var allBoards []board
	var winOrder []int
	squidwin := true
	winningBoardKey, winningdigit := 0, 0

	scanner := openFile()
	scanner.Scan()
	draw := strings.Split(scanner.Text(), ",")

	// Setup boards
	scanner.Scan() //skip first empty line after draw numbers
	for scanner.Scan() {
		var bingobord board
		for i := 0; len(scanner.Text()) > 0; i++ {
			line := strings.Fields(scanner.Text())
			var row []bingoval

			for _, value := range line {
				var a bingoval
				a.number, _ = strconv.Atoi(value)
				row = append(row, a)
			}
			bingobord.bord = append(bingobord.bord, row)
			scanner.Scan()
		}
		allBoards = append(allBoards, bingobord)
	}
Draw: // Start Drawing, Stop drawing at first or last BINGO, depending on squidwin parameter
	for _, value := range draw {
		d, err := strconv.Atoi(value)
		if err != nil {
			break
		}
		for _, value := range allBoards {
			for _, row := range value.bord {
				for fieldId, field := range row {
					if field.number == d {
						row[fieldId].hit = true
					}
				}
			}
		}
		//hits registered
		for key, value := range allBoards {
			bingo := false
			//row check
			for _, row := range value.bord {
				bingoR := true
				for _, field := range row {
					if !field.hit {
						bingoR = false
					}

				}
				if bingoR {
					winningBoardKey = key
					winningdigit = d
					bingo = true
				}
			}
			//column check
			if !bingo {
				for i := 0; i < len(value.bord); i++ {
					bingoC := true
					for j := 0; j < len(value.bord[0]); j++ {
						if !value.bord[j][i].hit {
							bingoC = false
						}
					}
					if bingoC {
						winningBoardKey = key
						winningdigit = d
						bingo = true
					}
				}
			}
			if bingo {
				if !squidwin {
					break Draw
				} else {
					if !allBoards[key].bingo {
						allBoards[key].bingo = true
						winOrder = append(winOrder, key)
						if len(winOrder) == len(allBoards) {
							break Draw
						}
					}
				}
			}
		}
	}
	fmt.Printf("Winning Board: %v | Winning Digit %v\n", winningBoardKey, winningdigit)
	println(calculateResult(allBoards[winningBoardKey]) * winningdigit)
}

func calculateResult(b board) (total int) {
	for _, row := range b.bord {
		for _, field := range row {
			if !field.hit {
				total += field.number
			}
		}
	}
	return
}

func openFile() (scanner *bufio.Scanner) {
	if file, err := os.Open("day4.txt"); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		log.Fatal(err)
	}
	return
}
