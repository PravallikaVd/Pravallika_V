package main

import "fmt"

func main() {
    c := make(chan int, 3) // Buffered channel with capacity 3

    c <- 10
    c <- 20
	c <- 30


    fmt.Println("Value received : ",<-c) // 1
    
    fmt.Println("Value received : ",<-c) //2
	
    fmt.Println("Value received : ",<-c) //3
}

