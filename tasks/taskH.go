package tasks

import (
	"fmt"
	"time"
)

var temp int

func MainH() {
	for i := 0; i < 1000; i++ {
		go func() {
			temp++
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(temp)
}


// до 1.22    = выведет 980
// после 1.22 = выведет 1000
