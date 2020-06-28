package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Println("ERROR",err)
		return
	}
	meters := feet * 0.3048
	yards := feet * 0.3333
	fmt.Printf("%g feet is %g meters.\n", feet,meters)
	fmt.Printf("%g feet is %g yards\n", feet, yards)
}
