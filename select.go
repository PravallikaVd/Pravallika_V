package main

import (
	"fmt"
	"time"
)

func worker(ch1 chan int, ch2 chan int) {
	ch1 <- 10
	ch1 <- 30
	time.Sleep(1 * time.Second)
	ch2 <- 20
	ch2 <- 40
	close(ch1) 
	close(ch2) 
}

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	go worker(ch1, ch2)

	
	for {
		select {
		case value1, ok := <-ch1:
			if ok {
				fmt.Println("Received from channel 1:", value1)
			} else {
				ch1 = nil
			}
		case value2, ok := <-ch2:
			if ok {
				fmt.Println("Received from channel 2:", value2)
			} else {
				ch2 = nil 
			}
		}


		if ch1 == nil && ch2 == nil {
			break
		}
	}
	time.Sleep(1 * time.Second)
}
