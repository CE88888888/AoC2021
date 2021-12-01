package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	scanner := openFile()

	increasedCountSingle, equalOrDecreasedCountSingle := 0, 0
	increasedCountSet, equalOrDecreasedSet := 0, 0
	old := 0
	a, b, c := 0, 0, 0

	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			break
		}

		// Opdracht 1
		increasedCountSingle, equalOrDecreasedCountSingle = simpleCompare(current, old, increasedCountSingle, equalOrDecreasedCountSingle)
		old = current

		//Opdracht 2
		if c > 0 {
			if (a + b + c) < (b + c + current) {
				increasedCountSet++
			} else {
				equalOrDecreasedSet++
			}

			a, b, c = b, c, current

		}

		// Eerste 3 cycles
		if a == 0 {
			a = current
		} else if b == 0 {
			b = current
		} else if c == 0 {
			c = current
		}
	}

	fmt.Printf("Single Increase/Decrease: %v/%v\n", increasedCountSingle, equalOrDecreasedCountSingle)
	fmt.Printf("Set of 3 Increase/Decrease: %v/%v\n", increasedCountSet, equalOrDecreasedSet)
}

func simpleCompare(current int, old int, increasedCountSingle int, equalordecreasedCountSingle int) (int, int) {
	if current > old && old > 0 {
		increasedCountSingle++
	} else {
		equalordecreasedCountSingle++
	}
	return increasedCountSingle, equalordecreasedCountSingle
}

func openFile() *bufio.Scanner {
	file, err := os.Open("INday1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}
