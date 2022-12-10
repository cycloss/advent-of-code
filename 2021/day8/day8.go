package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type knownsMap map[int]map[byte]bool

func main() {
	var file, err = os.Open("day8.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v\n", err)
	}
	defer file.Close()
	solve(file)
}

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)
	var easyDigitsCount = 0
	var outputSum = 0
	for buff.Scan() {
		var line = buff.Text()
		easyDigitsCount += countEasyDigits(line)
		outputSum += sumOutput(line)
	}
	fmt.Printf("Part 1 Solution: %d\n", easyDigitsCount)
	fmt.Printf("Part 2 Solution: %d\n", outputSum)

}

var uniqueSegemntSet = map[int]bool{2: true, 4: true, 3: true, 7: true}

func countEasyDigits(line string) int {
	var count = 0
	var split = strings.Split(line, " ")
	var found = false
	for _, v := range split {
		if v == "|" {
			found = true
			continue
		}
		if found {
			_, contains := uniqueSegemntSet[len(v)]
			if contains {
				count++
			}
		}
	}
	return count
}

func sumOutput(line string) int {
	var nums, output = getNumsAndOutput(line)
	var knowns = createMapOfKnowns(nums)
	var fiveSegs = createNSegMap(nums, 5)
	var sixSegs = createNSegMap(nums, 6)
	var segMap = map[byte]byte{}
	var topSeg = findTopSeg(knowns)
	segMap[topSeg] = 'a'
	var bottomSeg = findBottomSeg(sixSegs, knowns[4], topSeg)
	segMap[bottomSeg] = 'g'
	var bottomLeft = findBottomLeftSeg(knowns, bottomSeg)
	segMap[bottomLeft] = 'e'
	var topRight = findTopRight(knowns[1], sixSegs)
	segMap[topRight] = 'c'
	var bottomRight = findBottomRight(knowns[1], topRight)
	segMap[bottomRight] = 'f'
	var middle = findMiddle(fiveSegs, segMap)
	segMap[middle] = 'd'
	var topLeft = findTopLeft(knowns[8], segMap)
	segMap[topLeft] = 'b'
	// for k, v := range segMap {
	// 	fmt.Printf("key: %s, val: %s\n", string(k), string(v))
	// }
	return calculateOutputSum(output, segMap)
}

func getNumsAndOutput(line string) ([]string, []string) {
	var split = strings.Split(line, " ")
	var nums = []string{}
	var output = []string{}
	var found = false
	for _, v := range split {
		if v == "|" {
			found = true
			continue
		}
		if found {
			output = append(output, v)
		} else {
			nums = append(nums, v)
		}
	}
	return nums, output
}

func createMapOfKnowns(nums []string) knownsMap {
	var knowns = knownsMap{}
	for _, v := range nums {
		switch len(v) {
		case 2:
			knowns[1] = createSetFromNum(v)
		case 3:
			knowns[7] = createSetFromNum(v)
		case 4:
			knowns[4] = createSetFromNum(v)
		case 7:
			knowns[8] = createSetFromNum(v)
		}
	}
	return knowns
}

func createNSegMap(nums []string, n int) []map[byte]bool {
	var segMap = []map[byte]bool{}
	for _, v := range nums {
		if len(v) == n {
			segMap = append(segMap, createSetFromNum(v))
		}
	}
	return segMap
}

func createSetFromNum(num string) map[byte]bool {
	var set = map[byte]bool{}
	for _, v := range []byte(num) {
		set[v] = true
	}
	return set
}

func findTopSeg(knowns knownsMap) byte {
	var one = knowns[1]
	var seven = knowns[7]
	var diff = diff(one, seven)
	return takeOne(diff)
}

func findBottomSeg(sixSegs []map[byte]bool, four map[byte]bool, top byte) byte {
	var inter = sixSegs[0]
	for _, k := range sixSegs {
		inter = intersection(inter, k)
	}
	var not4 = map[byte]bool{}
	for k := range inter {
		var _, found = four[k]
		if !found {
			not4[k] = true
		}
	}
	return takeNot(not4, top)
}

func findBottomLeftSeg(knowns knownsMap, bottom byte) byte {
	var eight = knowns[8]
	var merged = union(knowns[1], knowns[4])
	merged = union(merged, knowns[7])
	var diff = diff(eight, merged)
	return takeNot(diff, bottom)
}

func findTopRight(one map[byte]bool, sixSegs []map[byte]bool) byte {
	var foundCount = map[byte]int{}
	for _, seg := range sixSegs {
		for k := range seg {
			foundCount[k]++
		}
	}
	var twoSet = map[byte]bool{}
	for k, v := range foundCount {
		if v == 2 {
			twoSet[k] = true
		}
	}
	return takeOne(intersection(one, twoSet))
}

func findBottomRight(one map[byte]bool, topRight byte) byte {
	return takeNot(one, topRight)
}

func findMiddle(fiveSegs []map[byte]bool, segMap map[byte]byte) byte {
	var inter = fiveSegs[0]
	for _, v := range fiveSegs {
		inter = intersection(inter, v)
	}
	var b byte
	for k := range inter {
		var _, found = segMap[k]
		if !found {
			b = k
		}
	}
	return b
}

func findTopLeft(eight map[byte]bool, segMap map[byte]byte) byte {
	var b byte
	for k := range eight {
		var _, found = segMap[k]
		if !found {
			b = k
		}
	}
	return b
}

var sets = map[byte]map[byte]bool{'0': {'a': true, 'b': true, 'c': true, 'e': true, 'f': true, 'g': true}, '1': {'c': true, 'f': true}, '2': {'a': true, 'c': true, 'd': true, 'e': true, 'g': true}, '3': {'a': true, 'c': true, 'd': true, 'f': true, 'g': true}, '4': {'b': true, 'c': true, 'd': true, 'f': true}, '5': {'a': true, 'b': true, 'd': true, 'f': true, 'g': true}, '6': {'a': true, 'b': true, 'd': true, 'e': true, 'f': true, 'g': true}, '7': {'a': true, 'c': true, 'f': true}, '8': {'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true}, '9': {'a': true, 'b': true, 'c': true, 'd': true, 'f': true, 'g': true}}

func calculateOutputSum(output []string, segMap map[byte]byte) int {

	var bytes = []byte{}
	var translated = translateOutput(output, segMap)
	for _, transMap := range translated {
		for k, v := range sets {
			if sameSet(v, transMap) {
				bytes = append(bytes, k)
			}
		}
	}

	var res, err = strconv.Atoi(string(bytes))
	if err != nil {
		log.Fatal(err, "output", output, "segmap", segMap, "trans", translated)
	}
	return res
}

func translateOutput(output []string, segMap map[byte]byte) []map[byte]bool {
	var translated = []map[byte]bool{}
	for _, v := range output {
		var bytes = []byte(v)
		var set = map[byte]bool{}
		for _, b := range bytes {
			var translation = segMap[b]
			set[translation] = true
		}
		translated = append(translated, set)
	}
	return translated
}

func diff(set1, set2 map[byte]bool) map[byte]bool {
	var diff = map[byte]bool{}
	for k := range set1 {
		var _, found = set2[k]
		if !found {
			diff[k] = true
		}
	}
	for k := range set2 {
		var _, found = set1[k]
		if !found {
			diff[k] = true
		}
	}
	return diff
}

func intersection(set1, set2 map[byte]bool) map[byte]bool {
	var intersection = map[byte]bool{}
	for k := range set1 {
		var _, found = set2[k]
		if found {
			intersection[k] = true
		}
	}
	for k := range set2 {
		var _, found = set1[k]
		if found {
			intersection[k] = true
		}
	}
	return intersection
}

func union(set1, set2 map[byte]bool) map[byte]bool {
	var union = map[byte]bool{}
	for k := range set1 {
		union[k] = true
	}
	for k := range set2 {
		union[k] = true
	}
	return union
}

func sameSet(set1, set2 map[byte]bool) bool {
	for k := range set1 {
		var _, found = set2[k]
		if !found {
			return false
		}
	}
	for k := range set2 {
		var _, found = set1[k]
		if !found {
			return false
		}
	}
	return true
}

func takeNot(set map[byte]bool, b byte) byte {
	var notTop byte
	for k := range set {
		if k != b {
			notTop = k
		}
	}
	return notTop
}

func takeOne(set map[byte]bool) byte {
	var b byte
	for k := range set {
		b = k
	}
	return b
}
