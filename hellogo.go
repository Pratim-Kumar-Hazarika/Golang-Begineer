package main

import (
	"fmt"
	"reflect"
)

//alias
var pl = fmt.Println

func main() {
	pl(reflect.TypeOf(42))
	pl(reflect.TypeOf(3.1415))
	pl(reflect.TypeOf(3 + 4i))
	pl(reflect.TypeOf("hello"))
}
	
