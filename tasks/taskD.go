package tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

func MainD() {

	// добавление json структуры в текстовик
	var curUser user
	curUser.Age = 15
	curUser.Name = "john"
	curUser.Value = 1001

	b, _ := json.Marshal(&curUser)

	// авто создание файла, чтение, и добавление
	readFile(b)

}

type user struct {
	Name  string
	Value int
	Age   int
}

func readFile(stroka []byte) {
	b, err := os.ReadFile("new.json")
	if err != nil {
		panic("")
	}

	c := []byte(stroka)
	res := append(b, c...)

	os.Stdout.Write(res)

	err = os.WriteFile("new.json", res, 0644)
	if err != nil {
		panic("")
	}
}

func MainD2() {

	// создание стека
	var stack Stackk
	stack.push(5)
	stack.push(7)
	stack.push(9)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())

	fmt.Println("  ")
	fmt.Println("end")
	fmt.Println("  ")

	// создание очереди
	var que Queue
	que.push(10)
	que.push(9)
	que.push(6)

	fmt.Println(que.pop())
	fmt.Println(que.pop())
	fmt.Println(que.pop())
	fmt.Println(que.pop())
	fmt.Println(que.pop())
}

type StackNode struct {
	value int
	prev  *StackNode
}
type Stackk struct {
	last *StackNode
}

// добавить в стек
func (cur *Stackk) push(value int) {
	var newNode StackNode
	newNode.value = value

	if cur.last == nil {
		cur.last = &newNode
		return
	}

	temp := cur.last
	cur.last = &newNode
	cur.last.prev = temp

}

// убрать из стека
func (cur *Stackk) pop() int {
	if cur.last == nil {
		return 0
	}
	var result int = cur.last.value

	temp := cur.last.prev
	cur.last = temp

	return result
}

// посмотреть что наверхе стека
func (cur *Stackk) peak() int {
	return cur.last.value
}

type QueueNode struct {
	value int
	next  *QueueNode
	prev  *QueueNode
}

type Queue struct {
	first *QueueNode
	last  *QueueNode
}

func (cur *Queue) push(val int) {
	var element QueueNode
	element.value = val

	if cur.last == nil {
		cur.first = &element
		cur.last = &element
		return
	}

	temp := cur.last
	cur.last = &element
	temp.next = &element
	element.prev = temp

}
func (cur *Queue) pop() int {
	if cur.first == nil {
		return 0
	}

	var res int = cur.first.value

	temp := cur.first
	cur.first = temp.next

	return res
}
func (cur *Queue) peak() int {
	return cur.first.value
}
