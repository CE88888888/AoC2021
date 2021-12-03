package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//	maxbin := 4095
	gammaRate, epsilonRate := 0, 0
	var evenArray [12]int
	var oddArray [12]int
	var gammaArray [12]int

	scanner := openFile()
	for scanner.Scan() {
		input := scanner.Text()
		current, err := strconv.ParseInt(input, 2, 16)
		if err != nil {
			log.Fatal()
		}

		for i := 0; i < 12; i++ {

			if current&1 == 1 {

				oddArray[i] += 1
			} else {

				evenArray[i] += 1
			}
			current = current >> 1

		}
	}

	for i := 0; i < 12; i++ {

		if oddArray[i] > evenArray[i] {
			gammaArray[i] = 1
		}

	}

	for key, value := range gammaArray {
		if key == 0 {
			gammaRate = 1 * value
			epsilonRate = 1 * (value ^ 1)
		} else {
			gammaRate = gammaRate + (IntPow(2, key) * value)
			epsilonRate = epsilonRate + (IntPow(2, key) * (value ^ 1))

		}

	}
	fmt.Printf("%08b\n", gammaRate)
	fmt.Printf("%08b\n", epsilonRate)
	println(gammaRate * epsilonRate)

}

func openFile() (scanner *bufio.Scanner) {
	if file, err := os.Open("day3.txt"); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		log.Fatal(err)
	}
	return
}
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
