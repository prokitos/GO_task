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
