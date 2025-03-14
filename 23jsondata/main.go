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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
                "coursename": "Reactjs",
                "price": 299,
                "website": "prratimc.comn",
                "tags": ["web-dev","js"]
        }
	`)
	// var bestCourse course 

	// checkValid := json.Valid(jsonDataFromWeb)
	// if checkValid {
	// 	fmt.Println("Json was valid")
	// 	json.Unmarshal(jsonDataFromWeb ,&bestCourse)
	// 	fmt.Printf("%#v\n", bestCourse)
	// }else{
	// 	fmt.Println("JSON WAS NOT VALID")
	// }
	/// some cases where you just want to add data to key value

	
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k,v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k,v,v)
	}

}