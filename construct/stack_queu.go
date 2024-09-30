package construct

import "fmt"

func TestStackQueue() {
	// var tempStack Stack
	// tempStack.Add(5)
	// tempStack.Add(7)
	// tempStack.Add(9)
	// tempStack.Check()
	// tempStack.Remove()
	// tempStack.Check()

	// var tempQueue Queue
	// tempQueue.Add(5)
	// tempQueue.Add(7)
	// tempQueue.Add(9)
	// tempQueue.Check()
	// tempQueue.Remove()
	// tempQueue.Check()
}

type Stack struct {
	firstNode *Node
	lastNode  *Node
}
type Queue struct {
	firstNode *Node
	lastNode  *Node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (inst *Stack) Add(value int) {
	var newNode *Node = &Node{value: value}

	if inst.lastNode == nil {
		inst.lastNode = newNode
		inst.firstNode = newNode
		return
	}

	temp := inst.lastNode
	temp.next = newNode
	newNode.prev = temp
	inst.lastNode = newNode
}
func (inst *Stack) Remove() {
	temp := inst.lastNode
	if temp.prev != nil {
		inst.lastNode = temp.prev
		inst.lastNode.next = nil
	}
}
func (inst *Stack) Check() {
	var tempNode *Node = inst.firstNode
	for tempNode != nil {
		fmt.Print(tempNode.value)
		tempNode = tempNode.next
	}
	fmt.Println()
}

func (inst *Queue) Add(value int) {
	var newNode *Node = &Node{value: value}

	if inst.lastNode == nil {
		inst.lastNode = newNode
		inst.firstNode = newNode
		return
	}

	temp := inst.lastNode
	temp.next = newNode
	newNode.prev = temp
	inst.lastNode = newNode
}
func (inst *Queue) Remove() {
	temp := inst.firstNode
	if temp.next != nil {
		inst.firstNode = temp.next
		inst.firstNode.prev = nil
	}
}
func (inst *Queue) Check() {
	var tempNode *Node = inst.firstNode
	for tempNode != nil {
		fmt.Print(tempNode.value)
		tempNode = tempNode.next
	}
	fmt.Println()
}
