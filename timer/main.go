package main

import (
	"time"
	"fmt"
	//c "github.com/gabeart10/go-1/colors"
)

func main() {
	msec := 0
	sec := 0
	min := 0
	hour := 0
	day := 0
	week := 0
	for {
		fmt.Print("Millisecond: " + string(msec) + "\n")
		fmt.Println("Second: " + string(sec) + "\n")
		fmt.Println("Minute: " + string(min) + "\n")
		fmt.Println("Hour: " + string(hour) + "\n")
		fmt.Println("Day: " + string(day) + "\n")
		fmt.Println("Week: " + string(week))
		msec += 1
		time.Sleep(0.001)
	}
}	
	
