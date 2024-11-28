package stats

import (
	"fmt"
	"errors"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minimumFunction(grades ...float64) float64 {
	min := grades[0]
	for _, grade := range grades {
		if min >= grade {
			min = grade
		}
	}
	return min
}

func maximumFunction(grades ...float64) float64 {
	max := grades[0]
	for _, grade := range grades {
		if max <= grade {
			max = grade
		}
	}
	return max
}

func averageFunction(grades ...float64) float64 {
	sum := 0.00
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func operation(op string) (func(...float64) float64, error) {
	if op == minimum {
		return minimumFunction, nil
	} else if op == average {
		return averageFunction, nil
	} else if op == maximum {
		return maximumFunction, nil
	} else {
		return nil, errors.New("Not a valid operation")
	}
}

func main() {	
	minFunc, err := operation(minimum)
	if err != nil {
		return
	}
	averageFunc, err := operation(average)
	if err != nil {
		return
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Minimum grade:", minValue, "Average grade:", averageValue, "Maximum grade:", maxValue)
}