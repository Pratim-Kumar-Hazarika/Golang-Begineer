package main

import "fmt"

func main(){
	fmt.Println("Welcome to methods in golang!!")
	//no inheritance in golang ; No super or parent
	pratim := User{"Pratim","hpro@gmal.com",true,25}
	fmt.Println(pratim)
	fmt.Printf("Pratim details are %+v\n",pratim )
	fmt.Printf("Name is %v and %v email is\n", pratim.Name, pratim.Email)

	pratim.GetStatus()
	pratim.NewMail()

	fmt.Printf("Name is %v and %v email is\n", pratim.Name, pratim.Email)
	  
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
	
}
func (u User) GetStatus(){
		fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email ="test@go.com"
	fmt.Println("Email of this user is ", u.Email)
}

//Pass it as *User to change email : Pointer concept
// func (u User) NewMail(){ 
// 	u.Email ="test@go.com"
// 	fmt.Println("Email of this user is ", u.Email)
// }