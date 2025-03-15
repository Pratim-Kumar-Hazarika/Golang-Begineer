package main

import (
	"fmt"
	"log"
	"mymongodb/router"
	"net/http"

	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Welcome to mongodb  api!!")
	r := router.Router()
	fmt.Println("Server is getting started....")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listenig at port 4000...")
	
}