package main

import (
	"fmt"
	"moduleA/tasks"
)

func main() {
	fmt.Println("start")

	// tasks.MainA()
	// tasks.MainB()
	// tasks.MainC()
	// tasks.MainD2()
	// tasks.MainE()
	// tasks.MainF()
	tasks.MainG()
	// fmt.Scan()

	//construct.MainConstruct()

	fmt.Println("end of programm")
}

type Tstru struct {
	Name string
	Age  int
}
