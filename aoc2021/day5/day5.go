package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var file, err = os.Open("day5.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v\n", err)
	}
	defer file.Close()
	solve(file)
}

type vector2 struct {
	x, y int
}

type ventLine struct {
	start, end *vector2
}

func (v *ventLine) isDiagonal() bool {
	var sameHorizontal = v.start.x == v.end.x
	var sameVertical = v.start.y == v.end.y
	return !sameHorizontal && !sameVertical
}

func createVentLine(line string) *ventLine {
	var x1, y1, x2, y2 int
	var _, err = fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
	if err != nil {
		log.Fatalf(" Error: %v. Failed to parse line: %s", err, line)
	}
	// if x2 < x1 {
	// 	var temp = x1
	// 	x1 = x2
	// 	x2 = temp
	// }
	// if y2 < y1 {
	// 	var temp = y1
	// 	y1 = y2
	// 	y2 = temp
	// }
	var start, end = &vector2{x1, y1}, &vector2{x2, y2}
	return &ventLine{start, end}
}

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)
	var ventLines = []*ventLine{}
	for buff.Scan() {
		var line = buff.Text()
		var ventLine = createVentLine(line)
		ventLines = append(ventLines, ventLine)
	}
	var countMap = map[vector2]int{}
	for _, v := range ventLines {
		// if !v.isDiagonal() {
		updateCountMapWithLine(v, countMap)
		// }
	}
	var dangerCount = countDangerCoords(countMap)
	fmt.Printf("Part 1 Solution: %d\n", dangerCount)
	fmt.Printf("Part 2 Solution: %d\n", 0)
}

func updateCountMapWithLine(vl *ventLine, counts map[vector2]int) {
	var xEnd, yEnd = false, false
	for x, y := vl.start.x, vl.start.y; ; {
		var v = vector2{x, y}
		fmt.Print()
		counts[v]++
		if xEnd && yEnd {
			break
		}

		if vl.start.x < vl.end.x {
			// increase x until finished
			if x < vl.end.x {
				x++
				if x == vl.end.x {
					xEnd = true
				}
			}
		} else if vl.start.x > vl.end.x {
			if x > vl.end.x {
				x--
				if x == vl.end.x {
					xEnd = true
				}
			}
		} else {
			xEnd = true
		}

		if vl.start.y < vl.end.y {
			// increase x until finished
			if y < vl.end.y {
				y++
				if y == vl.end.y {
					yEnd = true
				}
			}
		} else if vl.start.y > vl.end.y {
			if y > vl.end.y {
				y--
				if y == vl.end.y {
					yEnd = true
				}
			}
		} else {
			yEnd = true
		}

	}
}

func countDangerCoords(counts map[vector2]int) int {
	var count = 0
	for _, v := range counts {
		if v > 1 {
			count++
		}
	}
	return count
}
