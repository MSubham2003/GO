package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Printf("Type of currTime is %T\n", currTime)


	//02-01-2006,15:04:05, Monday is fixed cuz on this day go launched and it is in (dd-mm-yyyy,hh:mm:ss, day) format
	formattedTime := time.Now().Format("02-01-2006,15:04:05, Monday")
	fmt.Println(formattedTime)


	formattedTime2 := time.Now().Format("02-01-2006, 3:04 PM")
	fmt.Println(formattedTime2)


	//String to Time
	layout := "2006-01-02"
	str := "2021-07-01"
	t, _ := time.Parse(layout, str)
	fmt.Println(t)

	//Time to String
	layout2 := "2006-01-02"
	t2 := time.Now()
	str2 := t2.Format(layout2)
	fmt.Println(str2)

	//add 1 more day in current time
	new_date := currTime.Add(24*time.Hour)
	fmt.Println(new_date.Format(layout))
}
