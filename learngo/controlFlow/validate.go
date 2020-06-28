package main
import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args)!=3 {
		fmt.Println("usage: [username] [password]")
		return
	}
	u, p := args[1], args[2]
	if u!= "abc" {
		fmt.Println("Access denied for: ",u)
	} else if p != "1234" {
		fmt.Println("Invalid password for:", u)
	} else {
		fmt.Println("ACCESS GRANTED")
	}
}
