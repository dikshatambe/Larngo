package main
import (
	"os"
	"fmt"
	"strings"
)

const corpus =""+ "lazy cats jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
	search:
		for i,w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}
			if q==w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				//break //to skip duplicate
				continue queries //break from parent loop
			}
		}
	}
}
