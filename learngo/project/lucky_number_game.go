package main
import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)
const (
	maxTurns = 5
	usage = 'Welcome to the Lucky Number Game! The program will pick %d random numbers.Your mission is to guess one of those numbers.The greateryour number is, the harder it gets.Wanna Play?'
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) !=1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess <0 {
		fmt.Println("Plese select a positive number")
		return 
	}

	for turn :=0 ; turn<maxTurns;turn++ {
		n := rand.Intn(guess+1)

		if n == guess {
			fmt.Println("YOU WON!!")
			return
		}
	}
	fmt.Println("YOU LOST..TRY AGAIN!!")
}
