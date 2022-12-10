package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("day2/day2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	solve(file)
}

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)
	var x, y, y2 = 0, 0, 0
	for buff.Scan() {
		var line = buff.Text()
		var instruction string
		var num int
		fmt.Sscanf(line, "%s %d", &instruction, &num)
		handlePart1(instruction, num, &x, &y)
		handlePart2(instruction, num, &y, &y2)
	}
	fmt.Printf("Forward: %d, Depth: %d\n", x, y)
	fmt.Printf("Answer %d\n", x*y)
	fmt.Printf("Part 2 Forward: %d, Depth: %d\n", x, y2)
	fmt.Printf("Answer 2 %d\n", x*y2)

}

func handlePart1(instruction string, num int, x *int, y *int) {
	switch instruction {
	case "forward":
		*x += num
	case "down":
		*y += num
	case "up":
		*y -= num
	}
}

func handlePart2(instruction string, num int, aim *int, depth *int) {
	if instruction == "forward" {
		*depth += *aim * num
	}
}
