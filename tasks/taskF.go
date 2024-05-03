package tasks

import "fmt"

func MainF() {

	// дана последовательность букв. найти самую большую длину последовательности, в которой буквы не повторяются.

	var temp string = "abcabvgcdfbaf"
	var fast int = 0
	var slow int = 0
	var res int = 0
	var totalRes int = 0

	curSet := make(map[int]emptyStruct)

	for fast < len(temp) {
		fastChar := temp[fast]
		slowChar := temp[slow]
		_, ok := curSet[int(fastChar)]

		// если уже есть такое значение в сете
		if ok {
			delete(curSet, int(slowChar))
			res--
			slow++
		} else {
			curSet[int(fastChar)] = emptyStruct{}
			res++
			fast++
		}
		if totalRes < res {
			totalRes = res
		}
	}

	// ответ 7.   a b c (a b v g c d f) b a f
	fmt.Println(totalRes)

}

// пустая структура, которая занимает месат меньше чем bool. нужна для более легковесной реализации set
type emptyStruct struct{}
