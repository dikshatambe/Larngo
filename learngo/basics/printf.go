package main
import "fmt"
func main() {
	var (
	planet = "venus"
	distance = 261
	orbital = 224.701
	hasLife = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v million kms\n", distance)
	fmt.Printf("Orbital period: %.1f days\n", orbital)
	fmt.Printf("Does %v have life?: %v\n", planet, hasLife)
}
