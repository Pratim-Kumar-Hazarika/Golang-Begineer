package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name     string `json:"coursename"`
	Price    int     `json:"price"`
	Platform string    `json:"website"`
	Password string    `json:"-"`
	Tags     []string   `json:"tags,omitempty"`
}
func main()  {
	fmt.Println("Welcome to create a Json Data in golang")
	EncodeJson()
}
func EncodeJson(){
	pratimCourses := []course{
		{"Reactjs",299,"prratimc.comn","abc123", []string{"web-dev","js"}},
		{"Nodejs",199,"prratimc.comn","vcc123", []string{"full-stack"}},
		{"Js",299,"prratimc.comn","addbc123", nil},
	}
	///package this data as JSON data

	finalJson,err := json.MarshalIndent(pratimCourses,"","\t")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(finalJson))
}