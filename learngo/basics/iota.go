package main
import "fmt"
func main() {
	const (
	monday = iota
	tuesday
	wed
	thurs
	fri
	sat
	sun
	)
	fmt.Println(monday,tuesday,wed,thurs,fri,sat,sun)
}
