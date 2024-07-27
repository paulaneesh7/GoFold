package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Printf("Format of currentTime: %T\n", currentTime) // time.Time format in Go

	// All the built-in methods of time package in Go
	fmt.Println("Current time is: ", currentTime)
	fmt.Println("Year: ", currentTime.Year())
	fmt.Println("Month: ", currentTime.Month())
	fmt.Println("Day: ", currentTime.Day())
	fmt.Println("Hour: ", currentTime.Hour())
	fmt.Println("Minute: ", currentTime.Minute())
	fmt.Println("Second: ", currentTime.Second())
	fmt.Println("Nanosecond: ", currentTime.Nanosecond())
	fmt.Println("Location: ", currentTime.Location())
	fmt.Println("Weekday: ", currentTime.Weekday())
	fmt.Println("Unix: ", currentTime.Unix())
	fmt.Println("UnixNano: ", currentTime.UnixNano())
	fmt.Println("Timezone: ", currentTime.Location().String())

	fmt.Println("Formatted time: ", currentTime.Format("2006-01-02 15:04:05"))
	formatted := currentTime.Format("02-01-2006")
	fmt.Println("Formatted data: ", formatted)

	


}