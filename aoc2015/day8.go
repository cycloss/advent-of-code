package main

import (
	"fmt"
	"io/ioutil"
)

type token struct {
	codeCount, charCount int
	literal              string
}

type tokeniser struct {
	index int
	bytes []byte
}

func newTokeniser(b []byte) *tokeniser {
	return &tokeniser{-1, b}
}

func (t *tokeniser) next() bool {
	t.index++
	return t.index < len(t.bytes)
}

func main() {
	var bytes, _ = ioutil.ReadFile("inputs/day8.txt")
	var tokens = tokenise(bytes)
	var totalCode, totalChar = calculateTotal(tokens)
	var res = totalCode - totalChar
	fmt.Println("Result: %d", res)
}

func tokenise(bytes []byte) []token {

}

func calculateTotal(tokens []token) (int, int) {

}
