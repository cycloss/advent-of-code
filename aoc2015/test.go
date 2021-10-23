package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n, _ = strconv.Atoi("100,000")
	fmt.Println(n)
}
