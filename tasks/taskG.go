package tasks

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func MainG() {

	// i = ++; d = --; s = *=; o = add to array
	// deadFish("iiisdoso")

	// ( = не повторяется; ) = повторяется.
	// duplicateEncode("recede")

	// в квадрат числа от 0 до первого X, потом посчитать сколько цифр от второго Y во всех полученных квадратах.
	// nbDig(5, 3) // передаём число X и цифру Y

	// factorial(10)

	// приходит рецепт и наши текущие продукты. нужно посчитать сколько тортов можно испечь.
	// cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200})

	// сделать из прямоугольника со сторонами X Y = квадраты со сторонами M
	fmt.Println(squaresInRect(4, 4))
}

func squaresInRect(lng int, wdth int) []int {
	var res []int
	if lng == wdth {
		return res
	}
	for lng != wdth {
		if lng < wdth {
			wdth, lng = lng, wdth
		}
		res = append(res, wdth)
		lng = lng - wdth
	}
	res = append(res, wdth)
	return res
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
	// слишком долго.
	// var allNumber string
	// var result int
	// for i := 0; i <= n; i++ {
	// 	allNumber += strconv.Itoa(i * i)
	// }
	// for _, items := range allNumber {
	// 	if string(items) == strconv.Itoa(d) {
	// 		result++
	// 	}
	// }
	// return result

	// быстрее
	var buffer bytes.Buffer
	var result int
	for i := 0; i <= n; i++ {
		buffer.WriteString(strconv.Itoa(i * i))
	}
	for _, items := range buffer.String() {
		if string(items) == strconv.Itoa(d) {
			result++
		}
	}
	return result

	// без записи всего в одну строку
	// var result int
	// for i := 0; i <= n; i++ {
	// 	allNumber := strconv.Itoa(i * i)
	// 	for _, items := range allNumber {
	// 		if string(items) == strconv.Itoa(d) {
	// 			result++
	// 		}
	// 	}
	// }
	// return result
}

func factorial(n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func cakes(recipe, available map[string]int) int {
	var result int = -1

	for keys := range recipe {
		var temp = available[keys] / recipe[keys]
		if temp < result || result < 0 {
			result = temp
		}
	}

	return result
}
