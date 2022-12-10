package aoc2015

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	location     string
	destinations map[*node]int
}

type solution struct {
	placeMap          map[string]*node
	totalDestinations int
}

func Day9() {
	// generate map with values linked together in a tree
	var placeMap = createPlaceMap()
	// do bfs to find shorts route that visits all
	// create a 'set' for each iteration to find out which have been visited
	var solution = solution{placeMap, len(placeMap)}
	var _, shortestDistance = solution.findShortestRoute(placeMap["AlphaCentauri"], 0, 0)
	fmt.Printf("Shortest distance: %d", shortestDistance)
}

func createPlaceMap() map[string]*node {

	var placeMap = map[string]*node{}
	var file, _ = os.Open("inputs/day9.txt")
	defer file.Close()
	var reader = bufio.NewScanner(file)
	for reader.Scan() {
		var line = reader.Text()
		var tokens = strings.Split(line, " ")
		insertLineIntoMap(placeMap, tokens)
	}

	return placeMap
}

func insertLineIntoMap(placeMap map[string]*node, tokens []string) {
	var from = tokens[0]
	var fromNode = getNode(from, placeMap)
	var distMap = fromNode.destinations

	var to = tokens[2]
	var toNode = getNode(to, placeMap)

	var distTo, _ = strconv.Atoi(tokens[4])

	distMap[toNode] = distTo
}

func getNode(locationName string, placeMap map[string]*node) *node {
	var n = placeMap[locationName]
	if n == nil {
		n = &node{locationName, map[*node]int{}}
		placeMap[locationName] = n
	}
	return n
}

func (s *solution) findShortestRoute(from *node, visitedCount int, totalDist int) (int, int) {
	var currentShortest, visited = (1 << 63) - 1, 0
	for dest, dist := range from.destinations {
		var shortest, shortestCount = s.findShortestRoute(dest, visitedCount+1, totalDist+dist)
		if shortest < currentShortest && visitedCount+1 == s.totalDestinations {
			currentShortest = shortest
			visited = shortestCount
		}
	}

	return currentShortest, visited

}
