package main

import (
	"fmt"
	"internal/cpu"
)

type Person struct {
	name string
	age  int
}

func main() {
	var m = map[Person]bool{}
	var p1 = Person{"ted", 28}
	var p2 = Person{"hugo", 27}
	m[p1] = true
	m[p2] = true
	fmt.Println(cpu.Name())
	var p3 = Person{"ted", 28}

	if m[p3] {
		fmt.Println("Already exists")
	}
}
