package main
import (
	"fmt"
	"my_module/calculator"
)
func  main() {
     x,y := 8, 2
	 fmt.Println("Addition of x and y is", calculator.Add(x,y))
	 fmt.Println("Subtraction of x and y is", calculator.Subtract(x,y))
	 fmt.Println("Multiply of x and y is", calculator.Multiply(x,y))
	 fmt.Println("Divide of x and y is", calculator.Divide(x,y))

}