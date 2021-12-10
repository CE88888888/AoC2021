package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func main() {
	//scanner := openFile("ex10.txt")
	scanner := openFile("day10.txt")
	illegals := []rune{}
	incompletes := []string{}

Line:
	for scanner.Scan() {
		textrow := scanner.Text()
		buffer := []rune{}

		for _, c := range textrow {
			switch c {
			case '(':
				buffer = append(buffer, c)
			case '[':
				buffer = append(buffer, c)
			case '{':
				buffer = append(buffer, c)
			case '<':
				buffer = append(buffer, c)
			case ')':
				if buffer[len(buffer)-1] == '(' {
					buffer = buffer[:len(buffer)-1]
				} else {
					illegals = append(illegals, c)
					continue Line
				}
			case ']':
				if buffer[len(buffer)-1] == '[' {
					buffer = buffer[:len(buffer)-1]
				} else {
					illegals = append(illegals, c)
					continue Line
				}

			case '}':
				if buffer[len(buffer)-1] == '{' {
					buffer = buffer[:len(buffer)-1]
				} else {
					illegals = append(illegals, c)
					continue Line
				}
			case '>':
				if buffer[len(buffer)-1] == '<' {
					buffer = buffer[:len(buffer)-1]
				} else {
					illegals = append(illegals, c)
					continue Line
				}

			}
		}
		if len(buffer) > 0 {
			incompletes = append(incompletes, string(buffer))
		}

	}

	println("Exercise A: ", calculateScoreIL(illegals))

	linescores := calculateScoreAC(incompletes)
	sort.Ints(linescores)
	println("Exercise B: ", linescores[len(linescores)/2])

}

func calculateScoreAC(incompletes []string) []int {
	scoreMapAC := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	linescores := []int{}
	for _, line := range incompletes {
		linescore := 0
		for i := len(line) - 1; i > -1; i-- {
			chars := []rune(line)
			linescore = (linescore * 5) + scoreMapAC[chars[i]]
		}
		linescores = append(linescores, linescore)
	}
	return linescores
}

func calculateScoreIL(illegals []rune) int {
	scoreMapIllegal := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	illegalScore := 0
	for _, c := range illegals {
		illegalScore += scoreMapIllegal[c]
	}
	return illegalScore
}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		log.Fatal(err)
	}
	return
}
