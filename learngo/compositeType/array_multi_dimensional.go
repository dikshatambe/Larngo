package main
import "fmt"
func main() {
	students := [...][3]float64{
		{3,6,4},								  {4,8,9}}

	var sum float64
	//sum += students[0][0] + students[0][1] + students[0][2]
	//sum += students[1][0] + students[1][1] + students[1][2]
	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		} 
	}
	const N = float64(len(students) * len(students[0]))

	fmt.Printf("Average is: %g", sum/N)
} 
