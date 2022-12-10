package templates

import (
	"fmt"

	"github.com/cycloss/advent-of-code/utils"
)

func Template() {

	lines := utils.GetFileLines("inputs/day10")

	for _, line := range lines {
		fmt.Println(line)
	}

}
