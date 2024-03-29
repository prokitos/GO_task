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
	intSortOddEven([]int{1, 2, 3, 4, 5, 6, 7, 8})

	// удалить дубликаты которые идут подряд

	// удалить все дубликаты из массива

	// удалить цифры 2 из массива.

	// сумма только положительных чисел
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

	for _, i := range input {
		fmt.Println(i)
	}
}
