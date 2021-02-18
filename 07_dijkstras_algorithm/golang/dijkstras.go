package main

import (
	"fmt"
	"math"
)

var graph map[string]map[string]int
var costs map[string]int
var parents map[string]string
var visited []string

func dijkstra(graph map[string]map[string]int, start, end string) (res []string) {
	costs, parents = initData(graph, start)

	// Find the lowest-cost node that you haven't processed yet.
	node := findLowestCostNode(costs)
	// If you've processed all the nodes, this while loop is done.

	for node != "" {
		cost := costs[node]
		neighbors := graph[node]

		for nnode, nbcost := range neighbors {
			newCost := cost + nbcost
			if costs[nnode] > newCost {
				// ... update the cost for this node.
				costs[nnode] = newCost
				// This node becomes the new parent for this neighbor.
				parents[nnode] = node
			}
		}
		// Mark the node as processed.
		visited = append(visited, node)
		// Find the next node to process, and loop.
		node = findLowestCostNode(costs)
	}

	fmt.Println(costs)
	fmt.Println(parents)

	len := len(parents)
	res = make([]string, len+1)
	for len >= 0 {
		res[len] = end
		len--
		if v, ok := parents[end]; ok {
			end = v
		}
	}
	return res
}

func initData(graph map[string]map[string]int, start string) (costs map[string]int, parents map[string]string) {
	//costs & parent
	costs = make(map[string]int)
	parents = make(map[string]string)
	visited = make([]string, 0)

	for node, v := range graph {
		if node == start {
			// costs = v
			for k, val := range v {
				costs[k] = val
				parents[k] = start
			}
		} else {
			if _, ok := costs[node]; !ok {
				costs[node] = math.MaxInt64
			}
			if _, ok := parents[node]; !ok {
				parents[node] = ""
			}
		}
	}
	return
}

func findLowestCostNode(costs map[string]int) string {
	lowestCost := math.MaxInt64
	lowestCostNode := ""

	for node, cost := range costs {
		if cost < lowestCost && !contains(visited, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func contains(set []string, node string) bool {
	for _, item := range set {
		if node == item {
			return true
		}
	}
	return false
}
