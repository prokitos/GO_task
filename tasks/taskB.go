package tasks

import (
	"fmt"
)

func MainB() {

	//slovarTest()
	//sortTest()
	//convertTest()
	// printEveryType("don 16")
}

func slovarTest() {
	dict := make(map[string]string)

	dict["pen"] = "ruchka"
	dict["rand"] = "random"

	for i := range dict {
		fmt.Print(i, "  ")
		fmt.Println(dict[i])
	}
}

func sortTest() {
	//mass := []int{0, 4, 7, 9, 1, 4, 6}
	// mass := []string{"aa", "c", "dd", "aaaaa", "bbb"}
	//sort.Strings(mass)
	// sort.Slice(mass, func(x, y int) bool {
	// 	return len(mass[x]) > len(mass[y])
	// })
	//sort.Ints(mass)
	//sort.Sort(sort.Reverse(sort.IntSlice(mass)))
	// for i := range mass {
	// 	fmt.Println(mass[i])
	// }

}

func convertTest() {

	// var temp int = 10
	// var str string = strconv.Itoa(temp)
	// fmt.Println(str)

	// var temp string = "10"
	// chislo, _ := strconv.Atoi(temp)
	// fmt.Println(chislo)

}

func printEveryType[T any](s T) {
	fmt.Println(s)
}
