package construct

import "fmt"

func cPointer() {

	smartBase := Smartphone{"basic smart"}

	// изначально оба гика ссылаются на одно и тоже значение
	geek1 := Geek{&smartBase}
	geek2 := Geek{&smartBase}
	fmt.Println("----------------------")
	fmt.Println(geek1.smartphone.name)
	fmt.Println(geek2.smartphone.name)

	// мы пытаемся поменять указатель внутри функции, но не получается
	poorReplace(geek1.smartphone)
	fmt.Println("----------------------")
	fmt.Println(geek1.smartphone.name)
	fmt.Println(geek2.smartphone.name)

	// мы меняем оригинальное значение, на которое ссылается указатель, поэтому у всех гигов теперь другой телефон
	basicReplace(geek1.smartphone)
	fmt.Println("----------------------")
	fmt.Println(geek1.smartphone.name)
	fmt.Println(geek2.smartphone.name)

	// мы меняем ссылку на значение у указателя, и он теперь ссылается другую область в памяти, поэтому телефон меняется только у одного гика
	coolReplace(&geek2.smartphone)
	fmt.Println("----------------------")
	fmt.Println(geek1.smartphone.name)
	fmt.Println(geek2.smartphone.name)
	fmt.Println("----------------------")

	mapAndSliceTest()
	testTransfer()
}

func testTransfer() {
	var testArray [50]int
	testArray[0] = 101
	arrayChanger(testArray)
	fmt.Println(testArray[0])

	maper := make(map[int]int)
	maper[0] = 101
	mapChanger(maper)
	fmt.Println(maper[0])

	var testSlice []int = make([]int, 100)
	testSlice[0] = 101
	sliceChanger(testSlice)
	fmt.Println(testSlice[0])

}

func mapChanger(input map[int]int) {
	input[0] = 102
}

func arrayChanger(input [50]int) {
	input[0] = 102
}

func sliceChanger(input []int) {
	input[0] = 102
}

func mapAndSliceTest() {

	// нельзя взять указатель значения в мапе из-за эвакуации данных
	var temp map[int]int = make(map[int]int)
	temp[5] = 10
	// var pointers *int = &temp[5] // Ошибка при компиляции

	//////////////////////////////

	// можно получить указатель значения в слайсе, но он может быть неверным
	var slices1 []int = []int{99, 1, 2, 3}

	// мы получаем указатель на первое значение в слайсе (99)
	var pointers1 *int = &slices1[0]

	// добавляем много значение в слайс, тем самым вызывая перераспределение памяти и создание нового слайса с большим размером
	for i := 0; i < 100; i++ {
		slices1 = append(slices1, i)
	}

	// в итоге наш указатель ссылается на старый слайс который уже не используется, и вывод будет (99) вместо (101)
	slices1[0] = 101
	fmt.Println(*pointers1)

	//////////////////////////////

	// но можно зарезервировать слайс заранее
	var slices []int = make([]int, 110)
	placer := []int{99, 1, 2, 3}
	// добавить в него нужные данные
	slices = append(slices, placer...)
	// взять указатель
	var pointers *int = &slices[0]

	// заполнить мусором на 100 (что входит в резерв 110)
	for i := 0; i < 100; i++ {
		slices = append(slices, i)
	}

	// и в итоге указатель валиден, потомучто из 110 выделенной памяти использовалось только 104
	slices[0] = 101
	fmt.Println(*pointers)
}

// есть некий смартфон
type Smartphone struct {
	name string
}

// гик носит некий смартфон
type Geek struct {
	smartphone *Smartphone
}

// при передаче любого параметра в функцию, мы передаем его копию

// мы пытаемся поменять копию указателя, и в итоге в оригинале ничего не меняется
func poorReplace(s *Smartphone) {
	s = &Smartphone{"poor smart"}
}

// здесь мы перетираем оригинальные значенния obichniy на krutoi, и указатель указывает на тоже место в памяти.
func basicReplace(s *Smartphone) {
	*s = Smartphone{"normal smart"}
}

// здесь мы выделяем новую область памяти для krutoi, и наш указатель указывает на новую область памяти.
func coolReplace(s **Smartphone) {
	*s = &Smartphone{"cool smart"}
}
