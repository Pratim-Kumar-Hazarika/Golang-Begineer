package main

import "fmt"

func main()  {
	fmt.Println("Welcome to my arrays in golang")
	// Fruit list is  [Apple Tomato  Peach]
	var fruitList [4] string
	fruitList[0] ="Apple"
	fruitList[1] ="Tomato"
	fruitList[3] ="Peach"



	fmt.Println("Fruit list is ",fruitList)
	fmt.Println("Fruit list is ",len(fruitList))

	var vegList =[5]string{"porato","beans","mushroom"}
	fmt.Println("Veg  list is ",vegList)
	fmt.Println("Veg list is ",len(vegList))

}