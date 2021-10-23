package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type instruction func(*bool)

func main() {

	var commands = generateCommands()
	var lights = generateGrid(1000, 1000)
	for i, _ := range commands {
		commands[i].execute(lights)

	}
	var total = findLitCount(lights)
	fmt.Printf("%d lights lit after instructions\n", total)
}

func generateGrid(x, y int) [][]bool {
	var grid = make([][]bool, x)
	for i, _ := range grid {
		grid[i] = make([]bool, y)
	}
	return grid
}

func generateCommands() []*command {
	var file, _ = os.Open("inputs/day6.txt")
	var buff = bufio.NewScanner(file)
	defer file.Close()
	var commands = []*command{}
	for buff.Scan() {
		var line = buff.Text()
		var command = commandFromString(line)
		commands = append(commands, command)
	}
	return commands
}

type command struct {
	from, to *vector2
	ins      instruction
}

type vector2 struct {
	x, y int
}

func commandFromString(s string) *command {
	var command = &command{}
	var split = strings.Split(s, " ")
	// parse instruction
	command.ins = getInstruction(split[1])

	// parse coordinates
	var coodStart = getCoordStart(split)
	var coordIns = split[coodStart:]
	var from, to = fromToFromCoordString(coordIns)
	command.from = from
	command.to = to
	return command
}

func getInstruction(instStr string) func(*bool) {
	switch instStr {
	case "on":
		return turnOn
	case "off":
		return turnOff
	default:
		return toggle
	}
}

func getCoordStart(instStr []string) int {
	if instStr[0] == "turn" {
		return 2
	}
	return 1
}

func fromToFromCoordString(s []string) (*vector2, *vector2) {
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

func (c *command) execute(lights [][]bool) {

	for x := c.from.x; x <= c.to.x; x++ {
		for y := c.from.y; y <= c.to.y; y++ {
			var lightp = &lights[x][y]
			c.ins(lightp)
		}
	}
}

func (v *vector2) toOneD() int {
	return v.x*1000 + v.y
}

func turnOn(b *bool) {
	*b = true
}

func turnOff(b *bool) {
	*b = false
}

func toggle(b *bool) {
	*b = !*b
}

func findLitCount(lights [][]bool) int {
	var count = 0
	for _, r := range lights {
		for _, l := range r {
			if l {
				count++
			}
		}

	}
	return count
}
