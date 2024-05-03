package construct

import (
	"fmt"
)

func cSliceArray() {

	temp := [4]string{"tesla", "volvo", "toyota", "tesla"}    // создание массива из 4 элементов
	cars1 := [...]string{"tesla", "volvo", "toyota", "tesla"} // создание МАССИВА из 4 элементов. тоже самое что сверху, но не надо указывать количество
	cars2 := []string{"tesla", "volvo", "toyota", "tesla"}    // создание слайса  из 4 элементов
	// cars1 = append(cars1, "lada")  // ошибка при компиляции, нельзя добавить элемент в массив
	// cars2 = append(cars2, "lada")  // все хорошо, можно добавить элемент в слайс

	fmt.Println(temp)
	fmt.Println(cars1)
	fmt.Println(cars2)

	// удобная передача данные в функцию, и получение слайса внутри функции
	testArray("tesla", "lada")
	// не очень удобная передача данных в функцию, но зато удобно уже готовые массивы. внутри тоже будет слайс
	testArray2([]string{"tesla", "lada"})

	// добавление одного слайса в другой слайс
	a1 := []int{0, 1, 2, 3, 4}
	a2 := []int{5, 6, 7, 8}
	// аппенд добавляет по одному элементу, но через оператор упаковки можно сразу слайс
	a1 = append(a1, a2...)
	fmt.Println(a1)

	massTest("a", []string{"a", "b", "c"})
	massTest(5, []int{1, 2, 3})
	massTest(5, []int{1, 2, 5})

	genTest("z", []int{1, 2, 3})
}

func testArray(input ...string) {
	input = append(input, "kalina")
	fmt.Println(input)
}

func testArray2(input []string) {
	input = append(input, "kalina")
	fmt.Println(input)
}

// проверка есть ли значение в слайсе
func massTest[T comparable](val T, values []T) {
	for _, i := range values {
		if val == i {
			fmt.Println("its contain value in slice")
			return
		}
	}

	fmt.Println("no value in slice")
}

// тест дженериков
func genTest[T any, U any](val U, values []T) {

	// переводим букву z в строку, а потом в число
	var len = len(values)
	comp := fmt.Sprintf("%+v", val)
	cmp := []byte(comp)

	if int(cmp[0]) > len {
		fmt.Println("bigger")
	} else {
		fmt.Println("smaller")
	}

}
