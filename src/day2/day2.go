package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	xpos  int
	depth int
	aim   int
}

func (s *Submarine) MoveSimple(direction string, amount int) {
	switch direction {
	case "up":
		s.depth = s.depth - amount
	case "down":
		s.depth = s.depth + amount
	case "forward":
		s.xpos = s.xpos + amount
	}
}

func (s *Submarine) MoveAdvanced(direction string, amount int) {
	switch direction {
	case "up":
		s.aim = s.aim - amount
	case "down":
		s.aim = s.aim + amount
	case "forward":
		s.xpos = s.xpos + amount
		if s.aim != 0 {
			s.depth = s.depth + s.aim*amount
		}
	}
}

func (s Submarine) String() string {
	str := []string{"Submarine: xpos = ", strconv.FormatInt(int64(s.xpos), 10),
		" | depth = ", strconv.FormatInt(int64(s.depth), 10),
		" | aim = ", strconv.FormatInt(int64(s.aim), 10),
		" | multiplied = ", strconv.FormatInt(int64(s.depth*s.xpos), 10),
	}
	return strings.Join(str, "")
}

/*
Day 2, exercise 1: move a submarine with simple controls.
Day 2, exercise 2: move a submarine with advanced controls
*/

func main() {

	submarineSimple := Submarine{xpos: 0, depth: 0, aim: 0}
	submarineAdvanced := Submarine{xpos: 0, depth: 0, aim: 0}

	scanner := openFile()

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		direction := input[0]
		amount, err := strconv.Atoi(input[1])

		if err != nil {
			log.Fatal(err)
		}

		submarineSimple.MoveSimple(direction, amount)
		submarineAdvanced.MoveAdvanced(direction, amount)

	}
	fmt.Println(submarineSimple)
	fmt.Println(submarineAdvanced)

}
func openFile() *bufio.Scanner {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}
