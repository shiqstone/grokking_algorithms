package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]int)

	exp := []string{"start", "b", "a", "fin"}

	res := dijkstra(graph, "start", "fin")
	ast := assert.New(t)
	ast.Equal(exp, res)
}
