package main

import "fmt"

func getDividedSet(values []int32, hightestValue int32) []int32 {
	var setDivided []int32
	var i int32

	for i = values[0]; i <= hightestValue; i++ {
		result := true

		for _, v := range values {
			if i%v != 0 {
				result = false
				break
			}
		}

		if result {
			setDivided = append(setDivided, i)
		}
	}

	return setDivided
}

func process(a []int32, b []int32) int32 {
	lastValue := a[len(a)-1]
	hightestValue := lastValue * lastValue

	getDividedA := getDividedSet(a, hightestValue)
	fmt.Println(getDividedA)

	var setDivided []int32
	for _, valueDA := range getDividedA {
		result := true

		for _, bValue := range b {

			fmt.Println(valueDA, bValue, bValue%valueDA)
			if bValue%valueDA != 0 {
				result = false
				break
			}
		}

		if result {
			fmt.Printf("got it %v\n", valueDA)
			setDivided = append(setDivided, valueDA)
		}
	}

	fmt.Println(setDivided)
	return setDivided
}

func main() {
	var a []int32 = []int32{2, 4}
	var b []int32 = []int32{16, 32, 96}

	result := process(a, b)
	fmt.Println(result)
}
