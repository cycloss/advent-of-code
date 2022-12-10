package aoc2015

import (
	"fmt"
	"io/ioutil"
)

func Day1() {

	var openP, closeP = byte('('), byte(')')

	var fileBytes, _ = ioutil.ReadFile("inputs/day1.txt")
	var total = 0

	for i, b := range fileBytes {
		if b == openP {
			total++
		} else if b == closeP {
			total--
		}
		if total == -1 {
			fmt.Printf("Entered basement on pos: %d\n", i+1)
		}
	}
	fmt.Println(total)

}
