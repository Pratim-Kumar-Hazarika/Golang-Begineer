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
	cv1 := 1.5
	cv2 := int(cv1)
	pl(cv2)
}