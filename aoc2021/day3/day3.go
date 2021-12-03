package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("day3/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	solve(file)
}

const binLen = 12

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)
	var counts = make([]int, binLen, binLen)
	var binaryList = [][]byte{}
	for buff.Scan() {
		var line = buff.Text()
		var binary = []byte(line)
		updateCounts(binary, counts)
		binaryList = append(binaryList, binary)
	}
	var gamma = calculateGamma(counts)
	var epsilon = ^gamma & int(math.Pow(2, binLen)-1)
	fmt.Println(gamma, epsilon)
	fmt.Printf("Part 1 Solution: %d\n", gamma*epsilon)

	var oxygen = calculateSupportRating(binaryList, 0, true)
	var scrubber = calculateSupportRating(binaryList, 0, false)
	fmt.Printf("Part 2 Solution: %d\n", oxygen*scrubber)

}

func updateCounts(binary []byte, counts []int) {
	for i, v := range binary {
		if v == '1' {
			counts[i]++
		} else if v == '0' {
			counts[i]--
		} else {
			log.Fatalf("Invalid digit: %v\n", v)
		}
	}
}

func calculateGamma(counts []int) int {
	var binary = make([]byte, binLen, binLen)
	for i, v := range counts {
		if v < 0 {
			binary[i] = '0'
		} else if v > 1 {
			binary[i] = '1'
		} else {
			log.Fatalf("0 value detected in count")
		}
	}
	num, err := strconv.ParseInt(string(binary), 2, 64)
	if err != nil {
		log.Fatalf("Could not parse %s into an int\n", string(binary))
	}
	return int(num)
}

func calculateSupportRating(bl [][]byte, i int, oxygen bool) int {
	if len(bl) == 1 {
		num, err := strconv.ParseInt(string(bl[0]), 2, 64)
		if err != nil {
			log.Fatalf("Could not parse %s into an int\n", string(bl[0]))
		}
		return int(num)
	}
	var bit = getMostCommonBit(bl, i)
	var bl2 = reduceBinaryList(bl, i, bit, oxygen)
	return calculateSupportRating(bl2, i+1, oxygen)
}

func getMostCommonBit(bl [][]byte, i int) byte {
	var count = 0
	for _, v := range bl {
		var bit = v[i]
		if bit == '1' {
			count++
		} else {
			count--
		}
	}
	if count < 0 {
		return '0'
	} else if count > 0 {
		return '1'
	} else {
		return 'n'
	}
}

func reduceBinaryList(bl [][]byte, i int, mostCommonBit byte, oxygen bool) [][]byte {
	var bl2 = [][]byte{}
	for _, v := range bl {
		var bit = v[i]
		if oxygen {
			if bit == mostCommonBit {
				bl2 = append(bl2, v)
			} else if mostCommonBit == 'n' && bit == '1' {
				bl2 = append(bl2, v)
			}
		} else {
			if bit != mostCommonBit && mostCommonBit != 'n' {
				bl2 = append(bl2, v)
			} else if mostCommonBit == 'n' && bit == '0' {
				bl2 = append(bl2, v)
			}
		}
	}
	return bl2
}
