package tasks

import "fmt"

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
	var mass []int = FindMultiples(5, 25)
	for i := range mass {
		fmt.Println(mass[len(mass)-1-i])
	}

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
