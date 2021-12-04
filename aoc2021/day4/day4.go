package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var fileName = "day4.txt"
	fmt.Printf("Part 1 Solution: %d\n", solvePart1(fileName))
	fmt.Printf("Part 2 Solution: %d\n", solvePart2(fileName))
}

func solvePart1(file string) int {
	var b = createBingoFromFile(file)
	for _, v := range b.numbers {
		var winner = markAllBoardsWithNumber(v, b.boards)
		if winner != nil {
			return calculateWinningScore(winner, v)
		}
	}
	log.Fatal("no solution could be found")
	return -1
}

func createBingoFromFile(fileName string) *bingo {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %v\n", err)
		return nil
	}
	defer file.Close()
	return newBingoFromFile(file)
}

// returns last winning board found in round, otherwise nil
func markAllBoardsWithNumber(num int, boards []*board) *board {
	var winner *board
	for _, b := range boards {
		if b.won {
			continue
		}
		var coord = b.remainingNumbers[num]
		if coord != nil {
			b.markCoordinate(coord)

			delete(b.remainingNumbers, num)
			if b.winsWithCoordinate(coord) {
				winner = b
			}
		}
	}
	return winner
}

func (b *board) markCoordinate(c *coordinate) {
	// second row (yi 1) with 3w board with gives rowI 3 which is start of row y1
	var rowI = c.y * b.boardWidth
	// second column (xi 1) adds to rowI 3 to give to of 4 which is middle of 3 x 3
	var columnI = c.x
	var translatedIndex = rowI + columnI
	b.grid[translatedIndex] = true
}

func (b *board) winsWithCoordinate(c *coordinate) bool {
	if b.won {
		return false
	}
	if b.rowWinsWithCoordinate(c) || b.columnWinsWithCoordinate(c) {
		b.won = true
		return true
	} else {
		return false
	}
}

func (b *board) rowWinsWithCoordinate(c *coordinate) bool {
	// check row of coord, fix row and change column
	for x2 := 0; x2 < b.boardWidth; x2++ {
		var cToTest = &coordinate{x2, c.y}
		if !b.checkCoordinate(cToTest) {
			return false
		}
	}
	return true
}

func (b *board) checkCoordinate(c *coordinate) bool {
	// second row (yi 1) with 3w board with gives rowI 3 which is start of row y1
	var rowI = c.y * b.boardWidth
	// second column (xi 1) adds to rowI 3 to give to of 4 which is middle of 3 x 3
	var columnI = c.x
	var translatedIndex = rowI + columnI
	return b.grid[translatedIndex]
}

func (b *board) columnWinsWithCoordinate(c *coordinate) bool {
	// check column of coord, fix column  and change row
	for y2 := 0; y2 < b.boardHeight; y2++ {
		var cToTest = &coordinate{c.x, y2}
		if !b.checkCoordinate(cToTest) {
			return false
		}
	}
	return true
}

func calculateWinningScore(b *board, num int) int {
	var remainingSum = 0
	for k := range b.remainingNumbers {
		remainingSum += k
	}
	return remainingSum * num
}

func solvePart2(file string) int {
	var b = createBingoFromFile(file)
	var lastWinner *board
	var lastNumber int
	for _, v := range b.numbers {
		var winner = markAllBoardsWithNumber(v, b.boards)
		if winner != nil {
			lastWinner = winner
			lastNumber = v
		}
	}
	if lastWinner == nil {
		log.Fatal("no solution could be found")
		return -1
	} else {
		return calculateWinningScore(lastWinner, lastNumber)
	}
}
