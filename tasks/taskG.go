package tasks

import (
	"bytes"
	"fmt"
	"math"
	"sort"
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
	// fmt.Println(squaresInRect(4, 4))

	// дано время для нескольких участников. найти разницу между большим и меньшим временем, среднее время, и медиану.
	// fmt.Println(stati("01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"))

	// получить век из года
	// fmt.Println(WhatCentury("7901"))

	// посчитать количество цифр если сложить все числа от 1 до N
	// fmt.Println(AmountOfPages(10))
	// fmt.Println(ReverseAmountOfPages(11))

	// верный ли айпи
	// fmt.Println(Is_valid_ip("1.1.1.1"))

	// найти сколько чисел от X до Y делятся нацело на Z.
	// fmt.Println(DivisibleCount(6, 11, 2))
}

func DivisibleCount(x, y, k uint64) uint64 {

	// сложность 0(N)
	// var res uint64 = 0
	// for i := x; i <= y; i++ {
	// 	if i%k == 0 {
	// 		res++
	// 	}
	// }
	// return res

	// сложность O(1)
	var res uint64 = 0
	res = y / k
	res -= x / k
	if x%k == 0 {
		res++
	}
	return res
}

func Is_valid_ip(ip string) bool {

	allNums := strings.Split(ip, ".")
	if len(allNums) != 4 {
		return false
	}
	for _, items := range allNums {
		curItem, err := strconv.Atoi(items)
		if curItem < 0 || curItem > 255 || err != nil || strconv.Itoa(curItem) != items {
			return false
		}
	}

	return true
}

func ReverseAmountOfPages(summary int) int {

	var summaryStr = strconv.Itoa(summary)
	var res int = summary * len(summaryStr)

	for i := 1; i < len(summaryStr); i++ {
		temp := strings.Repeat("9", i)
		intTemp, _ := strconv.Atoi(temp)
		res = res - intTemp
	}

	return res
}

func AmountOfPages(summary int) int {

	sum := 0
	digit := 1
	count := 0
	a := 10
	for sum < summary {
		count += 1
		if count%a == 0 {
			digit += 1
			a *= 10
		}
		sum += digit
	}

	if sum > summary {
		count -= 1
	}
	return count
}

func WhatCentury(year string) string {

	// проверка что век хотябы первый
	lenIs := len(year) - 2
	if lenIs <= 0 {
		return ""
	}

	// проверка что если век с двумя нулями в конце то будет прошлый, а иначе век + 1
	result, _ := strconv.Atoi(year[0:lenIs])
	zeroCheck := year[lenIs:len(year)]
	zeroNum, _ := strconv.Atoi(zeroCheck)
	if zeroNum != 0 {
		result++
	}

	// окончания для веков
	var suffix string = "th"
	if result < 11 || result > 13 {
		switch result % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}

	return strconv.Itoa(result) + suffix
}

// 01|01|18 Average: 01|38|05 Median: 01|32|34
func stati(strg string) string {

	if strg == "" {
		return ""
	}

	allTeams := strings.Split(strg, ",")
	var allTeamsInt []int
	var resRanges int
	var resAverge int
	var resMedian int
	var maxLen int = len(allTeams)

	for _, items := range allTeams {
		newItem := strings.Split(items, "|")
		var val int = 0
		for iter, curval := range newItem {
			temp, _ := strconv.Atoi(strings.TrimSpace(curval))
			var pow int = int(math.Pow(60, float64(2-iter)))
			val += temp * pow
		}
		//val := hmsToNumber(temp) // вместо цикла внутри цикла
		resAverge += val
		allTeamsInt = append(allTeamsInt, val)
	}

	sort.Ints(allTeamsInt)

	if len(allTeamsInt)%2 == 0 {
		resMedian = (allTeamsInt[maxLen/2] + allTeamsInt[maxLen/2-1]) / 2

	} else {
		resMedian = allTeamsInt[maxLen/2]
	}

	resRanges = allTeamsInt[maxLen-1] - allTeamsInt[0]
	resAverge = resAverge / maxLen

	// response := "Range: " + numberToHms(resRanges) + " Average: " + numberToHms(resAverge) + " Median: " + numberToHms(resMedian)
	response := fmt.Sprintf("Range: %s Average: %s Median: %s", numberToHms(resRanges), numberToHms(resAverge), numberToHms(resMedian))
	return response
}

//	func hmsToNumber(value []string) int {
//		// можно было в основной функции в цикле считать сразу
//		var resTime int = 0
//		temp, _ := strconv.Atoi(strings.TrimSpace(value[2]))
//		resTime += temp
//		temp, _ = strconv.Atoi(strings.TrimSpace(value[1]))
//		resTime += temp * 60
//		temp, _ = strconv.Atoi(strings.TrimSpace(value[0]))
//		resTime += temp * 60 * 60
//		return resTime
//	}
func numberToHms(value int) string {

	return fmt.Sprintf("%02d|%02d|%02d", value/3600, (value%3600)/60, value%60)

	// var result string
	// times := value
	// temp := times / 60 / 60
	// if temp < 10 {
	// 	result += "0" + strconv.Itoa(temp) + "|"
	// } else {
	// 	result += strconv.Itoa(temp) + "|"
	// }
	// times -= temp * 60 * 60
	// temp = times / 60
	// if temp < 10 {
	// 	result += "0" + strconv.Itoa(temp) + "|"
	// } else {
	// 	result += strconv.Itoa(temp) + "|"
	// }
	// times -= temp * 60
	// if times < 10 {
	// 	result += "0" + strconv.Itoa(times)
	// } else {
	// 	result += strconv.Itoa(times)
	// }
	// return result
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
