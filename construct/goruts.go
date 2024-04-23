package construct

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var global int = 0
var mutex sync.Mutex
var atmvar int32

func cGoroutines() {

	// работа с каналами
	// chanTest()

	// работа с горутинами
	// gorutTest()

	// работа с возвратом каналов
	chanReturnTest()
}

func gorutTest() {
	// запись в переменную 1000 горутин одновременно
	global = 0
	for i := 0; i < 1000; i++ {
		go addValueToGlobal()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(global)

	// запись в переменную 1000 горутин одновременно + мьютекс
	global = 0
	for i := 0; i < 1000; i++ {
		go addValueToGlobalMutex()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(global)

	// запись в атомарную переменную 1000 горутин одновременно
	global = 0
	for i := 0; i < 1000; i++ {
		go addAtomic()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atmvar)
}

func chanTest() {

	chaIn := make(chan int, 3)
	chaOut := make(chan int, 3)
	chaIn <- 5
	chaOut <- 5
	chanOneDirection(chaIn, chaOut) // проверка каналов (только для чтения и только для записи) внутри функций

	// запись в закрытый канал
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	close(intCh) // без закрытия будет краш при попытке чтения 3 и 4 значения в канале
	// intCh <- 10	// краш при записи в закрытый канал

	// чтение из закрытого канала. все нормально
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh) // выводит значение по умолчанию (ноль), так как в канале нет больше значений
	fmt.Println(<-intCh) // выводит значение по умолчанию (ноль), так как в канале нет больше значений

	// чтение всех записей из канала, и вывод сообщения если идет чтение пустого значения
	intChy := make(chan int, 3)
	intChy <- 10
	intChy <- 3
	close(intChy) // такой способ работает только если канал закрыт
	for i := 0; i < len(intChy); i++ {
		val, opened := <-intChy
		if opened {
			fmt.Println(val)
		} else {
			fmt.Println("channel is empty")
		}
	}

}

func chanReturnTest() {
	fmt.Println(<-createChanAndReturn(5))
}

func createChanAndReturn(n int) chan int {
	ch := make(chan int, 1) // создаем канал. если мы не укажем размер, то будет ошибка дедлок
	ch <- n
	return ch

	// chans := make(chan int) // иначе добавлять через горутину
	// go func() {
	// 	chans <- n
	// }()
	// return chans
}

func chanOneDirection(chaIn chan<- int, chaOut <-chan int) {
	// все хорошо
	fmt.Print(<-chaOut)
	chaIn <- 5

	// ошибки. они возникают на этапе компиляции
	// fmt.Print(<-chaIn)
	// chaOut <- 5

}

func addValueToGlobal() {
	global = global + 1
}

func addValueToGlobalMutex() {
	mutex.Lock()
	global = global + 1
	mutex.Unlock()
}

func addAtomic() {
	atomic.AddInt32(&atmvar, 1)
}
