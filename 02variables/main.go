package main

import "fmt"

//Capital L -> this is a public variable
const LoginToken string = "LOGIN___TOKEN"

func main(){
	var username string = "pratim"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool =true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal  uint8 =255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)



	var smallFloat  float64 =255.543444
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	///default ValUE AND SOME ALISAES

	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	//implicit type 
	var website = "www.prratim.com"
	fmt.Println(website)

	//no var style
	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
		fmt.Printf("Variable is of type: %T \n", LoginToken)
}