package lock

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomicAdd(t *testing.T) {
	count := 1
	var flag int64
	flag = 0
	for {
		if atomic.CompareAndSwapInt64(&flag, 0, 1) {
			count++
			atomic.StoreInt64(&flag, 0)
			break
		}
	}
	fmt.Println("count is", count)
}
