package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx2, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	data := "abc"
	longRunningThingManager(ctx2, data)
}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}
	ch := make(chan wrapper, 1)
	go func() {
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{result, err}
	}()
	select {
	case data := <-ch:
		fmt.Println("ok")
		return data.result, data.err
	case <-ctx.Done():
		fmt.Println("ng")
		return "", ctx.Err()
	}
}

func longRunningThing(ctx context.Context, data string) (string, error) {
	fmt.Println("Sleep at two seconds")
	time.Sleep(2 * time.Second)
	return data, nil
}
