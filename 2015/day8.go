package aoc2015

import (
	"fmt"
	"io/ioutil"
)

type token struct {
	codeCount, charCount, encodedCount int
}

type tokeniser struct {
	index int
	bytes []byte
}

func Day8() {
	var bytes, _ = ioutil.ReadFile("inputs/day8.txt")
	var tokeniser = newTokeniser(bytes)
	var tokens = tokeniser.tokeniseBytes()
	var codeTotal, charTotal, encodedTotal = calculateTotal(tokens)
	var res = codeTotal - charTotal
	fmt.Printf("Result: %d\n", res)
	// plus 2 for surrounding quotes
	var encodedDiff = encodedTotal - codeTotal
	fmt.Printf("Encoded diff: %d\n", encodedDiff)
}

func newTokeniser(b []byte) *tokeniser {
	return &tokeniser{-1, b}
}

func (t *tokeniser) tokeniseBytes() []*token {
	var tokens = []*token{}
	for t.next() {
		var token = t.getNextToken()
		tokens = append(tokens, token)
	}
	return tokens
}

func (t *tokeniser) next() bool {
	t.index++
	return t.index < len(t.bytes)
}

func (t *tokeniser) current() byte {
	return t.bytes[t.index]
}

const backslash = byte('\\')
const quote = byte('"')
const newline byte = '\n'

func (t *tokeniser) getNextToken() *token {

	var current = t.current()
	switch current {
	case newline:
		return &token{}
	case backslash:
		return t.handleEscaped()
	case quote:
		// add 3 as should be surrounding quote
		return &token{codeCount: 1, encodedCount: 3}
	default:
		return &token{charCount: 1, codeCount: 1, encodedCount: 1}
	}
}

func (t *tokeniser) handleEscaped() *token {
	if !t.next() {
		return &token{charCount: 1, encodedCount: 1, codeCount: 1}
	}
	var current = t.current()
	switch current {
	case backslash:
		// must escape two back slashes
		return &token{codeCount: 2, charCount: 1, encodedCount: 4}
	case quote:
		return &token{codeCount: 2, charCount: 1, encodedCount: 4}
	default:
		// must be hex sequence
		t.index += 2
		return &token{codeCount: 4, charCount: 1, encodedCount: 5}
	}
}

func calculateTotal(tokens []*token) (int, int, int) {
	var codeTotal, charTotal, encodedTotal int
	for _, token := range tokens {
		codeTotal += token.codeCount
		charTotal += token.charCount
		encodedTotal += token.encodedCount
	}

	return codeTotal, charTotal, encodedTotal
}
