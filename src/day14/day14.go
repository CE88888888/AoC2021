package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type elements map[string]int

func (e *elements) CountStringArray(poly []string) {
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

	//get mapping rules
	rules := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		rules[line[0:2]] = string(line[6])
	}

	paircount := newPairCount(rules, poly)
	eleB := getElements(rules)

	//Exercise A
	eleA := getElements(rules)
	for i := 0; i < 10; i++ {
		poly = growOnce(poly, rules)
		print(i)
	}
	eleA.CountStringArray(poly)
	min, max := getLowMax(eleA)
	fmt.Println(max - min)

	//Exercise B
	for i := 0; i < 40; i++ {
		paircount = growAPair(paircount, rules)
	}
	eleB = translatePairToSingle(paircount, eleB, rules)
	minB, maxB := getLowMax(eleB)
	fmt.Println(((maxB - minB) / 2) + 1)

}

func newPairCount(rules map[string]string, poly []string) map[string]int {
	paircount := map[string]int{}
	for key := range rules {
		paircount[key] = 0
	}

	for i := 0; i+1 < len(poly); i++ {
		pair := poly[i] + poly[i+1]
		paircount[pair]++
	}
	return paircount
}

func translatePairToSingle(paircount map[string]int, eleB elements, rules map[string]string) elements {
	eleR := eleB
	for key, value := range paircount {
		eleR[string(key[0])] += value
		eleR[string(key[1])] += value
	}
	return eleR
}

func growAPair(pc map[string]int, rules map[string]string) map[string]int {
	newpc := make(map[string]int, len(pc))
	for k, v := range pc {
		newpc[k] = v
	}

	for key, value := range pc {
		if value > 0 {
			np1 := string(key[0]) + rules[key]
			np2 := rules[key] + string(key[1])

			newpc[np1] += value
			newpc[np2] += value
			newpc[key] -= value
		}
	}
	return newpc
}

func getLowMax(eles elements) (min int, max int) {
	keyMin, keyMax := "", ""
	for key, value := range eles {
		keyMin = key
		min = value
		break
	}
	for key, value := range eles {
		if value > max {
			keyMax = key
			max = value
		}
		if value < min {
			keyMin = key
			min = value
		}
	}
	fmt.Println("Min:", keyMin, " Max:", keyMax)
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
