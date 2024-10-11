package main
import (
	"fmt"
	"time"
)

func numbers() {
	for i :=0; i <= 5; i++ {
		fmt.Println("ID  : ", i)
		time.Sleep(100*time.Millisecond)
	}
}

func strings() {

	fmt.Println("START")
	str :=[4]string{"Bob","Alice","Joe","Ray"}
	for i :=0; i < len(str) ; i++ {
		fmt.Println("Name: ", str[i])
		time.Sleep(200*time.Millisecond)
	}
	fmt.Println("END")
}

func main() {

	go numbers()
	go strings()
	time.Sleep(1*time.Second)


}