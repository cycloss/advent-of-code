package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var file, err = os.Open("day1/day1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	solve(file)
}

type ringBuffer struct {
	index  int
	cap    int
	buffer []*int
}

func newBuffer(cap int) *ringBuffer {
	var buff = make([]*int, cap, cap)
	return &ringBuffer{0, cap, buff}
}

func (b *ringBuffer) sum() int {
	var total = 0
	for _, v := range b.buffer {
		total += *v
	}
	return total
}

func (b *ringBuffer) getWriteIndex() int {
	return b.index % b.cap
}

func (b *ringBuffer) insertValue(val int) {
	b.buffer[b.getWriteIndex()] = &val
	b.index++
}

func (b *ringBuffer) full() bool {
	return b.buffer[b.getWriteIndex()] != nil
}

const windowSize1 = 1
const windowSize2 = 3

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)

	var ringBuff1 = newBuffer(windowSize1)
	var increaseCount1 = 0
	var increaseCount2 = 0
	var ringBuff2 = newBuffer(windowSize2)

	for buff.Scan() {
		var depth = buff.Text()
		var currentDepth, err = strconv.Atoi(depth)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert line: %s. Error: %v\n", depth, err))
		}
		if windowIncreased(ringBuff1, currentDepth) {
			increaseCount1++
		}
		if windowIncreased(ringBuff2, currentDepth) {
			increaseCount2++
		}
	}
	fmt.Printf("Sea bed increased: %d times\n", increaseCount1)
	fmt.Printf("Sea bed increased: %d times in sliding windows of %d\n", increaseCount2, windowSize2)

}

func windowIncreased(depthBuffer *ringBuffer, newVal int) bool {

	if depthBuffer.full() {
		var prevSum = depthBuffer.sum()
		depthBuffer.insertValue(newVal)
		var currentSum = depthBuffer.sum()
		return currentSum > prevSum
	} else {
		depthBuffer.insertValue(newVal)
		return false
	}
}
