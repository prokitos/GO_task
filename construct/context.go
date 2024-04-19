package construct

import (
	"context"
	"fmt"
	"time"
)

func cContext() {
	fmt.Println("-------------------")

	// мы запускаем горутину, и ждем либо пока она сделается, либо пока пройдёт 2 секунды.
	// после этого мы идем дальше, и отменяем горутину если она не закончилась
	TimeoutCont()

	// мы запускаем две горутины, с контекстом withCancel
	// вторая горутина делается за секунду, и после этого отменяет первую
	CancelCont()

	time.Sleep(time.Second * 8)
	fmt.Println("-------------------")
}

func TimeoutCont() {

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	var nchan chan int = make(chan int)

	go testGOfunc(ctx, nchan)

	select {
	case <-ctx.Done():
		fmt.Println("goruntina timeout")
	case <-nchan:
		fmt.Println("goruntina complete")
	}
}

func CancelCont() {

	ctx, cancel := context.WithCancel(context.Background())

	var nchan chan int = make(chan int)
	go testGoFast(ctx, nchan)
	go testGOfunc(ctx, nchan)

	select {
	case <-ctx.Done():
		fmt.Println("goruntina timeout")
		cancel()
	case <-nchan:
		fmt.Println("goruntina complete")
		cancel()
	}
}

func testGOfunc(ctx context.Context, cchan chan int) {

	fmt.Println("gorutina start")
	time.Sleep(time.Second * 5)

	// если конекст отменен, то уходим
	if ctx.Err() != nil {
		return
	}

	fmt.Println("gorutina end")
	cchan <- 2

}

func testGoFast(ctx context.Context, cchan chan int) {

	time.Sleep(time.Second * 1)
	cchan <- 2
}
