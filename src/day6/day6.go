package main

func main() {
	days := 256

	visjes := []int{1, 2, 1, 3, 2, 1, 1, 5, 1, 4, 1, 2, 1, 4, 3, 3, 5, 1, 1, 3, 5, 3, 4, 5, 5, 4, 3, 1, 1, 4, 3, 1, 5, 2, 5, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 1, 4, 4, 4, 1, 4, 4, 1, 4, 2, 1, 1, 1, 1, 3, 5, 4, 3, 3, 5, 4, 1, 3, 1, 1, 2, 1, 1, 1, 4, 1, 2, 5, 2, 3, 1, 1, 1, 2, 1, 5, 1, 1, 1, 4, 4, 4, 1, 5, 1, 2, 3, 2, 2, 2, 1, 1, 4, 3, 1, 4, 4, 2, 1, 1, 5, 1, 1, 1, 3, 1, 2, 1, 1, 1, 1, 4, 5, 5, 2, 3, 4, 2, 1, 1, 1, 2, 1, 1, 5, 5, 3, 5, 4, 3, 1, 3, 1, 1, 5, 1, 1, 4, 2, 1, 3, 1, 1, 4, 3, 1, 5, 1, 1, 3, 4, 2, 2, 1, 1, 2, 1, 1, 2, 1, 3, 2, 3, 1, 4, 5, 1, 1, 4, 3, 3, 1, 1, 2, 2, 1, 5, 2, 1, 3, 4, 5, 4, 5, 5, 4, 3, 1, 5, 1, 1, 1, 4, 4, 3, 2, 5, 2, 1, 4, 3, 5, 1, 3, 5, 1, 3, 3, 1, 1, 1, 2, 5, 3, 1, 1, 3, 1, 1, 1, 2, 1, 5, 1, 5, 1, 3, 1, 1, 5, 4, 3, 3, 2, 2, 1, 1, 3, 4, 1, 1, 1, 1, 4, 1, 3, 1, 5, 1, 1, 3, 1, 1, 1, 1, 2, 2, 4, 4, 4, 1, 2, 5, 5, 2, 2, 4, 1, 1, 4, 2, 1, 1, 5, 1, 5, 3, 5, 4, 5, 3, 1, 1, 1, 2, 3, 1, 2, 1, 1}
	vissen := newVissen(visjes)

	println(fokVissen(days, vissen))

}

func fokVissen(days int, vissen [9]int) int {
	for i := 0; i < days; i++ {
		visnieuw := vissen[0]
		vissen[0] = vissen[1]
		vissen[1] = vissen[2]
		vissen[2] = vissen[3]
		vissen[3] = vissen[4]
		vissen[4] = vissen[5]
		vissen[5] = vissen[6]
		vissen[6] = vissen[7] + visnieuw
		vissen[7] = vissen[8]
		vissen[8] = visnieuw

	}
	return sumVissen(vissen)
}

func newVissen(visjes []int) [9]int {
	var vissen [9]int
	for _, value := range visjes {
		switch value {
		case 0:
			vissen[0]++
		case 1:
			vissen[1]++
		case 2:
			vissen[2]++
		case 3:
			vissen[3]++
		case 4:
			vissen[4]++
		case 5:
			vissen[5]++
		case 6:
			vissen[6]++
		case 7:
			vissen[7]++
		case 8:
			vissen[8]++
		}
	}
	return vissen
}

func sumVissen(in [9]int) int {
	result := 0
	for _, value := range in {
		result += value
	}
	return result
}
