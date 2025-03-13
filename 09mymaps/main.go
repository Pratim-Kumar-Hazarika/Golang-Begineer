package main

import "fmt"

func main()  {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)
	languages["Js"] ="Javascript"
	languages["Rb"] ="Ruby"
	languages["Py"] ="Python"
	fmt.Println("List of all languages ", languages)

	fmt.Println("Js shorts for: ", languages["Js"])

	delete(languages,"Rb")
	fmt.Println("List of all languages ", languages)

	//loops are interesting in golang

	for key, value := range languages {
		// fmt.Println("For key %v, value is %v\n",key,value )
		fmt.Printf("For key %v, value is %v\n",key,value )
	}
}

// fmt.Println	 : Prints the entire string as-is, ignores format specifiers (%v).
// fmt.Printf    : Processes format specifiers (%v, %s, %d etc.) and replaces them with values.