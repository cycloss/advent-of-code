package main

import (
	"fmt"
	"io/ioutil"
)

type Worker struct {
	name       string
	visited    map[Vector2]bool
	currentPos *Vector2
}

func newWorkerWithMap(name string, m map[Vector2]bool) *Worker {
	var worker = &Worker{}
	worker.name = name
	worker.currentPos = &Vector2{0, 0}
	worker.visited = m
	return worker
}

func newWorker(name string) *Worker {
	return newWorkerWithMap(name, map[Vector2]bool{{0, 0}: true})
}

func (w *Worker) move(v *Vector2) {
	w.currentPos.add(v)
	w.visited[*w.currentPos] = true
}

func (w *Worker) getTotal() int {
	return len(w.visited)
}

type Vector2 struct {
	x int
	y int
}

func (v *Vector2) add(v2 *Vector2) {
	v.x += v2.x
	v.y += v2.y
}

func main() {
	var bytes, _ = ioutil.ReadFile("inputs/day3.txt")

	var loneSanta = newWorker("lone santa")

	var sharedMap = map[Vector2]bool{{0, 0}: true}
	var santa = newWorkerWithMap("santa", sharedMap)
	var roboSanta = newWorkerWithMap("robo santa", sharedMap)

	var vectorMap = map[byte]Vector2{byte('<'): {-1, 0}, byte('>'): {1, 0}, byte('^'): {0, 1}, byte('v'): {0, -1}}

	for i, b := range bytes {
		var vector = vectorMap[b]
		loneSanta.move(&vector)
		if i%2 == 0 {
			santa.move(&vector)
		} else {
			roboSanta.move(&vector)
		}
	}
	fmt.Printf("%d houses visited by %s\n", loneSanta.getTotal(), loneSanta.name)
	fmt.Printf("%d houses visited by %s and %s\n", len(sharedMap), santa.name, roboSanta.name)
}
