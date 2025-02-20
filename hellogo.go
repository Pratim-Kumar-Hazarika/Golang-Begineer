package main

import (
	"fmt"
	"reflect"
	"strconv"
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
	cv1 := "5000000000"
	 cv2 , err := strconv.Atoi(cv1)
	 pl(cv2, err, reflect.TypeOf(cv2))
}