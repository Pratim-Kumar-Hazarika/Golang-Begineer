package main

import "fmt"

func main(){
	fmt.Println("Welcome to Structs in golang!!")
	//no inheritance in golang ; No super or parent
	pratim := User{"Pratim","hpro@gmal.com",true,25}
	fmt.Println(pratim)
	fmt.Printf("Pratim details are %+v\n",pratim )
	fmt.Printf("Name is %v and %v email is", pratim.Name, pratim.Email)
	
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}