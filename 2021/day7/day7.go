package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var file, err = os.Open("day7.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v\n", err)
	}
	defer file.Close()
	solve(file)
}

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)
	var crabs = getCrabs(buff)
	var medianI = getMedianI(crabs)
	// 184 is best for part 1
	var fuel = calculateFuelUsage(crabs, medianI-34)
	fmt.Printf("Part 1 Solution: %d\n", fuel)
	fmt.Printf("Part 2 Solution: %d\n", 0)

}

func getCrabs(buff *bufio.Scanner) []int {
	var nums = []int{}
	for buff.Scan() {
		var line = buff.Text()
		var numsRaw = strings.Split(line, ",")
		for _, v := range numsRaw {
			var num, err = strconv.Atoi(v)
			if err != nil {
				log.Fatalf("%v. %s", err, v)
			}
			nums = append(nums, num)
		}
	}
	return nums
}

func getMedianI(crabs []int) int {
	sort.Ints(crabs)
	return len(crabs) / 2
}

func calculateFuelUsage(crabs []int, alignI int) int {
	var total = 0
	for _, v := range crabs {
		var diff = absDiffInt(v, alignI)
		// triangle number formula
		total += diff * (diff + 1) / 2
	}
	return total
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
