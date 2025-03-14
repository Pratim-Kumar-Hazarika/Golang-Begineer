package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Post requests in golang...")
	PerfromPostRequest()
}

func PerfromPostRequest(){
	const myurl = "http://localhost:8080/post"
	requestBody := strings.NewReader(`
		{
			"coursename":"Le'ts go with golang",
			"price":0,
			"platform":"prrati.com"
		}
	`)
	response , err := http.Post(myurl,"application/json",requestBody)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}