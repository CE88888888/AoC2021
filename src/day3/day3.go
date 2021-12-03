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
	nobits := 12

	var evenArray [12]int
	var oddArray [12]int
	var gammaArray [12]int
	var fullSlice []int64
	var oxySlice []int64
	var co2Slice []int64

	scanner := openFile()
	for scanner.Scan() {
		input := scanner.Text()
		current, err := strconv.ParseInt(input, 2, 16)
		if err != nil {
			log.Fatal()
		}

		fullSlice = append(fullSlice, current)
		for i := 0; i < nobits; i++ {

			if current&1 == 1 {

				oddArray[i] += 1
			} else {

				evenArray[i] += 1
			}
			current = current >> 1

		}
	}

	for i := 0; i < nobits; i++ {

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

	oxySlice, co2Slice = filterSplitByBits(fullSlice, oxySlice, co2Slice, nobits-1)

	for i := nobits - 2; len(oxySlice) > 1; i-- {
		println(i)
		oxySlice = filterByBits(oxySlice, i, true)
	}
	for i := nobits - 2; len(co2Slice) > 1; i-- {
		println(i)
		co2Slice = filterByBits(co2Slice, i, false)
	}

	fmt.Printf("%08b\n", gammaRate)
	fmt.Printf("%08b\n", epsilonRate)
	println(gammaRate * epsilonRate)

	println(oxySlice[0] * co2Slice[0])
}

func filterSplitByBits(fullSlice []int64, oxySlice []int64, co2Slice []int64, bitNumber int) ([]int64, []int64) {
	for key, value := range fullSlice {
		power := IntPow(2, bitNumber)
		println(key)
		if (value & int64(power)) == int64(power) {
			oxySlice = append(oxySlice, value)
		} else {
			co2Slice = append(co2Slice, value)
		}
	}
	if len(oxySlice) >= len(co2Slice) {
		return oxySlice, co2Slice
	} else {
		return co2Slice, oxySlice
	}

}
func filterByBits(inSlice []int64, bitNumber int, most bool) (outSlice []int64) {
	power := IntPow(2, bitNumber)
	one := outSlice
	zero := outSlice
	for _, value := range inSlice {
		if (value & int64(power)) == int64(power) {
			one = append(one, value)
		} else {
			zero = append(zero, value)
		}
	}
	if most {
		if len(one) >= len(zero) {
			outSlice = one
		} else {
			outSlice = zero
		}
	} else {
		if len(one) >= len(zero) {
			outSlice = zero
		} else {
			outSlice = one
		}
	}

	return
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
