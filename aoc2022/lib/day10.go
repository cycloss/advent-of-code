package lib

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cycloss/advent-of-code/utils"
)

type instruction struct {
	name   string
	value  int
	cycles int
}

func Day10() {

	lines := utils.GetFileLines("inputs/day10")
	ins := getInstructions(lines)

	// how many cycles have been expended
	curCycle := 0
	x := 1
	totalSigStrength := 0
	for _, in := range ins {
		// fmt.Printf("%+v\n", in)
		for c := in.cycles; c > 0; c-- {

			curCycle++
			if (curCycle-20)%40 == 0 {
				// out of cycles for current block
				fmt.Printf("current cycle: %d, x register: %d\n", curCycle, x)
				totalSigStrength += x * curCycle
			}
		}
		// only after the last cycle has finished should the op be applied
		// as the question is about during cycle
		switch in.name {
		case "addx":
			x += in.value
		default:
		}

	}
	fmt.Printf("total sig strength: %d\n", totalSigStrength)

}

func getInstructions(lines []string) []instruction {
	ins := []instruction{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		in := instruction{split[0], 0, 1}
		if len(split) > 1 {
			in.cycles = 2
			n, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
			in.value = n
		}
		ins = append(ins, in)
	}
	return ins
}
