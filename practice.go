package main
import (
	"fmt"
)
func main() {

	arr := [6]int{1,2,3,4,5,6}
	fmt.Println(arr)
	var a [5]int
	fmt.Println(a)
	fmt.Println("len:", len(a))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
		}
    }
    fmt.Println("2d: ", twoD)


}