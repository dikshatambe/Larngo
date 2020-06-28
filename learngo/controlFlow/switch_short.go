package main
import "fmt"
func main() {
	switch i := 10; {
	case i>0:
		fmt.Println("Positive")
	case i<0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
