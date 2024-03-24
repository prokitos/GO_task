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
	//fmt.Print(CamelCase("hello case"))

	// дано число. привести его к следу
	//fmt.Print(RoundToNext5(17))

	// перевести буквы в численные коды, и потом в инт
	// ABC = 65 + 66 + 67 = 656667 (INT)
	//fmt.Print(Calc("abcdef"))

	// вывести число, которое встречается чаще всего в массиве. если количество одинаковое, то выбрать макс число
	// var mass = []int{12, 10, 8, 12, 7, 6, 4, 10, 12}
	// fmt.Print(HighestRank(mass))

	// получить массив фибоначи чисел
	// var mass = fibonachiFound(7)
	// for _, element := range mass {
	// 	fmt.Println(element)
	// }

	// Добавить пробел между каждым символом в строке
	//fmt.Print(Spacify("hello world"))

	// дан массив разных цветов. сколько есть пар одинаковых цветов в этом массиве?
	//fmt.Print(NumberOfPairs([]string{"red", "blue", "red", "red", "red", "red"}))

	// посчитть количество гласных
	//fmt.Print(GetCount("hello event poll"))

	// дана строка символов A. найти N слово, и повторить его M Раз
	//fmt.Print(ModifyMultiply("hello every vodu nay", 2, 4))

	// реверснуть массив чисел
	// var mass = ReverseList([]int{3, 5, 6, 7, 8})
	// for _, element := range mass {
	// 	fmt.Println(element)
	// }

	// все буквы в ловер кейс
	//fmt.Print(toLowerCase("HELLo"))

	// отсортировать людей по росту
	// var mass = sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})
	// for _, element := range mass {
	// 	fmt.Println(element)
	// }

	// найти количество четных чисел в массиве
	//fmt.Print(findNumbers([]int{12, 345, 2, 6, 7896}))

	// дан массив из N чифр, вернуть все различные комбинации из этих чисел
	// var nums []int = []int{1, 2, 3}
	// temper := permute(nums)
	// for _, i := range temper {
	// 	for _, j := range i {
	// 		fmt.Print(j)
	// 	}
	// 	fmt.Println()
	// }

	// Дано предложение, вывести количество слов которые не содержат больших букв, или содержат только первую большую букву
	// var stroka string = "hello How aRe you CoNtestant"
	// fmt.Print(calculateValideWord(stroka)) // 3, hello How you

	// проверить что строка является окронимом массива слов
	// var stroka []string = []string{"alice", "bob", "charlie"}
	// fmt.Print(isAcronym(stroka, "abc"))

	// найти произведение (двух максимальных чисел - 1)
	// fmt.Print(maxProduct([]int{3, 4, 5, 2}))

	// заменить цифры в строке на символы. цифра меняется на прошлую букву + эта цифра
	// fmt.Println(replaceDigits("a1b2c")) // = abbdc

	// сортировать строку в порядке цифр
	// fmt.Print(sortSentence("is2 sentence4 This1 a3"))

	// отсортировать массив чтобы сначала шли четные а потом нечетные
	// nums := []int{3, 1, 2, 4}
	// temp := sortArrayByParity(nums)
	// for _, i := range temp {
	// 	fmt.Print(i)
	// }

	// умножить два самых больших числа и два самых маленьких, и найти разницу.
	fmt.Print(maxProductDifference([]int{3, 1, 2, 4}))
}

func maxProductDifference(nums []int) int {
	sort.Ints(nums)

	left := nums[0] * nums[1]
	rigth := nums[len(nums)-1] * nums[len(nums)-2]

	return rigth - left
}

func sortArrayByParity(nums []int) []int {

	// сортировка лямбдой
	// sort.Slice(nums, func(i, j int) bool { return nums[i]%2 == 0 })
	// return nums

	// сортировка в ручную
	result := make([]int, len(nums))
	first := 0
	last := len(nums) - 1

	for _, i := range nums {
		if i%2 == 0 {
			result[first] = i
			first++
		} else {
			result[last] = i
			last--
		}
	}

	return result
}

func sortSentence(s string) string {
	mass := strings.Fields(s)
	massResult := make([]string, len(mass))

	for _, item := range mass {
		temp := []rune(item)
		for i := 0; i < len(temp); i++ {
			num := temp[i] - '0'
			if num >= 0 && num <= 9 {
				temp = temp[:len(temp)-1]
				massResult[num-1] = string(temp)
			}
		}
	}

	return strings.Join(massResult, " ")
}

func replaceDigits(s string) string {

	// так как строки изменять нельзя, делаем массив рун
	str := []rune(s)

	for i := 0; i < len(s); i++ {

		// если попалась цифра, то
		if s[i] >= '0' && s[i] <= '9' {
			// берем прошлую букву, и складываем с текущей цифрой (которая получилось при отнимании нуля)
			temp := int(s[i-1]) + int(s[i]-'0')
			str[i] = rune(temp)
		}
	}

	// конвертим руны в строку и возвращаем
	return string(str)
}

func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	var first int = nums[0]
	var second int = nums[1]

	for i := 2; i < len(nums); i++ {
		cur := nums[i]
		min := min(first, second)
		max := max(first, second)

		if cur > min {
			min = cur
			first = min
			second = max
		}

	}

	return (first - 1) * (second - 1)
}

func isAcronym(words []string, s string) bool {

	if len(words) != len(s) {
		return false
	}

	for i, j := range words {
		if j[0] != s[i] {
			return false
		}
	}

	return true
}

func calculateValideWord(input string) int {
	var result int = 0

	mass := strings.Fields(input)
	for i := 0; i < len(mass); i++ {
		var counter int = 0
		temp := mass[i]

		for j := 0; j < len(mass[i]); j++ {
			if unicode.ToUpper(rune(temp[j])) == rune(temp[j]) {

				counter++
				if j > 0 {
					counter++
					break
				}

			}
		}

		if counter < 2 {
			result++
		}
	}

	return result
}

func permuteSupport(res *[][]int, val []int, position int) {

	// выполнять рукурсии до тех пор, пока вложенность рекурсии не дойдёт до количества цифр в комбинации
	if position == len(val) {
		return
	}

	// перебор по возможным вариациям
	for i := 0; i <= position; i++ {

		// создаем копию пришедших чисел, и делаем их новым объектом, чтобы массив не состоял из ссылок друг на друга
		valCopy := make([]int, len(val))
		copy(valCopy, val)

		// аналог swap из c++. меняем местами два элемента во временном массиве
		valCopy[i], valCopy[position] = valCopy[position], valCopy[i]

		// запускаем новый круг рекурсии, но с углубленностью + 1, и отправляем временный свапнутый массив
		permuteSupport(res, valCopy, position+1)

		// уникальное условие, которое игнорирует одинаковые числа на каждой итерации в углубленности рекурсии. и только потом добавляет это число в список результата
		if position > len(val)-2 {
			*res = append(*res, valCopy)
		}

	}

	// после перебора фор выход из функции
	return
}

func permute(nums []int) [][]int {
	var result [][]int

	// начинаем рекурсию, с углубленостью 0. не надо ничего возвращать так как передаем ссылку
	permuteSupport(&result, nums, 0)

	return result
}

func findNumbers(nums []int) (result int) {

	// !!!! long duration !!!!
	// for _, i := range nums {
	// 	if len(strconv.Itoa(i))%2 == 0 {
	// 		result++
	// 	}
	// }
	// return

	for _, i := range nums {
		if cheakNumEven(i) {
			result++
		}
	}
	return
}

func cheakNumEven(nums int) bool {

	var lenCount int = 0
	for nums >= 1 {
		nums = nums / 10
		lenCount++
	}

	return lenCount%2 == 0

}

func sortPeople(names []string, heights []int) (result []string) {

	// создаем структуру, чтобы данные были в одном месте
	type people struct {
		name   string
		height int
	}

	// добавляем данные в одну структуры
	tempStruct := make([]people, 0, len(names))
	for i := 0; i < len(names); i++ {
		tempCur := people{
			name:   names[i],
			height: heights[i],
		}
		tempStruct = append(tempStruct, tempCur)
	}

	// делаем кастомную сортировку лямбдой
	sort.Slice(tempStruct, func(i, j int) bool { return tempStruct[i].height > tempStruct[j].height })

	// вытаскиваем данные в переменную result
	for _, i := range tempStruct {
		result = append(result, i.name)
	}

	return

}

func toLowerCase(s string) (result string) {
	for _, i := range s {
		result += string(unicode.ToLower(i))
	}

	return result
	// return strings.ToLower(s)
}

func ReverseList(lst []int) (result []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		result = append(result, lst[i])
	}

	return result
}

func ModifyMultiply(str string, loc, num int) string {
	var stroki []string = strings.Split(str, " ")
	var result string

	for i := 0; i < num; i++ {
		result += stroki[loc] + "-"
	}

	result = result[:len(result)-1]
	return result
}

func GetCount(str string) (count int) {

	for _, i := range str {
		if i == 'u' || i == 'e' || i == 'a' || i == 'o' || i == 'i' {
			count++
		}

		// switch i {
		// case 'a', 'e', 'i', 'o', 'u':
		//   count++
		// }
	}

	return
}

// создаем переменную вывода сразу в функции, и нагружаем её
func NumberOfPairs(gloves []string) (result int) {

	dictionary := make(map[string]int, len(gloves))

	for _, i := range gloves {
		dictionary[i]++

		if dictionary[i] == 2 {
			result++
			dictionary[i] = 0
		}
	}

	return
}

func Spacify(s string) string {

	// var arr []string
	// for i := 0; i < len(s); i++ {
	// 	arr = append(arr, string(s[i]))
	// }
	// res := strings.Join(arr, " ")
	//return res

	// либо через slpit сразу получить массив символов, и потом добавить пробел к каждому элементу массива
	return strings.Join(strings.Split(s, ""), " ")
}

func fibonachiFound(num int) []int {
	var result = []int{0, 1}

	var first int = 0
	var second int = 1
	for i := 2; i < num; i++ {
		second = second + first
		first = second - first
		result = append(result, second)
	}

	return result
}

func HighestRank(nums []int) int {
	myMap := make(map[int]int)

	for _, i := range nums {
		myMap[i]++
	}

	var max int = 0
	var result int = 0
	for i, j := range myMap {
		if (j > max) || (j == max && i > result) {
			max = j
			result = i
		}
	}

	return result
}

func Calc(s string) int {
	var result int = 0
	var preResult string

	for _, i := range s {
		preResult += strconv.Itoa(int(i))
	}

	result, _ = strconv.Atoi(preResult)
	return result
}

func RoundToNext5(n int) int {
	for n%5 != 0 {
		n++
	}
	return n

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
