package tasks

import (
	"fmt"
	"sort"
	"strings"
)

func MainC() {

	// строку в массив строк
	// strToMass("hey yourk")

	// сортировка по длинне строк
	// strCustomSort([]string{"hello", "jump", "may", "online"})

	// возвести все числа массива в квадрат
	// intQuadMass([]int{1, 2, 3, 4, 5})

	// копировать четные элементы в Один массива, а нечетные в другой
	// intSortOddEven([]int{1, 2, 2, 4, 5, 6, 7, 8})

	// сумма только положительных чисел
	// sumOfPositive([]int{1, 2, -3, 4, -5})

	// удалить все дубликаты из массива
	// delAllRepeat([]int{1, 7, 7, 2, 3, 4, 3, 7})

	// удалить дубликаты которые идут подряд
	// delMiniRepeat([]int{1, 1, 7, 7, 2, 3, 4, 4, 4, 3, 7})

	// удалить цифры 2 из массива.
	delSpecNum([]int{1, 7, 7, 2, 3, 4, 3, 2}, 2)
}

func delSpecNum(input []int, specNum int) {
	newitems := []int{}

	for _, i := range input {
		if i != specNum {
			newitems = append(newitems, i)
		}
	}

	for _, i := range newitems {
		fmt.Print(i, " ")
	}
}

func delMiniRepeat(input []int) {

	// если две одинаковые цифры подряд, то удаляем одну, и на один шаг назад.
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			input = append(input[:i], input[i+1:]...)
			i--
		}
	}

	for _, i := range input {
		fmt.Print(i, " ")
	}
}

func delAllRepeat(input []int) {

	// делаем словарь и новый массив для результатов
	globalMap := make(map[int]bool)
	tempRes := make([]int, 0)

	// храним результат в новом массиве, так как словарь хранит данные в несортированом порядке.
	for _, i := range input {
		if globalMap[i] == false {
			tempRes = append(tempRes, i)
		}
		globalMap[i] = true
	}

	for _, i := range tempRes {
		fmt.Print(i, " ")
	}
}

func sumOfPositive(input []int) {
	var result int = 0

	for _, i := range input {
		if i > 0 {
			result += i
		}
	}

	fmt.Println(result)
}

func strToMass(input string) {
	mass := strings.Fields(input)
	for _, i := range mass {
		fmt.Println(i)
	}
}

func strCustomSort(input []string) {
	sort.Slice(input, func(i, j int) bool { return len(input[j]) < len(input[i]) })
	for _, i := range input {
		fmt.Println(i)
	}
}

func intQuadMass(input []int) {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = v * v
	}

	for _, i := range result {
		fmt.Println(i)
	}
}

func intSortOddEven(input []int) {

	sort.Slice(input, func(i, j int) bool { return input[i]%2 == 0 })

	maxLen := len(input)
	midLen := 0

	// найти точку, где четные переходят в нечетные
	for i := 0; i < maxLen; i++ {
		if input[i]%2 != 0 {
			midLen = i
			break
		}
	}

	// вывод четных
	odd := input[0:midLen]
	fmt.Println("")
	fmt.Print("odd ")
	for _, i := range odd {
		fmt.Print(i, " ")
	}

	// вывод нечетных
	even := input[midLen:maxLen]
	fmt.Println("")
	fmt.Print("even ")
	for _, i := range even {
		fmt.Print(i, " ")
	}
}
