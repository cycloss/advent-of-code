package main

import "fmt"

func main() {
	var m = map[string]bool{"a": true}
	var p = m["b"]
	fmt.Printf("%v", p)
}
