package staire

import "fmt"

func printStair(index int, specificChar string, max int) {
	printValue := ""
	for i := 1; i <= max; i++ {
		if i >= index {
			printValue += specificChar
		} else {
			printValue += " "
		}
	}
	fmt.Println(printValue)
}

func main() {
	var stair int32 = 25
	stairInt := int(stair)
	specificChar := "#"

	for i := stairInt; i > 0; i-- {
		printStair(i, specificChar, stairInt)
	}
}
