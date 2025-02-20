package main

import (
	"fmt"
	"reflect"
)

//alias
var pl = fmt.Println

func main() {
	pl(reflect.TypeOf(42))
}
	
