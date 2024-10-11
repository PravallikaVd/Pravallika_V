package main

import (
	"fmt"
	"time"
)
func sent(ch chan int){
	ch <- 100
	fmt.Println("Sender: Sent 100")
	time.Sleep(200*time.Millisecond)
}
func receive(ch chan int){
	fmt.Println("Receiver: Received ", <-ch)
	time.Sleep(400*time.Millisecond)
}

func main(){
	ch := make(chan int)
	go sent(ch)
	go receive(ch)
	time.Sleep(1*time.Second)
	
}
