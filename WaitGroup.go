package main
import (
	"fmt"
	"sync"
	"time"
)
func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine 1 started")
	time.Sleep(3*time.Second)
	fmt.Println("Goroutine 1 finished")
}
func task2(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Goroutine 2 started")
	time.Sleep(2*time.Second)
	fmt.Println("Goroutine 2 finished")
}
func task3(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Goroutine 3 started")
	time.Sleep(1*time.Second)
	fmt.Println("Goroutine 3 finished")
}


func main() {
	var wg sync.WaitGroup
    wg.Add(3)

	go task1(&wg)
	go task2(&wg)
	go task3(&wg)


	wg.Wait()
    fmt.Println("All goroutines finished")
}