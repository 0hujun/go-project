package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request_id", "123")
}

func doSomethingCool(ctx context.Context) {
	id := ctx.Value("request_id")
	fmt.Println(id)
}

func TestContextValue(t *testing.T) {
	fmt.Println("go context")
	ctx := context.Background()

	ctx = enrichContext(ctx)

	doSomethingCool(ctx)
}

func doing(ctx context.Context, done chan bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(time.Second * 1)
		done <- true
	}
}

// TestContextDone 设置超时时间为2秒，每秒输出一次doing直到超时
func TestContextDone(t *testing.T) {
	fmt.Println("go context")
	// 构造一个超时器
	after := time.After(time.Duration(5) * time.Second)
	// 构造一个2秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// 构造一个通道与goroutine通信
	done := make(chan bool)
	// 函数最后推出context
	defer cancel()
	// 处理事物
	go doing(ctx, done)
	select {
	case <-ctx.Done():
		fmt.Println("ctx timeout done")
	case <-after:
		fmt.Println("ctx timeout")
	case <-done:
		fmt.Println("we done first")
	}

	time.Sleep(2 * time.Second)
}
