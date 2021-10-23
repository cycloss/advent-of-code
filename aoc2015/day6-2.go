package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type instruction2 func(*int)

func main() {

	var commands = generateCommands()
	var lights = generateGrid(1000, 1000)
	for i, _ := range commands {
		commands[i].execute(lights)

	}
	var total = findBrightness(lights)
	fmt.Printf("%d total brightness after instructions\n", total)
}

func generateGrid(x, y int) [][]int {
	var grid = make([][]int, x)
	for i, _ := range grid {
		grid[i] = make([]int, y)
	}
	return grid
}

func generateCommands() []*command2 {
	var file, _ = os.Open("inputs/day6.txt")
	var buff = bufio.NewScanner(file)
	defer file.Close()
	var commands = []*command2{}
	for buff.Scan() {
		var line = buff.Text()
		var command = commandFromString(line)
		commands = append(commands, command)
	}
	return commands
}

type command2 struct {
	from, to *vector2
	ins      instruction2
}

type vector2 struct {
	x, y int
}

func commandFromString(s string) *command2 {
	var command = &command2{}
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

func getInstruction(instStr string) func(*int) {
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

func (c *command2) execute(lights [][]int) {

	for x := c.from.x; x <= c.to.x; x++ {
		for y := c.from.y; y <= c.to.y; y++ {
			var lightp = &lights[x][y]
			c.ins(lightp)
		}
	}
}

func turnOn(i *int) {
	*i++
}

func turnOff(i *int) {
	if *i > 0 {
		*i--
	}
}

func toggle(i *int) {
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
