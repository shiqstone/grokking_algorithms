package main

import (
	"math"
)

func floyd(graph map[string]map[string]int, start, end string) (res []string) {
	vert, path, n, nmap, smap := initArrData(graph)

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if vert[i][j] > (vert[i][k] + vert[k][j]) {
					vert[i][j] = vert[i][k] + vert[k][j]
					path[i][j] = k
				}
			}
		}
	}
	// fmt.Println(vert, path)
	resij := getPath(path, smap[start], smap[end])
	res = make([]string, 0)
	for _, p := range resij {
		res = append(res, nmap[p])
	}
	return res
}

func initArrData(graph map[string]map[string]int) (map[int]map[int]int, map[int]map[int]int, int, map[int]string, map[string]int) {
	nmap := make(map[int]string)
	smap := make(map[string]int)
	n := 0
	vert := make(map[int]map[int]int)
	path := make(map[int]map[int]int)
	for node, _ := range graph {
		vert[n] = make(map[int]int)
		path[n] = make(map[int]int)
		nmap[n] = node
		smap[node] = n
		n++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = -1
			if i == j {
				vert[j][j] = 0
				continue
			}
			if v, ok := graph[nmap[i]][nmap[j]]; ok {
				vert[i][j] = v
			} else {
				vert[i][j] = int(math.MaxUint32)
			}
		}
	}
	return vert, path, n, nmap, smap
}

func getPath(path map[int]map[int]int, start, end int) []int {
	res := make([]int, 0)
	k := path[start][end]
	if k == -1 {
		res = append(res, start, end)
	} else {
		resik := getPath(path, start, k)
		reskj := getPath(path, k, end)
		res = append(res, resik...)
		res = append(res, reskj[1:]...)
	}
	return res
}
