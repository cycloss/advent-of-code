package aoc2015

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

var intMax = getIntMax()

func Day2() {
	var file, _ = os.Open("inputs/day2.txt")
	var buff = bufio.NewScanner(file)
	defer file.Close()

	var totalArea = 0
	var totalRibbon = 0
	for buff.Scan() {

		var lwh [3]int
		fmt.Sscanf(buff.Text(), "%dx%dx%d", &lwh[0], &lwh[1], &lwh[2])
		totalArea += calculateRequiredPaper(lwh)
		totalRibbon += calculateRequiredRibbon(lwh)
	}
	fmt.Printf("Total sq ft: %d\n", totalArea)
	fmt.Printf("Total ft of ribbon: %d\n", totalRibbon)
}

func calculateRequiredPaper(lwh [3]int) int {
	var area = calculateArea(lwh)
	var smallest, secondSmallest = findSmallestSides(lwh)
	return area + smallest*secondSmallest
}

func findSmallestSides(lwh [3]int) (int, int) {
	var smallest = intMax
	var secondSmallest = intMax
	for _, v := range lwh {
		if v < smallest {
			secondSmallest = smallest
			smallest = v
		} else if v < secondSmallest {
			secondSmallest = v
		}
	}
	return smallest, secondSmallest
}

func calculateArea(lwh [3]int) int {
	return 2*lwh[0]*lwh[1] + 2*lwh[1]*lwh[2] + 2*lwh[0]*lwh[2]
}

func calculateRequiredRibbon(lwh [3]int) int {
	var smallest, secondSmallest = findSmallestSides(lwh)
	var around = smallest*2 + secondSmallest*2

	return around + calculateVolume(lwh)
}

func calculateVolume(lwh [3]int) int {
	var vol = lwh[0]
	for i := 1; i < len(lwh); i++ {
		vol *= lwh[i]
	}
	return vol
}

// finds max integer on 32 or 64 bit runtime environments
func getIntMax() int {
	var i = 1
	var size = reflect.TypeOf(i).Size() * 8
	if size > 32 {
		return ^(i << 63)
	} else {
		return ^(i << 31)
	}
}
