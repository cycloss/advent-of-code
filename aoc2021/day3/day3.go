package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const binLen = 12

func main() {
	file, err := os.Open("day3/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var nums = convertFileToInts(file)
	fmt.Printf("Part 1 Solution: %d\n", solvePart1(nums))
	fmt.Printf("Part 2 Solution: %d\n", solvePart2(nums))
}

func convertFileToInts(file *os.File) []int {

	var buff = bufio.NewScanner(file)
	var nums = []int{}
	for buff.Scan() {
		var line = buff.Text()
		var num, _ = strconv.ParseInt(line, 2, 64)
		nums = append(nums, int(num))
	}
	return nums
}

func solvePart1(nums []int) int {
	var maskBit = 1
	var gamma = 0
	for i := 0; i < binLen; i++ {
		gamma |= extractMostCommonBit(nums, maskBit)
		maskBit <<= 1
	}
	var epsilon = ^gamma & int(math.Pow(2, binLen)-1)

	return gamma * epsilon
}

func extractMostCommonBit(nums []int, maskBit int) int {
	var count = 0
	for _, v := range nums {
		var bit = v & maskBit
		if bit > 0 {
			count++
		} else {
			count--
		}
	}
	if count > 0 {
		return maskBit
	} else if count < 0 {
		return 0
	} else {
		return -1
	}
}

func solvePart2(nums []int) int {
	var oxygenRating = reduceNumbers(nums, 1<<(binLen-1), oxygenFilter)
	var scrubberRating = reduceNumbers(nums, 1<<(binLen-1), scrubberFilter)
	return oxygenRating * scrubberRating
}

/// returns true if the filter is passed
type filterFunc func(int, int, int) bool

func reduceNumbers(nums []int, mask int, filter filterFunc) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var commonBit = extractMostCommonBit(nums, mask)
	var nums2 = []int{}
	for _, v := range nums {
		if filter(v, commonBit, mask) {
			nums2 = append(nums2, v)
		}
	}
	return reduceNumbers(nums2, mask>>1, filter)
}

func oxygenFilter(num int, commonBit int, mask int) bool {
	if commonBit > 0 {
		return (mask & num) != 0
	} else if commonBit == 0 {
		return (mask & num) == 0
	} else {
		return (mask & num) != 0
	}
}

func scrubberFilter(num int, commonBit int, mask int) bool {
	if commonBit > 0 {
		return (mask & num) == 0
	} else if commonBit == 0 {
		return (mask & num) != 0
	} else {
		return (mask & num) == 0
	}
}
