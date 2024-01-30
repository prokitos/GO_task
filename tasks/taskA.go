package tasks

import (
	"fmt"
	"strconv"
)

func MainA() {

	// вернуть массив значений, между этими числами
	// var mass []int = Between(1, 5)
	// for element := range mass {
	// 	fmt.Println(element)
	// }

	// вернуть массив чисел кратных n
	// var mass []int = FindMultiples(5, 25)
	// for _, element := range mass {
	// 	fmt.Println(element)
	// }

	// вернуть массив чисел кратных n в обратном порядке
	// var mass []int = FindMultiples(5, 25)
	// for i := range mass {
	// 	fmt.Println(mass[len(mass)-1-i])
	// }

	// конверт строки в число
	// fmt.Print(StringToNumber("102"))

	// написать указанную строку N раз
	// fmt.Print(RepeatStr(10, "ha"))

	// подсчитать сумму негативных и количество позитивных чисел
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	res := CountPositivesSumNegatives(arr)
	fmt.Print(res[0], "  ", res[1])

}

func CountPositivesSumNegatives(numbers []int) []int {
	mass := []int{0, 0}

	for _, i := range numbers {
		if i > 0 {
			mass[0]++
		} else {
			mass[1] += i
		}
	}

	return mass
}

func RepeatStr(repetitions int, value string) string {

	var result string
	for i := 0; i < repetitions; i++ {
		result += value
	}

	//strings.Repeat(value, repititions)
	return result
}

func StringToNumber(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		// Лог ошибки
		return 1
	}
	return i
}

func Between(a, b int) []int {

	temp := []int{}

	for i := a; i <= b; i++ {
		temp = append(temp, i)
	}

	return temp
}

func FindMultiples(integer, limit int) []int {
	temp := []int{}

	for it := integer; it <= limit; it += integer {
		temp = append(temp, it)
	}

	return temp
}
