package main
import (
	"fmt"
)

func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n ; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	fmt.Println("Primes numbers from 1 to 10 are:")
	for i  := range 10 {
		if Prime(i) {
			fmt.Println(i," ")
			sum +=i
		}
	}
	fmt.Println("\nSum of prime numbers from 1 to 10 is:", sum)

}