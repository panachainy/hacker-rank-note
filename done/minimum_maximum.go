package minimum

import "fmt"

func getSum(inputs []int32, index int) int64 {
	var sum int64

	for i, value := range inputs {
		if i != index {
			sum += int64(value)
		}
	}

	return sum
}

func main() {
	inputs := []int32{769082435, 210437958, 673982045, 375809214, 380564127}
	var maxSum int64
	var minSum int64

	var result int64 = 0

	for i, _ := range inputs {
		// do
		if i == 0 {
			result = getSum(inputs, i)
			maxSum = result
			minSum = result
		} else {
			// while
			result = getSum(inputs, i)

			if result >= maxSum {
				maxSum = result
			} else {
				if result <= minSum {
					minSum = result
				}
			}
		}
	}

	fmt.Printf("%v %v", minSum, maxSum)
}
