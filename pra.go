package main
import (
	"fmt"
)

func Prime(num int) bool {

	if num <=1 {
		return false
	}

	for i :=2; i*i <= num ; i++ {
		if num%2 == 0 {
			return false
		}
	}
	return true
}

func main () {
	sum :=0
	for i := range 20 {
	if Prime(i) {
		fmt.Println(i)
		sum += i
	}
}
	fmt.Println(sum)
}