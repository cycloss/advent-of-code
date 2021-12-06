package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	var file, err = os.Open("day6.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v\n", err)
	}
	defer file.Close()
	solve(file)
}

func solve(file *os.File) {
	var buff = bufio.NewScanner(file)
	var fishes = createStartingFish(buff)
	var part1 = simulate(fishes, 80)
	var part2 = simulate(fishes, 256)
	fmt.Printf("Part 1 Solution: %d\n", part1)
	fmt.Printf("Part 2 Solution: %d\n", part2)

}

func createStartingFish(buff *bufio.Scanner) []int {
	var fishes = make([]int, breedingCycle)

	for buff.Scan() {
		var line = buff.Text()
		var rawNums = strings.Split(line, ",")
		for _, v := range rawNums {
			var num, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal(v)
			}
			fishes[num]++
		}
	}
	return fishes
}

const breedingCycle = 7
const juvenileStage = 2

func simulate(fish []int, days int) *big.Int {
	var adults = convertToBigIntArr(fish)
	var breedingP = 0
	var juveniles = createJuveniles()
	var juvenileP = 0
	for i := 0; i < days; i++ {
		var fishesToBreed = adults[breedingP]
		var newFish = big.NewInt(0)
		// breed new fish
		newFish.Set(fishesToBreed)
		// move oldest juveniles to group that just bred
		var juvenilesToMove = juveniles[juvenileP]
		fishesToBreed.Add(fishesToBreed, juvenilesToMove)
		// set oldest juveniles group to new fish
		juvenilesToMove.Set(newFish)
		breedingP = (breedingP + 1) % breedingCycle
		juvenileP = (juvenileP + 1) % juvenileStage

	}
	var adultTotal = totalFishBig(adults)
	return adultTotal.Add(adultTotal, totalFishBig(juveniles))
}

func convertToBigIntArr(fishes []int) []*big.Int {
	var fishesBig = make([]*big.Int, len(fishes))
	for i, v := range fishes {
		fishesBig[i] = big.NewInt(int64(v))
	}
	return fishesBig
}

func createJuveniles() []*big.Int {
	var juv = make([]int, juvenileStage)
	return convertToBigIntArr(juv)
}

func totalFishBig(fish []*big.Int) *big.Int {
	var total = big.NewInt(0)
	for _, v := range fish {
		total.Add(total, v)
	}
	return total
}
