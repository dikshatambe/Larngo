package main
import "fmt"
func main(){
	var (
	width uint8=255
	height =255
	)
	width++ //resets value to min i.e=0
	if int(width)<height {
	fmt.Printf("Height is greater\n")
	}
	fmt.Printf("HEight %d and width %d",height,width)
}
