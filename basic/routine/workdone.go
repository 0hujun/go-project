package routine

import (
	"fmt"
	"time"
)

func DoWork(done <-chan int) {
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		default:
			fmt.Println("doing work")
			time.Sleep(time.Second * 1)
		}
	}
}
