package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("Welcome to channels in golang")
	
	myCh := make(chan int ,2)
	wg := &sync.WaitGroup{}


	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	//receive only
	go func (ch <-chan int , wg *sync.WaitGroup)  {
		close(myCh)
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)


	//send only
	go func (ch chan<- int , wg *sync.WaitGroup)  {
		myCh <- 0
		close(myCh)
		
		
		
		// myCh <- 6

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}