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
