package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/cycloss/advent-of-code/utils"
)

const target = 2020

func main() {
	f, err := os.Open("aoc2020/day1/day1.txt")
	if err != nil {
		utils.ExitFatal("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	nums := make(map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		nums[num] = true
	}

	for num := range nums {
		toFind := target - num
		if _, exists := nums[toFind]; exists {
			fmt.Printf("%d * %d = %d\n", num, toFind, num*toFind)
			os.Exit(0)
		}
	}
	fmt.Println("num not found")

}
