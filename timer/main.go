package main

import (
	"time"
	"fmt"
	"strconv"
	c "github.com/gabeart10/go-1/colors"
)

func main() {
	msec := 0
	sec := 0
	min := 0
	hour := 0
	day := 0
	week := 0
	for {
		fmt.Println("Millisecond: " + strconv.Itoa(msec) + "\n")
		fmt.Println("Second: " + strconv.Itoa(sec) + "\n")
		fmt.Println("Minute: " + strconv.Itoa(min) + "\n")
		fmt.Println("Hour: " + strconv.Itoa(hour) + "\n")
		fmt.Println("Day: " + strconv.Itoa(day) + "\n")
		fmt.Println("Week: " + strconv.Itoa(week))
		msec += 1
		time.Sleep(time.Millisecond)
		if msec == 1000 {
			sec += 1
			msec -= 1000
		}
		fmt.Print(c.Clear)
	}
}	
	
