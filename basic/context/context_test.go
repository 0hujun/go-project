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

func doing(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(time.Second * 1)
	}
}

// TestContextDone 设置超时时间为2秒，每秒输出一次doing直到超时
func TestContextDone(t *testing.T) {
	fmt.Println("go context")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doing(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("ctx timeout done")
	}

	time.Sleep(2 * time.Second)
}
