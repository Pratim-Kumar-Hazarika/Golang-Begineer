package main

import "fmt"

func main()  {

	fmt.Println("Welcome to my functions in golang ")
	greeter()
	greeterTwo()
	result := adder(3,4)

	fmt.Println("Result is : ",result)

	proResult, myMessage := proAdder(2,2,4,4,44,4,)
	fmt.Println("Pro-result is : ",proResult)
	fmt.Println("Pro-message is : ",myMessage)
}

func adder(valOne int , valTwo int) int{
		return valOne +valTwo
}

func proAdder(values ...int)(int,string){
	total := 0

	for _,val :=range values {
		total += val
	}
	return total, "Hi pro result hehe"
}

func greeterTwo()  {
	fmt.Println("Another method")
}

func greeter(){
	fmt.Println("Hello from golang")
}