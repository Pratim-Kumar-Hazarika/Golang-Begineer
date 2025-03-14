package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to files in golang....")

	content := "This needs to go in a file - Prratim.com"
	file , err := os.Create("./prratim.text")

	checkNilError(err)

	length, err := io.WriteString(file,content)

	checkNilError(err)
	
	fmt.Println("Length is: " , length)
	defer file.Close()
	readFile("./prratim.text")
	
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	
	fmt.Println("Text datra inside the file\n",string(databyte))
}

func checkNilError (err error){
	if err != nil{
		panic(err)
	}
}