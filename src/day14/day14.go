package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type elements map[string]int

func (e *elements) Count(poly []string) {
	f := *e
	for _, value := range poly {
		f[value]++
	}
	e = &f
}

func main() {
	//scanner := openFile("ex14.txt")
	scanner := openFile("day14.txt")

	scanner.Scan()
	start := scanner.Text()
	scanner.Scan()
	poly := strings.Split(start, "")

	rules := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		rules[line[0:2]] = string(line[6])
	}

	eles := getElements(rules)

	for i := 0; i < 40; i++ {
		poly = growOnce(poly, rules)
		print(i)
	}
	eles.Count(poly)
	min, max := getLowMax(eles)

	print(start, poly[0])
	fmt.Println(min)
	fmt.Println(max)
}

func getLowMax(eles elements) (min int, max int) {
	for _, value := range eles {
		min = value
		break
	}
	for _, value := range eles {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return
}

func getElements(rules map[string]string) elements {
	ele := map[string]int{}
	for _, v := range rules {
		if _, ok := ele[v]; ok {
			//skip
		} else {
			ele[v] = 0
		}
	}
	return ele
}

func growOnce(poly []string, rules map[string]string) (newPoly []string) {
	for i := 0; i+1 < len(poly); i++ {
		pair := poly[i] + poly[i+1]
		newpair := make([]string, 2)
		newpair[0] = poly[i]
		newpair[1] = rules[pair]

		//newpair[2] = poly[i+1]
		newPoly = append(newPoly, newpair...)
		if i+2 == len(poly) {
			newPoly = append(newPoly, poly[i+1])
		}
	}

	return
}

func openFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		println("Err opening file")
	}
	return
}
