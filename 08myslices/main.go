package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to slices class in golang")
	// var fruitList =[]string{"Apple","Tomato","Peach"}

	// fmt.Println("Values", fruitList)

	// fmt.Printf("Type of fruitlis is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango","Banana")

	// fmt.Println("Added", fruitList)

	// fruitList = append(fruitList[1:3]) //works as slice
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:3]) //works as slice
	// fmt.Println(fruitList)

	// highScores  := make([]int,4)
	// highScores[0] =231
	// highScores[1] =954
	// highScores[2] =241
	// highScores[3] =233
	// // highScores[4] =777 does not work crash

	// highScores = append(highScores, 55,535,222) //realocates memory
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted((highScores)))

	// sort.Ints(highScores) //available in slices not in array ( methods)

	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted((highScores)))


	//how to remove a value from slices based on index

	var courses = []string{"reactjs","js","nodejs","swift","python","ruby"}
	// fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)
}