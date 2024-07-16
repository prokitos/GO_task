package construct

import "fmt"

func cInterfaces() {
	interfOne()
	// interfTwo()
	//  interfThr()
}

/////////////////////////////////////////////////////////////////////
// самое нормальное использование интерфейса !!!! если функции не реализованы, мы не можем присвоить экземпляр класса в переменную интерфейса.

func interfOne() {

	var curAcc BankAccount                   // создаем переменную типа интерфейса
	curAcc = &BaseAccount{ID: 0, Balance: 0} // закидываем в него экземпляр класса
	// var curAcc BankAccount = &BaseAccount{ID: 0, Balance: 0}  // можно в одну строчку вместо двух
	//curAcc = &AdvanceAccount{ID: 0, Balance: 0}	// другие классы закинуть в него не можем, так как они не наследованы (этот класс не реализовал все методы интерфейса)

	curAcc.Deposit(100) // в итоге мы обращаемся к методу интерфейса, и программа перенаправляет нас на реализованные методы класса
	fmt.Println(curAcc.GetBalance())
}

type BankAccount interface {
	Deposit(amount float64)
	GetBalance() float64
}

type BaseAccount struct {
	ID      int64   `json:"id"`
	Balance float64 `json:"balance"`
}
type AdvanceAccount struct {
	ID      int64   `json:"id"`
	Balance float64 `json:"balance"`
}

func (a *BaseAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BaseAccount) GetBalance() float64 {
	return a.Balance
}

/////////////////////////////////////////////////////////////////////
// использование интерфейса !!!! если функции интерфейса не реализованы, то экземпляр нельзя будет передать в фунцию

func interfTwo() {
	tesla := Car{name: "tesla"}
	drive(tesla) // в этом месте
}

type Vehicle interface {
	move()
}

type Car struct {
	name string
}

func drive(veh Vehicle) {
	veh.move()
}

func (c Car) move() {
	fmt.Println("car goes!!" + c.name)
}

/////////////////////////////////////////////////////////////////////
// использование интерфейса !!!! если функции не реализованы, то будет ошибка при выполненеии кода

func interfThr() {
	amdRayzen := Computer{}
	amdRayzen.TurnOn()
}

type Device interface {
	TurnOn()
}
type Computer struct {
	Device
}

func (c Computer) TurnOn() {
	fmt.Println("computer on")
}

/////////////////////////////////////////////////////////////////////
