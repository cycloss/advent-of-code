package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type op func([]int) int

func main() {

	// create dictionary with output as keys and value as instruction
	// recursively calculate what key a is through its value by going back until you get a number
	// can only ever be one output which is unique

	const valToFind = "a"
	var opMap = generateOpMap()
	var val = find(valToFind, opMap)
	fmt.Printf("Part 1 value of %s: %d\n", valToFind, val)
	cache = map[string]int{"b": val}
	var val2 = find(valToFind, opMap)
	fmt.Printf("Part 2 value of %s: %d\n", valToFind, val2)

}

func generateOpMap() map[string]*opBundle {
	var opMap = map[string]*opBundle{}
	var file, _ = os.Open("inputs/day7.txt")
	defer file.Close()
	var reader = bufio.NewScanner(file)
	for reader.Scan() {
		var line = reader.Text()
		var tokens = strings.Split(line, " ")
		var key = tokens[len(tokens)-1]
		var opBundle = opBundleFromTokens(tokens)
		opMap[key] = opBundle
	}
	return opMap
}

var cache = map[string]int{}

func find(val string, opMap map[string]*opBundle) int {
	var bundle = opMap[val]
	var intVals = []int{}
	for _, param := range bundle.params {
		var intVal, err = strconv.Atoi(param)
		if err != nil {
			var cached, exists = cache[param]
			if exists {
				intVal = cached
			} else {
				intVal = find(param, opMap)
			}

		}
		cache[param] = intVal
		intVals = append(intVals, intVal)
	}
	return bundle.op(intVals)
}

type opBundle struct {
	params []string
	op     op
}

func opBundleFromTokens(tokens []string) *opBundle {
	switch tokens[1] {
	case "RSHIFT":
		return newBinaryOpBundle(tokens, rShift)
	case "LSHIFT":
		return newBinaryOpBundle(tokens, lShift)
	case "AND":
		return newBinaryOpBundle(tokens, and)
	case "OR":
		return newBinaryOpBundle(tokens, or)
	}
	if tokens[0] == "NOT" {
		return newUnaryOpBundle(tokens[1], not)
	}
	return newUnaryOpBundle(tokens[0], none)
}

func newBinaryOpBundle(tokens []string, op op) *opBundle {
	return &opBundle{[]string{tokens[0], tokens[2]}, op}
}

func newUnaryOpBundle(token string, op op) *opBundle {
	return &opBundle{[]string{token}, op}
}

func rShift(params []int) int {
	return params[0] >> params[1]
}

func lShift(params []int) int {
	return params[0] << params[1]
}

func or(params []int) int {
	return params[0] | params[1]
}

func and(params []int) int {
	return params[0] & params[1]
}

func not(params []int) int {
	return ^params[0]
}

func none(params []int) int {
	return params[0]
}
