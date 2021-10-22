package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var file, _ = os.Open("inputs/day5.txt")
	var buff = bufio.NewScanner(file)
	defer file.Close()

	var niceV1Strings = 0
	var niceV2Strings = 0

	for buff.Scan() {
		var line = buff.Bytes()
		if checkLineV1Nice(line) {
			niceV1Strings++
		}
		if checkLineV2Nice(line) {
			niceV2Strings++
		}

	}

	fmt.Printf("Nice V1 strings: %d\n", niceV1Strings)
	fmt.Printf("Nice V2 strings: %d\n", niceV2Strings)

}

var vowels = map[byte]bool{byte('a'): true, byte('e'): true, byte('i'): true, byte('o'): true, byte('u'): true}

func checkLineV1Nice(line []byte) bool {
	var vowelCount = 0
	var sameCombo = 1
	var comboFound = false
	var previousLetter byte
	for i, currentLetter := range line {
		// combo check
		if i > 0 {
			if previousLetter == currentLetter {
				sameCombo++
			} else {
				sameCombo = 1
			}
			if sameCombo == 2 {
				comboFound = true
			}
		}
		previousLetter = currentLetter

		// vowels check
		if vowels[currentLetter] {
			vowelCount++
		}

		// danger combo check
		if isDangerCombo(line, i) {
			return false
		}

	}
	return vowelCount >= 3 && comboFound
}

var dangerCombos = map[byte][]byte{byte('a'): []byte("ab"), byte('c'): []byte("cd"), byte('p'): []byte("pq"), byte('x'): []byte("xy")}

func isDangerCombo(line []byte, i int) bool {
	var combo = dangerCombos[line[i]]
	if combo == nil {
		return false
	}

	for j := 0; j < len(combo); j++ {
		if i+j >= len(line) {
			return false
		}

		var comboByte = combo[j]
		var currentByte = line[i+j]
		if currentByte != comboByte {
			return false
		}
	}
	return true
}

func checkLineV2Nice(line []byte) bool {
	var pairDict = map[string]int{}
	var lastPair string = ""
	var mirrorFound = false
	var twoPairFound = false
	for i := range line {
		// check either side is same
		if i > 0 && i < len(line)-1 {
			var prev = line[i-1]
			var next = line[i+1]
			if prev == next {
				mirrorFound = true
			}
		}
		// exclude first
		if i > 0 {
			var pair = string(line[i-1 : i+1])
			if pair != lastPair {
				pairDict[pair]++
			}
			lastPair = pair
			if pairDict[pair] > 1 {
				twoPairFound = true
			}
		}
	}
	return mirrorFound && twoPairFound
}
