package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Get requests in golang...")
	PeformGetRequest()
}

func PeformGetRequest(){
	const myurl ="http://localhost:8080/get"
	response ,err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Staus code: ", response.StatusCode)
	fmt.Println("Content length ", response.ContentLength)

	var responseString strings.Builder
	content,_ := io.ReadAll(response.Body) ////Byte format
	byteCount,_ := responseString.Write(content)
	fmt.Println("Byte coubnt is: ",byteCount)
	fmt.Println(responseString.String())
	

	// fmt.Println(string(content))
}