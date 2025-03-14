package main

import "fmt"

func main()  {
	fmt.Println("Welcome to loops in golang")

	days := []string {"Sunday","Tuesday","Wednesday","Thursday","Friday"}
	fmt.Println(days)


	// for d:=0; d< len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	
	for index , day := range days {
		fmt.Printf("index is %v and value is %v\n", index,day)
	}

	rougeValue :=1

	for rougeValue <10 {
		if rougeValue ==5 {
			goto loco
		}
		// if rougeValue ==5{
		// 	break;
		// }
		fmt.Println("Value is :", rougeValue)
		rougeValue++
	}
	loco:
		fmt.Println("Jumping to prratim.com")
}