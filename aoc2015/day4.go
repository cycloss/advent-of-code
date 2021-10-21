package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"strconv"
)

const key = "iwrupvqb"
const mask = 15

func main() {

	var fiveFound = false
	for i := 1; ; i++ {
		var hasher = getNewHasher()
		hasher.Write([]byte(strconv.Itoa(i)))
		var hash = hasher.Sum(nil)
		if checkHashZeroes(hash, 5) && !fiveFound {
			fmt.Printf("First instance of hash with 5 zeros: %d\n", i)
			fiveFound = true
		}
		if checkHashZeroes(hash, 6) {
			fmt.Printf("First instance of hash with 6 zeros: %d\n", i)
			break
		}
	}

}

func getNewHasher() hash.Hash {
	var hash = md5.New()
	hash.Write([]byte(key))
	return hash
}

// byte is 2 hex characters
func checkHashZeroes(bs []byte, zeroCount int) bool {
	var byteIndex = 0
	// check one hex character at a time
	for i := 0; i < zeroCount; i++ {
		var b = bs[byteIndex]
		var modified byte
		if i%2 == 0 {
			// check first half
			// shift right 4
			// if zero then ok
			modified = b >> 4
		} else {
			// check second half and line up a new number
			// mask off first 4 bits
			modified = b & mask
			byteIndex++
		}
		if modified != 0 {
			return false
		}
	}
	return true
}
