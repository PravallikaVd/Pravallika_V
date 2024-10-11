package main
import (
    "fmt"
)


func main(){
    ch := make(chan string)
	go func(){
		ch <- "Hi"
	}()
    result := <- ch
    fmt.Println("Value received: ", result)

}