package main
import "fmt"
const (
	winter = 1	
	summer = 3
	yearly = winter + summer
)

func main() {
	books := [...]string{
	"Wings of Fire",
	"Why I Left Drinking Milk",
	"One Indian Girl",
	"Wings Of Fire 2nd Edition",
	}
	fmt.Printf("BOOKS : %#v\n", books)
}
