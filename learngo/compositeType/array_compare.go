package main
import "fmt"
func main() {
	var (
		blue = [3]int{1,2,3}
		red  = [3]int{1,2,3}
	)

	fmt.Printf("blue bookcase : %v\n",blue)
	fmt.Printf("red bookcase  : %v\n",red)
	fmt.Println("are they equal?",blue == red)
}	
