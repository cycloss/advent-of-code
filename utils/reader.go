package utils

import (
	"bufio"
	"log"
	"os"
)

func GetFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	var buff = bufio.NewScanner(file)
	defer file.Close()

	lines := []string{}

	for buff.Scan() {
		var line = buff.Bytes()
		lines = append(lines, string(line))
	}
	return lines
}
