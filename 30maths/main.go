package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int =2
	// var mynumbeTwo float64 = 4.5

	// fmt.Println("The sum is : ", mynumberOne + mynumbeTwo) 

	//random nunber
	fmt.Println(time.Now().UnixNano())
	fmt.Println(rand.Intn(5))
	
}