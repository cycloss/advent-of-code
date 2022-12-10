package test

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	"github.com/cycloss/advent-of-code/aoc2022/lib"
)

func init() {
	// change to previous directory as directory will be the test dir to start with
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cwd)
	err = os.Chdir(path.Dir(cwd))
	if err != nil {
		log.Fatal(err)
	}
}

func TestPlayground(t *testing.T) {
	fmt.Println(3 / 2)
}

func TestDay10(t *testing.T) {
	lib.Day10()
}
