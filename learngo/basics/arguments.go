package main
import ("fmt"; "os")
func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("Arguments:", os.Args[1])
	fmt.Println("Number of arguments", len(os.Args))
}
