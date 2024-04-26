package construct

import "fmt"

func cDefers() {
	// panicReconver()
	printTest()
}

func printTest() {

	// defer ложатся по принципу стека
	defer fmt.Println(10)
	defer fmt.Println(5)

	// defer запоминает значение которое было (делает копию на этапе создания defer)
	var x int = 10
	defer fmt.Println(x)
	x += 5

}

func panicReconver() {

	// panic("err") // если паника будет до recover, то случится краш

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	panic("err")

}
