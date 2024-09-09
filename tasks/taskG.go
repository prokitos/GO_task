package tasks

import (
	"strconv"
	"strings"
)

func MainG() {

	// i = ++; d = --; s = *=; o = add to array
	// deadFish("iiisdoso")

	// ( = не повторяется; ) = повторяется.
	// duplicateEncode("recede")

	// в квадрат числа от 0 до первого X, потом посчитать сколько цифр от второго Y во всех полученных квадратах.
	nbDig(20, 3) // передаём число X и цифру Y
}

func deadFish(input string) []int {

	result := []int{}
	// result := make([]int, 0)
	var curVal int = 0

	for _, items := range input {
		switch items {
		case 'i':
			curVal++
		case 'd':
			curVal--
		case 's':
			curVal = curVal * curVal
		case 'o':
			result = append(result, curVal)
		}
	}

	return result
}

func duplicateEncode(word string) string {

	var result string
	var seth = make(map[string]int, 0)
	// можно сделать в начале новую строку сразу в фулл ловеркейс, и проходить по ней.

	for _, items := range word {
		seth[strings.ToLower(string(items))]++
	}

	for _, items := range word {
		if seth[strings.ToLower(string(items))] == 1 {
			result += "("
		} else {
			result += ")"
		}
	}

	return result
}

func nbDig(n int, d int) int {
	// var allNumber string
	// var result int
	// for i := 0; i < n; i++ {
	// 	allNumber += strconv.Itoa(i * i)
	// }
	// for _, items := range allNumber {
	// 	if string(items) == strconv.Itoa(d) {
	// 		result++
	// 	}
	// }
	// return result

	var result int

	for i := 0; i <= n; i++ {
		allNumber := strconv.Itoa(i * i)
		for _, items := range allNumber {
			if string(items) == strconv.Itoa(d) {
				result++
			}
		}

	}

	return result
}
