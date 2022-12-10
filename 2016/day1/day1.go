package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/cycloss/advent-of-code/utils"
)

type vector struct {
	x, y int
}

func newVector(x, y int) vector {
	return vector{x: x, y: y}
}

// up, right, down, left
func makeCardinalVectors() []vector {
	return []vector{newVector(0, 1), newVector(1, 0), newVector(0, -1), newVector(-1, 0)}
}

func (v vector) add(v2 vector) vector {
	return newVector(v.x+v2.x, v.y+v2.y)
}

func (v vector) allIntermediaryTo(v2, cardinal vector) []vector {
	intermediary := []vector{}
	for ; v != v2; v = v.add(cardinal) {
		intermediary = append(intermediary, v)
	}
	return intermediary
}

type instruction struct {
	dirMod, dist int
}

func (i instruction) String() string {
	return fmt.Sprintf("dirMod: %d, dist: %d", i.dirMod, i.dist)
}

func newInstructionFromRaw(raw []byte) instruction {

	dirMod := func(firstChar byte) int {
		if firstChar == 'R' {
			return 1
		} else {
			return -1
		}
	}(raw[0])

	dist := func(raw []byte) int {
		numBytes := bytes.TrimSuffix(raw[1:], []byte{','})
		num, err := strconv.Atoi(string(numBytes))
		if err != nil {
			utils.ExitFatal("%v", err)
		}
		return num
	}(raw)

	return instruction{dirMod: dirMod, dist: dist}
}

func (i instruction) toVector(cardinal vector) vector {
	return vector{x: cardinal.x * i.dist, y: cardinal.y * i.dist}
}

func main() {
	cardinals := makeCardinalVectors()
	// 0 is north
	dirIndex := 0
	currentPos := vector{x: 0, y: 0}

	f, err := os.Open("./2016/day1/input.txt")
	if err != nil {
		utils.ExitFatal("%v", err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanWords)

	visited := map[vector]bool{}

	found := false

	for scan.Scan() {
		rawInstruction := scan.Bytes()
		instruction := newInstructionFromRaw(rawInstruction)
		dirIndex = mod(dirIndex+instruction.dirMod, 4)
		normalisedDirection := cardinals[dirIndex]
		walkVector := instruction.toVector(normalisedDirection)

		nextPos := currentPos.add(walkVector)
		intermediate := currentPos.allIntermediaryTo(nextPos, normalisedDirection)
		if !found {
		inner:
			for _, pos := range intermediate {
				if _, exists := visited[pos]; exists {
					fmt.Printf("%+v", pos)
					fmt.Printf("distance 2: %d\n", abs(pos.x)+abs(pos.y))
					found = true
					break inner
				} else {
					visited[pos] = true
				}
			}
		}
		currentPos = nextPos
	}

	fmt.Printf("distance 1: %d\n", abs(currentPos.x)+(currentPos.y))

}

func mod(a, b int) int {
	return (a%b + b) % b
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
