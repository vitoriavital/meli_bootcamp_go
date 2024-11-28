package average

import "fmt"

func calculateAverage(grades ...float64) float64 {
	var sum float64
	size := len(grades)

	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(size)
}

func main() {
	fmt.Println("The students average is:", calculateAverage(9.5, 6.2, 7.3, 4.5))
}