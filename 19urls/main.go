package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://localhost:3000/learn?coursename=golang&paymentId=eifhdjd"
func  main()  {
	fmt.Println("Welcome to urls in golang")
	
	fmt.Println(myurl)

	//parsing

	result , err := url.Parse(myurl)
	if err != nil{
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentId"])

	for _,value := range qparams{
		fmt.Println("Param is : ",value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
	    Host: "prratim.com",
		Path: "/hire",
		RawQuery: "user=pratim",
	}
	anotherUrl :=partsOfUrl.String()
	fmt.Println(anotherUrl)
}