package main

import (
	"fmt"
	"net/http"
	"sync"
)
var signals = []string{"test"}
var wg sync.WaitGroup //pointers
var mut sync.Mutex //pointer
func  main()  {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://prratim.com",
	}

	for _, web := range websitelist{
		wg.Add(1)
		 getStatusCode(web)
	}
	wg.Wait() //doesnot allow main method to finish
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res,err := http.Get(endpoint)
	if err != nil{
		fmt.Println("OOps in endpoint")
	}else{
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status dode for %s\n",res.StatusCode, endpoint)
	}
	defer res.Body.Close() // Close response body

	
}