package main
import "fmt"
import "os"
const (
	usage = "Usage: [username] [password]"
	errUser = "Access denied for %q.\n"
	errPwd = "Invalid password for %q.\n"
	accessOK = "Access granted for %q.\n"

	user, user2 = "abc", "pqr"
	pass, pass2 = "1234", "5678"
)
func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1],args[2]
	if u! = user && u!= user2 {
		fmt.Printf(errUser, u)
	} else if u==user && p == pass || u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}
}
