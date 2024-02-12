package tasks

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	// res := CountPositivesSumNegatives(arr)
	// fmt.Print(res[0], "  ", res[1])

	// инвертировать регист строки
	// fmt.Println(ToAlternatingCase("hello WORLDS"))

	// убрать все пробелы из строки
	// fmt.Println(NoSpace("test testov   teest "))

	// найти ближайшее квадратное число
	// fmt.Println(NearestSq(111))

	// убрать последний символ строки
	//fmt.Println(delLatest("kekas"))

	// сумма квадратов всех чисел
	// var mass = []int{1, 2, 3}
	// fmt.Print(SquareSum(mass))

	// вернуть тру если числа в массиве в порядке возрастания
	// var numbers = []int{1, 2, 4, 3, 4, 5}
	// fmt.Print(InAscOrder(numbers))

	// сортировать массив строк по длинне
	// var mass = []string{"aaaaa", "bb", "zzzz", "ccccccc", "rrr"}
	// mass = SortByLength(mass)
	// for _, element := range mass {
	// 	fmt.Println(element)
	// }

	// найти количество делителей числа
	//fmt.Print(Divisors(100))

	// кастомная сортировка строки. каждое слово имеет цифру внутри, и цифра означает позицию в строке
	//fmt.Print(Order("is2 Thi1s T4est 3a"))

	// написать функцию которая букву перед решеткой.  abc#d = abd, так как # убрала букву c
	//fmt.Print(CleanString("abc#d##c"))

	// строку в camel case
	fmt.Print(CamelCase("hello case"))
}

func CamelCase(s string) string {

	if len(s) == 0 {
		return ""
	}

	wordsMass := strings.Split(s, " ")
	var result string

	for _, elem := range wordsMass {

		result += strings.Title(elem)
		result += ""
	}

	return result

	// заменяем старую строку и пробел, на новую строку в Title и без пробела, бесконечное количество раз
	// var result string = strings.Replace(strings.Title(s)," ","",-1)
	// var result string = strings.ReplaceAll(strings.Title(s), " ", "")

	// fields автоматически делает массив строк по пробелу, и далем в титульном виде, и собираем строку без пробела
	// var result string = Join(Fields(Title(str)), "")
}

func CleanString(s string) string {

	if len(s) == 0 {
		return ""
	}
	var result string = ""

	// перебор по входной строке
	for _, elem := range s {

		// если такой символ, то уменьшаем размер строки на (текущий размер минус 1), и только если строка не пустая
		// иначе добавляем символ в результат
		if elem == '#' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result += string(elem)
		}

	}

	return result
}

func Order(s string) string {
	if len(s) == 0 {
		return ""
	}

	// // переводим строку в массив слов
	// wordsMass := strings.Split(s, " ")
	// res := make([]string, len(wordsMass))

	// // перебираем каждое слово
	// for _, slovo := range wordsMass {
	// 	for _, symbol := range slovo {
	// 		// если слово имеет цифру, то переводим её в отдельное число
	// 		index, err := strconv.Atoi(string(symbol))
	// 		if err != nil {
	// 			continue
	// 		}

	// 		// добавляем в массив это слово, по индексу числа
	// 		res[index-1] = slovo
	// 	}
	// }

	// // собираем массив слов в строку
	// return strings.Join(res, " ")

	// переводим строку в массив слов, и указываем регекс
	wordsMass := strings.Split(s, " ")
	regx := regexp.MustCompile("[1-9]")

	// делаем кастомную сортировку
	sort.SliceStable(wordsMass, func(i, j int) bool {
		// получаем регексом число из двух слов, и слово с большим числом кидаем в конец
		num_i := regx.FindString(wordsMass[i])
		num_j := regx.FindString(wordsMass[j])
		return num_i < num_j
	})

	return strings.Join(wordsMass, " ")
}

func Divisors(n int) int {
	var counter int = 1

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			counter++
		}

	}

	return counter
}

func SortByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool { return len(arr[i]) < len(arr[j]) })
	return arr
}

func InAscOrder(numbers []int) bool {

	var result bool = true
	var old int = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if old > numbers[i] {
			result = false
			break
		}
		old = numbers[i]
	}

	return result
}

func SquareSum(numbers []int) int {
	var result int = 0
	for _, i := range numbers {
		result += i * i
	}

	return result
}

func delLatest(str string) string {
	var lastChar string = string(str[len(str)-1])
	var result string = strings.TrimSuffix(str, lastChar)
	return result
}

func NearestSq(n int) int {
	var i int = 1
	var first int = 1
	var second int = 1
	for {
		first = second
		second = i * i

		if second == n {
			return second
		}

		if second > n && first < n {
			if second-n > n-first {
				return first
			} else {
				return second
			}
		}

		i++
	}

	// быстрое решение
	// число в корень
	// потом число округлить
	// и привести к инту
}

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func ToAlternatingCase(str string) string {
	var result string
	for _, i := range str {
		if i == unicode.ToLower(i) {
			result += string(unicode.ToUpper(i))
		} else if i == unicode.ToUpper(i) {
			result += string(unicode.ToLower(i))
		} else {
			result += string(i)
		}
	}

	return result
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
