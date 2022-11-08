package routine

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// TestBlockChannelDeadLock
func TestBlockChannelDeadLock(t *testing.T) {
	queue := make(chan int)
	queue <- 1
	got := <-queue
	fmt.Println(got)
}

func TestBlockChannelFine(t *testing.T) {
	queue := make(chan int)

	go func() {
		queue <- 1
	}()

	got := <-queue
	fmt.Println(got)
}

// TestChannelWaitGroup 启动10个goroutine每个执行结束的时间都是不固定的
// 使用wg计数add+1，done-1，wait等待全部完成
func TestChannelWaitGroup(t *testing.T) {
	queue := make(chan int)

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			// 启动一个goroutine计数加1
			wg.Add(1)
			go func() {
				// 结束一个goroutine
				defer wg.Done()
				queue <- rand.Intn(100)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			}()
		}
		wg.Wait()
		close(queue)
	}()

	for msg := range queue {
		fmt.Println(msg)
	}
}
