package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i = 1
	fmt.Println(^(i << 63))
	fmt.Println(reflect.TypeOf(i).Size())
}
