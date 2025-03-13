package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")

	presentTime := time.Now()
	fmt.Println("Current Time:", presentTime)

	// Correcting the format string
	fmt.Println("Formatted Time:", presentTime.Format("01-02-2006 15:04 Monday"))

	createdDate := time.Date(2020 , time.October , 10, 23,23,0,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
