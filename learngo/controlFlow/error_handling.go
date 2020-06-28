package main
import (
	"fmt"
	"strconv"
	"os"
)
func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number :",n)
	fmt.Println("Error message :", err)
	
}
