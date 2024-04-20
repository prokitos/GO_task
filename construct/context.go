package construct

import (
	"context"
	"fmt"
	"time"
)

// https://www.youtube.com/watch?v=fXzzF5y6UEU
// https://dzen.ru/a/YRzqvr7vuXsMnjG0

func cContext() {
	fmt.Println("-------------------")

	// мы запускаем горутину, и ждем либо пока она сделается, либо пока пройдёт 2 секунды.
	// после этого мы идем дальше, и отменяем горутину если она не закончилась
	// TimeoutCont()

	// мы запускаем две горутины, с контекстом withCancel. если одна сделалась, отменяем другие.
	// вторая горутина делается за секунду, и после этого отменяет первую
	// CancelCont()

	// мы запускаем горутину которая отменится через 2 секунды, и внутри этой горутины запускаем ещё одну, но во вторую горутину мы пробросим contextWithoutCancel
	// если во вторую горутину пробросить тотже котекст, то она тоже отменится. если наследоваться withCancel, или withTimeout(10hour), то отмена все ещё будет
	// так как родительский контекст более приоритетный. решение это contextWithoutCancel который наследуется от родителя, и продолжает работу даже если родительский контекст отменен.
	NoCancel()

	time.Sleep(time.Second * 8)
	fmt.Println("-------------------")
}

func NoCancel() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	var nchan chan int = make(chan int)

	go testGoCanceler(ctx, nchan)
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

func testGoCanceler(ctx context.Context, cchan chan int) {

	fmt.Println("gorutina start 1")
	// ctxN, _ := context.WithTimeout(ctx, 5*time.Second)	// оно не перезапишет, и старый таймаут будет более приорететный, в итоге в новой функции контекст отменится вместе с этим
	// ещё сработает создание абсолютно нового контекста, но так не по нормам. по нормам нужно наследоваться от пришедшего контекста.
	ctxN := context.WithoutCancel(ctx) // оно создаст контекст на основе этого, и новый контекст не отменится внутри новой функции

	go testGoNoCancel(ctxN, cchan)
	time.Sleep(time.Second * 3)

	// если конекст отменен, то уходим
	if ctx.Err() != nil {
		return
	}

	fmt.Println("gorutina end 1")
}

func testGoNoCancel(ctx context.Context, cchan chan int) {
	fmt.Println("gorutina start 2")
	time.Sleep(time.Second * 3)

	if ctx.Err() != nil {
		return
	}

	fmt.Println("gorutina end 2")
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
