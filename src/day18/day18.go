package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Magresult struct {
	left   string
	right  string
	sum    string
	result string
	mag    int
}

func main() {

	scanner := openFile("day18.txt")
	//scanner := openFile("ex18.txt")
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	//Exercise A
	result := input[0]
	for i := 1; i < len(input); i++ {
		result = addSnails(result, input[i])
		result = checkSplit(explodeSnail(result), result)
	}
	println(getMagnitude(result))

	//Exercise B
	results := []Magresult{}
	mags := []int{}
	//Actual pair and sum of pair is not required for the exercise
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i != j {
				sum := addSnails(input[i], input[j])
				mr := Magresult{
					left:   input[i],
					right:  input[j],
					sum:    addSnails(input[i], input[j]),
					result: checkSplit(explodeSnail(sum), sum),
				}
				x, _ := strconv.Atoi(getMagnitude(mr.result))
				mr.mag = x
				results = append(results, mr)
				mags = append(mags, x)
			}
		}
	}
	sort.Ints(mags)
	println(mags[len(mags)-1])

}

func getMagnitude(s string) string {
	r := regexp.MustCompile(`\d{1,4},\d{1,4}`)
	for matches := r.FindAllString(s, -1); matches != nil; matches = r.FindAllString(s, -1) {
		for _, pair := range matches {
			xc := regexp.MustCompile(`,`)
			index := xc.FindStringIndex(pair)
			x, _ := strconv.Atoi(pair[:index[0]])
			y, _ := strconv.Atoi(pair[index[1]:])
			m := 3*x + 2*y
			s = strings.Replace(s, "["+pair+"]", strconv.Itoa(m), 1)
		}
	}
	return s
}

func checkSplit(s1 string, source string) string {
	if s1 != source {
		source = s1
		s1 = checkExplode(explodeSnail(source), source)
		s1 = checkSplit(splitPair(s1), s1)
	}
	return s1
}

func checkExplode(s1 string, source string) string {
	if s1 != source {
		source = s1
		s1 = checkExplode(explodeSnail(source), source)
	}
	return s1
}

func addSnails(s, s2 string) string {
	return "[" + s + "," + s2 + "]"
}

func splitPair(s string) string {
	tosplit := 0
	index := 0
	substring := ""
	for i := 0; i < len(s)-1; i++ {
		x, err := strconv.Atoi(s[i : i+2])
		if err == nil && x > 9 {
			tosplit = x
			index = i
			break
		}
	}
	if tosplit > 0 {
		substring = s[index : index+2]
		x1 := int(math.Floor(float64(tosplit) / 2))
		x2 := int(math.Ceil(float64(tosplit) / 2))
		newpair := "[" + strconv.Itoa(x1) + "," + strconv.Itoa(x2) + "]"
		s = strings.Replace(s, substring, newpair, 1)
		return s
	}
	return s
}

func explodeSnail(s string) string {

	x := 0
	left := -1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			x++
			if x == 5 {
				replacement := ""
				pairstart, pairend := getPair(s[i:], i)
				pair := s[pairstart:pairend]
				xleft, xright := getNumbersFromPair(pair)
				if left != -1 {
					x1, indexx1 := checkLeftNumber(s, left)
					x1 = x1 + xleft
					replacement = strconv.Itoa(x1) + s[left+1:i]
					pair = s[indexx1:i] + pair
				}
				replacement = replacement + "0"
				right, xr := checkRightNumber(s[pairend:], pairend)
				if right != -1 {
					pair = pair + s[pairend:right+len(strconv.Itoa(xr))]
					xr = xr + xright
					xasstring := strconv.Itoa(xr)
					replacement = replacement + s[pairend:right] + xasstring
				}
				if !strings.Contains(s, pair) {
					panic("string construction gone awry")
				}
				s = strings.Replace(s, pair, replacement, 1)
				return s
			}
		case ']':
			x--
		case ',':
		default:
			left = i
		}
	}
	return s
}

func checkLeftNumber(s string, left int) (number int, index int) {
	if !(s[left-1] == ',' || s[left-1] == '[' || s[left-1] == ']') {
		if !(s[left-2] == ',' || s[left-2] == '[' || s[left-2] == ']') {
			number, _ := strconv.Atoi(s[left-2 : left+1])
			return number, left - 2
		}
		number, _ := strconv.Atoi(s[left-1 : left+1])
		return number, left - 1
	} else {
		return int(s[left] - '0'), left
	}

}

func getNumbersFromPair(pair string) (x1, x2 int) {
	commai := 0
	for i := 0; commai == 0; i++ {
		if pair[i] == ',' {
			commai = i
		}
	}
	left := pair[1:commai]
	right := pair[commai+1 : len(pair)-1]
	x1, _ = strconv.Atoi(left)
	x2, _ = strconv.Atoi(right)
	return
}

func getPair(s string, offset int) (i1, i2 int) {
	for i := 0; i < len(s); i++ {
		if s[i] == '[' && i1 == 0 {
			i1 = i + offset
		}
		if s[i] == ']' {
			i2 = i + offset + 1
			return
		}
	}
	return 0, 0
}

func checkRightNumber(s string, offset int) (index int, number int) {
	index = -1
	for i := 0; i < len(s); i++ {
		r := s[i]
		if !(r == '[' || r == ']' || r == ',') {
			index = i
			break
		}
	}
	if index > 0 {
		for j := index; j < len(s); j++ {
			if s[j] == ']' || s[j] == ',' {
				number, _ = strconv.Atoi(s[index:j])
				break
			}
		}
	}
	if index == -1 {
		return -1, 0
	}
	return index + offset, number
}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		println("Err opening file")
	}
	return
}
