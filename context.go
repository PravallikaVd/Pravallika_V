package main

import (
	"context"
	"fmt"
	"time"
)

func Task(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): 
		fmt.Println("Task completed")
	case <-ctx.Done(): 
		fmt.Println("Task was canceled:", ctx.Err())
	}
}

func main() {
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	
	go Task(ctx)
	time.Sleep(2 * time.Second)
	cancel() 

	
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}
