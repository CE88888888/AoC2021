package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type DisplayDigits struct {
	zero  string
	one   string
	two   string
	three string
	four  string
	five  string
	six   string
	seven string
	eigth string
	nine  string
}

func main() {
	scanner := *openFile()
	count := 0
	grandtotal := 0
	for scanner.Scan() {

		in := strings.Split(scanner.Text(), "|")

		patterns := strings.Fields(in[0])
		output := strings.Fields(in[1])

		display := new(DisplayDigits)
		var len5 []string
		var len6 []string

		//Determine 1,4,7,8 - fill length 5 & 6 pattern-slices
		matchPatternToDigit(patterns, display, &len5, &len6)
		//Strings with length 5 can be either 2,3 or 5
		determineLen5Values(len5, display)
		//Strings with length 6 can be either 0,6 or 9
		determineLen6Values(len6, display)

		// Based on input patterns determine the values for a-g
		decodeMap := decipherSegmentValues(display)

		//Start counting for Exercise A and while looping also translate the values based on the decodeMap
		var translated []string
		count, translated = newFunction(output, count, decodeMap, translated)
		//With the translated values determine the Digit output
		outputValues := determineDigits(translated)
		grandtotal = grandtotal + sliceToInt(outputValues)
	}
	fmt.Println(count)
	println(grandtotal)
}

func matchPatternToDigit(patterns []string, display *DisplayDigits, len5 *[]string, len6 *[]string) {
	for _, value := range patterns {
		switch len(value) {
		case 2:
			display.one = value
		case 3:
			display.seven = value
		case 4:
			display.four = value
		case 5:
			*len5 = append(*len5, value)
		case 6:
			*len6 = append(*len6, value)
		case 7:
			display.eigth = value
		}
	}
}

func determineDigits(translated []string) []int {
	mapToDigits := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	outputValues := []int{}

	for _, value := range translated {
		s := SortStringByCharacter(value)
		outputValues = append(outputValues, mapToDigits[s])
	}
	return outputValues
}

func newFunction(output []string, count int, decodeMap map[string]string, translated []string) (int, []string) {
	for _, value := range output {
		if (len(value) == 2) || (len(value) == 3) || (len(value) == 4) || (len(value) == 7) {
			count++
		}
		s := ""
		for i := 0; i < len(value); i++ {
			a := decodeMap[string(value[i])]
			s = s + a
		}
		s = SortStringByCharacter(s)
		translated = append(translated, s)
	}
	return count, translated
}

func decipherSegmentValues(d *DisplayDigits) map[string]string {
	segments := map[string]string{
		"a": "",
		"b": "",
		"c": "",
		"d": "",
		"e": "",
		"f": "",
		"g": "",
	}

	for _, value := range d.seven {
		if !strings.ContainsRune(d.one, value) {
			segments["a"] = string(value)
		}
	}

	for _, value := range d.eigth {
		if !strings.ContainsRune(d.six, value) {
			segments["c"] = string(value)

		}
	}

	for _, value := range d.one {
		if string(value) != segments["c"] {
			segments["f"] = string(value)
		}
	}

	for _, value := range d.eigth {
		if !strings.ContainsRune(d.nine, value) {
			segments["e"] = string(value)
		}
	}

	for _, value := range d.eigth {
		if !strings.ContainsRune(d.zero, value) {
			segments["d"] = string(value)
		}
	}

	for _, value := range d.eigth {
		s := d.four + d.seven + segments["e"]
		if !strings.ContainsRune(s, value) {
			segments["g"] = string(value)
		}
	}

	for _, value := range d.four {
		s := segments["c"] + segments["d"] + segments["f"]
		if !strings.ContainsRune(s, value) {
			segments["b"] = string(value)
		}
	}

	decodeMap := SwitchKVMap(segments)
	return decodeMap
}

func SwitchKVMap(m map[string]string) (switchedMap map[string]string) {
	switchedMap = make(map[string]string)
	for i, s := range m {
		switchedMap[SortStringByCharacter(s)] = i
	}
	return
}

func determineLen6Values(len6 []string, display *DisplayDigits) {
	for _, pattern := range len6 {
		if checkforSix(pattern, display) {
			display.six = pattern
		} else {
			if checkforNine(pattern, display) {
				display.nine = pattern
			} else {
				display.zero = pattern
			}
		}
	}
}

func determineLen5Values(len5 []string, display *DisplayDigits) {
	for _, pattern := range len5 {
		if checkforTwo(pattern, display) {
			display.two = pattern
		} else {
			if checkforThree(pattern, display) {
				display.three = pattern
			} else {
				display.five = pattern
			}
		}
	}
}

func checkforTwo(pattern string, display *DisplayDigits) (r bool) {
	a := 0
	for _, c := range display.four {
		if strings.ContainsRune(pattern, c) {
			a++
		}
	}
	return (a == 2)
}
func checkforThree(pattern string, display *DisplayDigits) (r bool) {
	a := 0
	for _, c := range display.seven {
		if strings.ContainsRune(pattern, c) {
			a++
		}
	}
	return (a == 3)
}
func checkforSix(pattern string, display *DisplayDigits) (r bool) {
	a := 0
	for _, c := range display.one {
		if strings.ContainsRune(pattern, c) {
			a++
		}
	}
	return (a == 1)
}
func checkforNine(pattern string, display *DisplayDigits) (r bool) {
	a := 0
	for _, c := range display.four {
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
func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}
