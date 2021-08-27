package time

import (
	"fmt"
	"strconv"
	"strings"
)

func process(s string) string {
	time := s[0:8]
	extensionTime := s[8:10]
	hour := time[0:2]
	newHour := ""

	if extensionTime == "PM" && hour == "12" {
		return time
	} else if extensionTime == "AM" && hour == "12" {
		newHour = "00"
	} else if extensionTime == "PM" {
		newHourInt := toInt(hour)
		newHour = toString(newHourInt + 12)
	} else {
		// AM case
		newHour = hour
	}

	return strings.Replace(time, hour, newHour, 1)
}

func main() {
	timeString := "07:05:45PM"
	result := process(timeString)
	fmt.Println("result")
	fmt.Println(result)
}

// convert string to int
func toInt(s string) int {
	intValue, _ := strconv.Atoi(s)
	return intValue
}

// convert int to string
func toString(i int) string {
	return strconv.Itoa(i)
}
