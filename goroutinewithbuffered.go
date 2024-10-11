package main
import (
	"fmt"
	"time"
)

func goroutine1(ch chan int){
	ch <- 5
	time.Sleep(100*time.Millisecond)
}
func goroutine2(ch chan int){
	ch <- 10
	time.Sleep(200*time.Millisecond)
}
func goroutine3(ch chan int){
	ch <- 15
	time.Sleep(300*time.Millisecond)
}

func main(){ 
	ch := make(chan int, 3)
    
	go goroutine1(ch)
	go goroutine2(ch)
	go goroutine3(ch)

	fmt.Println("Goroutine 1 sent ", <-ch)
	fmt.Println("Goroutine 2 sent ", <-ch)
	fmt.Println("Goroutine 3 sent ", <-ch)
    time.Sleep(1*time.Second)
}