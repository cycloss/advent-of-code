package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var file, err = os.Open("day9.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v\n", err)
	}
	defer file.Close()
	solve(file)
}

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)
	for buff.Scan() {
		var line = buff.Text()
		fmt.Println(line)
	}
	fmt.Printf("Part 1 Solution: %d\n", 0)
	fmt.Printf("Part 2 Solution: %d\n", 0)

}
