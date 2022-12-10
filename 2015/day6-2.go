package aoc2015

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type instruction2 func(*int)

func day6Part2() {

	var commands = generateCommands2()
	var lights = generateGrid2(1000, 1000)
	for i := range commands {
		commands[i].execute(lights)

	}
	var total = findBrightness(lights)
	fmt.Printf("%d total brightness after instructions\n", total)
}

func generateGrid2(x, y int) [][]int {
	var grid = make([][]int, x)
	for i := range grid {
		grid[i] = make([]int, y)
	}
	return grid
}

func generateCommands2() []*command2 {
	var file, _ = os.Open("inputs/day6.txt")
	var buff = bufio.NewScanner(file)
	defer file.Close()
	var commands = []*command2{}
	for buff.Scan() {
		var line = buff.Text()
		var command = commandFromString2(line)
		commands = append(commands, command)
	}
	return commands
}

type command2 struct {
	from, to *vector2
	ins      instruction2
}

func commandFromString2(s string) *command2 {
	var command = &command2{}
	var split = strings.Split(s, " ")
	// parse instruction
	command.ins = getInstruction2(split[1])

	// parse coordinates
	var coodStart = getCoordStart2(split)
	var coordIns = split[coodStart:]
	var from, to = fromToFromCoordString2(coordIns)
	command.from = from
	command.to = to
	return command
}

func getInstruction2(instStr string) func(*int) {
	switch instStr {
	case "on":
		return turnOn2
	case "off":
		return turnOff2
	default:
		return toggle2
	}
}

func getCoordStart2(instStr []string) int {
	if instStr[0] == "turn" {
		return 2
	}
	return 1
}

func fromToFromCoordString2(s []string) (*vector2, *vector2) {
	var from = &vector2{}
	var _, err = fmt.Sscanf(s[0], "%d,%d", &from.x, &from.y)
	var to = &vector2{}
	_, err = fmt.Sscanf(s[2], "%d,%d", &to.x, &to.y)
	if err != nil {
		fmt.Println(s)
		fmt.Println(err)
		panic("unable to scan string into vecto2 pair")
	}
	return from, to
}

func (c *command2) execute(lights [][]int) {

	for x := c.from.x; x <= c.to.x; x++ {
		for y := c.from.y; y <= c.to.y; y++ {
			var lightp = &lights[x][y]
			c.ins(lightp)
		}
	}
}

func turnOn2(i *int) {
	*i++
}

func turnOff2(i *int) {
	if *i > 0 {
		*i--
	}
}

func toggle2(i *int) {
	*i += 2
}

func findBrightness(lights [][]int) int {
	var brightness = 0
	for _, r := range lights {
		for _, l := range r {
			brightness += l
		}

	}
	return brightness
}
