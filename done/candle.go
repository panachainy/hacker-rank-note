package candle

import "fmt"

func process(candles []int32) {
	var highest int64 = 0
	var sum int64 = 0

	for _, value := range candles {
		var valueInt64 int64 = int64(value)

		if highest < valueInt64 {
			highest = valueInt64
			sum = 1

		} else if highest == valueInt64 {
			sum++
		}
	}

	fmt.Println("highest")
	fmt.Println(highest)

	fmt.Println("sum")
	fmt.Println(sum)
}

func main() {
	var candles []int32 = []int32{3441267, 3441267}
	process(candles)
}
