package main
import "fmt"
const (	
	winter = 1
	summer= 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string
	books[0] = "Wings of Fire"
	books[1] = "Why I Left Milk"
	books[2] = "stay golden"
	books[3] += books[0] + " 2nd Edition"
//	fmt.Printf("books : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)
	wBooks[0] = books[0]
/*	sBooks[0] = books[1]
	sBooks[1] = books[2]
	sBooks[2] = books[3]
*/

	for i :=0; i<len(sBooks); i++ {
	//for _,v :=range sBooks {
		sBooks[i] = books[i+1]
	}
	fmt.Printf("\nwinter : %#v\n", wBooks)
	fmt.Printf("\nsummer : %#v\n", sBooks)

}
