package tasks

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func MainB() {

	//slovarTest()
	//sortTest()
	//convertTest()
	//printEveryType("don 16")

	// oopTest()
	// binaryTreeTest()
	// gourutinesTest()

	// var mass = []int{0, 1, 2, 3, 4, 5}
	// testMass(mass)
	// var mass2 = testMass2(mass)

	// for i := range mass2 {
	// 	fmt.Println(mass2[i])
	// }

	oopTest()

}

func oopTest() {
	man := Persons{"Jack", 15, 100}
	man.AddMoney()  // + 100
	AddMoney2(&man) // + 200
	//fmt.Print(man.money) // = 400

	// тест наследования и интерфейса
	var cman CoolPersons
	cman.Persons = man       // 400
	cman.AddMoney()          // + 200 так как он крутой
	AddMoney2(&cman.Persons) // + 200
	fmt.Print(cman.money)    // = 800

}

func binaryTreeTest() {
	var mainTree binaryTree
	mainTree.AddNode(10)
	mainTree.AddNode(15)
	mainTree.AddNode(5)
	fmt.Println(mainTree.root.value)
}

func gourutinesTest() {

	start := time.Now()

	chann1 := make(chan int)
	chann2 := make(chan int)

	go timerCheck(chann1)
	go timerCheck(chann2)

	<-chann1
	<-chann2

	duration := time.Since(start)
	fmt.Println(duration)

	// https://blog.ildarkarymov.ru/posts/go-concurrency/
}

type binaryTree struct {
	root *nodeTree
}

type nodeTree struct {
	value     int
	leftNode  *nodeTree
	rightNode *nodeTree
}

func (thisTree *binaryTree) AddNode(val int) {

	if thisTree.root == nil {
		thisTree.root = &nodeTree{val, nil, nil}
		return
	}

	if thisTree.root != nil {

		temp := thisTree.root
		for true {
			if val > temp.value {
				if temp.rightNode == nil {
					temp.rightNode = &nodeTree{val, nil, nil}
					return
				}
				temp = temp.rightNode
			} else {
				if temp.leftNode == nil {
					temp.leftNode = &nodeTree{val, nil, nil}
					return
				}
				temp = temp.leftNode
			}
		}

	}

}

func testMass2(temp []int) []int {

	var masser = temp

	for i := range masser {
		masser[i] += 10
	}

	return masser
}

func testMass(temp []int) {

	for i := range temp {
		temp[i] += 10
	}

}

type IPerson interface {
	AddMoney()
}

type CoolPersons struct {
	Persons
}

type Persons struct {
	name  string
	age   int
	money int
}

func (x *Persons) AddMoney() {
	x.money += 100
}

func AddMoney2(x *Persons) {
	x.money += 200
}

func (x *CoolPersons) AddMoney() {
	x.money += 200
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

func timerCheck(zaglushka_kanala chan int) {

	var mass []int
	for i := 0; i < 10000; i++ {
		mass = append(mass, i)
	}

	rand.Shuffle(len(mass), func(i, j int) {
		mass[i], mass[j] = mass[j], mass[i]
	})

	sort.Ints(mass)

	// засечь сколько времиени с горутинами

	// for i := range mass {
	// 	fmt.Print(mass[i], "  ")
	// }

	zaglushka_kanala <- 1 // говорим что корутина закончена
}
