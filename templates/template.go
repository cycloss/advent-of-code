package templates

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Template() {

	file, err := os.Open("inputs/test")
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

	for _, line := range lines {
		fmt.Println(line)
	}

}
