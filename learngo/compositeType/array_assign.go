package main
import "fmt"
func main() {
	blue := [3]int{1,2,3}
	red := blue
	fmt.Printf("blue array : %#v\n", blue)
	fmt.Printf("red array  : %#v\n", red)
	blue[0] = 8
	fmt.Printf("blue array : %#v\n", blue)
        fmt.Printf("red array  : %#v\n", red)

}
