package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := *openFile()
	count := 0
	grandtotal := 0
	for scanner.Scan() {

		in := strings.Split(scanner.Text(), "|")
		patterns := strings.Fields(in[0])
		output := strings.Fields(in[1])

		display := map[int]string{
			0: "zero",
			1: "one",
			2: "two",
			3: "three",
			4: "four",
			5: "five",
			6: "six",
			7: "seven",
			8: "eigth",
			9: "nine",
		}

		var len5 []string
		var len6 []string

		//Exercise B
		//Determine 1,4,7,8 - fill length 5 & 6 pattern-slices
		for _, value := range patterns {
			switch len(value) {
			case 2:
				display[1] = value

			case 3:
				display[7] = value

			case 4:
				display[4] = value

			case 5:
				len5 = append(len5, value)
			case 6:
				len6 = append(len6, value)
			case 7:
				display[8] = value

			}
		}

		determineLen5Values(len5, display)
		determineLen6Values(len6, display)

		//switch map and use it to calculate the results
		switchedDisplay := SwitchKVMap(display)
		var total int
		total, count = calculateResult(output, count, switchedDisplay)

		grandtotal += total

	}
	fmt.Println("Exercise A: ", count)
	fmt.Println("Exercise B: ", grandtotal)
}

func SwitchKVMap(display map[int]string) (switchedMap map[string]int) {
	switchedMap = make(map[string]int)
	for i, s := range display {
		switchedMap[SortStringByCharacter(s)] = i
	}
	return
}

func calculateResult(output []string, count int, switchedDisplay map[string]int) (int, int) {
	total := ""
	for _, value := range output {

		if (len(value) == 2) || (len(value) == 3) || (len(value) == 4) || (len(value) == 7) {
			count++
		}
		// Awkard convert, receiving the numbers as characters from string
		total = total + strconv.Itoa(switchedDisplay[SortStringByCharacter(value)])

	}
	totalInt, _ := strconv.Atoi(total)
	return totalInt, count
}

func determineLen6Values(len6 []string, display map[int]string) {
	for _, pattern := range len6 {
		if checkforSix(pattern, display) {
			display[6] = pattern
		} else {
			if checkforNine(pattern, display) {
				display[9] = pattern
			} else {
				display[0] = pattern
			}
		}
	}
}

func determineLen5Values(len5 []string, display map[int]string) {
	for _, pattern := range len5 {
		if checkforTwo(pattern, display) {
			display[2] = pattern
		} else {
			if checkforThree(pattern, display) {
				display[3] = pattern
			} else {
				display[5] = pattern
			}
		}
	}
}

func checkforTwo(pattern string, display map[int]string) (r bool) {
	a := 0
	for _, c := range display[4] {
		if strings.ContainsRune(pattern, c) {
			a++
		}
	}
	return (a == 2)
}
func checkforThree(pattern string, display map[int]string) (r bool) {
	a := 0
	for _, c := range display[7] {
		if strings.ContainsRune(pattern, c) {
			a++
		}
	}
	return (a == 3)
}
func checkforSix(pattern string, display map[int]string) (r bool) {
	a := 0
	for _, c := range display[1] {
		if strings.ContainsRune(pattern, c) {
			a++
		}
	}
	return (a == 1)
}
func checkforNine(pattern string, display map[int]string) (r bool) {
	a := 0
	for _, c := range display[4] {
		if strings.ContainsRune(pattern, c) {
			a++
		}
	}
	return (a == 4)
}

func openFile() (scanner *bufio.Scanner) {
	if file, err := os.Open("day8.txt"); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		log.Fatal(err)
	}
	return
}
func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
