package main
import (
	"fmt"
	"strings"
)
func main() {
	words := strings.Fields("Hi Hello How Are You",)
	for j:=0; j<len(words);j++ {
		fmt.Printf("#%-2d:%q\n", j+1, words[j])
	}
	/*for i:=1;i<len(os.Args);i++ {
		fmt.Printf("%q\n", os.Args[i])
	}*/
}
