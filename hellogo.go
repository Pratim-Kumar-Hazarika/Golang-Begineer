package main

import (
	"fmt"
)

//alias
var pl = fmt.Println

// func main() {
// 	pl(reflect.TypeOf(42))
// 	pl(reflect.TypeOf(3.1415))
// 	pl(reflect.TypeOf(3 + 4i))
// 	pl(reflect.TypeOf("hello"))
// 	pl(reflect.TypeOf('⏰'))
// }
	
//Casting
func main(){
	//Convert string
	// cv1 := "5000000000"
	//  cv2 , err := strconv.Atoi(cv1)
	//  pl(cv2, err, reflect.TypeOf(cv2))

	// cV1 := 5000000000
	// cv2 := strconv.Itoa(cV1)
	// pl(cv2,reflect.TypeOf(cv2))

	//Convert to float
	// cv1 := "3.14"
	// if cv1 , err := strconv.ParseFloat(cv1, 64);
	// err == nil{
	// 	pl(cv1, reflect.TypeOf(cv1));
	// }

	//Conditional operators

	iAge :=18
	if iAge >= 18{
		pl("You are an adult")
	}else{
		pl("You are a minor")
	}

	/// Little difference
	if(iAge >=18) && (iAge <65){
		pl("You are an adult of genz")
	}
}