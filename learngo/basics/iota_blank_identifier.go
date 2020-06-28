package main
import "fmt"
func main() {
	const (
	EST = -(5 + iota)
	_	
	MST
	_
	PST
	)
	fmt.Println(EST, MST, PST)
}
