package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingo struct {
	numbers []int
	boards  []*board
}

type board struct {
	// non zero indexed
	boardWidth, boardHeight int
	won                     bool
	grid                    []bool
	remainingNumbers        map[int]*coordinate
}

type coordinate struct {
	// start from top left at 0, 0
	x, y int
}

func newBingoFromFile(f *os.File) *bingo {
	var buff = bufio.NewScanner(f)
	var nums = getNumbers(buff)
	var boards = []*board{}
	for nextBoard := getNextBoard(buff); nextBoard != nil; {
		boards = append(boards, nextBoard)
		nextBoard = getNextBoard(buff)

	}
	return &bingo{nums, boards}
}

func getNumbers(buff *bufio.Scanner) []int {
	var nums = []int{}
	buff.Scan()
	var line = buff.Text()
	var rawNums = strings.Split(line, ",")
	buff.Scan() // consume next
	for _, v := range rawNums {
		var num, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Couldn't parse into number: %s. Error: %v", v, err)
		}
		nums = append(nums, num)
	}
	return nums
}

func getNextBoard(buff *bufio.Scanner) *board {
	var width, height = 0, 0
	var nums = []int{}
	for scanned := buff.Scan(); scanned; {
		if !scanned { // hit end of file
			break
		}
		var line = buff.Text()

		if line == "" || !scanned { // hit end of board
			break
		}
		var row = generateNumsFromLine(line)
		nums = append(nums, row...)
		width = len(row)
		height++
		scanned = buff.Scan()

	}
	if len(nums) == 0 {
		return nil
	}
	var grid = make([]bool, len(nums))
	var coords = createCoordinateMap(width, height, nums)
	return &board{width, height, false, grid, coords}
}

func generateNumsFromLine(line string) []int {
	var lineNums = strings.Fields(line)
	var nums = []int{}
	for _, v := range lineNums {
		var num, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Couldn't parse into number: %s. Error: %v", v, err)
		}
		nums = append(nums, num)
	}
	return nums
}

func createCoordinateMap(width, height int, nums []int) map[int]*coordinate {
	var coords = map[int]*coordinate{}
	var i = 0
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			var num = nums[i]
			coords[num] = &coordinate{column, row}
			i++
		}
	}
	return coords
}
