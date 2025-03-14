package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main()  {
	fmt.Println("Post requests in golang...")
	PerfromPosFormRequest()
}

func PerfromPosFormRequest(){
	const myurl = "http://localhost:8080/postform"
	
	///form data
	data := url.Values{}
	data.Add("firstname","pratim")
	data.Add("lastname","hazarika")

	response , err := http.PostForm(myurl,data)
	if err!= nil{
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}