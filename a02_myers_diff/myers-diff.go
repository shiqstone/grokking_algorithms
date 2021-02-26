package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	charMode := flag.Bool("char", false, "diff in char mode")
	flag.Parse()

	var src []string
	var dst []string

	if *charMode {
		args := flag.Args()
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "usage: myers-diff -char [src_string] [dst_string]")
			os.Exit(1)
		}
		src = strings.Split(args[0], "")
		dst = strings.Split(args[1], "")
	} else {
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: myers-diff [src_file] [dst_file]")
			os.Exit(1)
		}

		var err error
		src, err = getFileLines(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		dst, err = getFileLines(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	generateDiff(src, dst)
}

func getFileLines(p string) ([]string, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type operation uint

const (
	INSERT operation = 1
	DELETE           = 2
	MOVE             = 3
)

var color = map[operation]string{
	INSERT: "\033[32m",
	DELETE: "\033[31m",
	MOVE:   "\033[39m",
}

func generateDiff(src, dst []string) {
	script := shortestEditScript(src, dst)

	srcIndex := 0
	dstIndex := 0

	for _, op := range script {
		switch op {
		case INSERT:
			fmt.Println(color[op] + "+" + dst[dstIndex])
			dstIndex++
		case MOVE:
			fmt.Println(color[op] + " " + src[srcIndex])
			srcIndex++
			dstIndex++
		case DELETE:
			fmt.Println(color[op] + "-" + src[srcIndex])
			srcIndex++
		}
	}
}

func shortestEditScript(src, dst []string) []operation {
	n := len(src)
	m := len(dst)
	max := n + m
	var trace []map[int]int
	var x, y int

loop:
	for d := 0; d <= max; d++ {
		//k represents the value of the current coordinate x-y

		v := make(map[int]int, d+2)
		trace = append(trace, v)

		//Notice the diagonal
		if d == 0 {
			t := 0
			for len(src) > t && len(dst) > t && src[t] == dst[t] {
				t++
			}
			v[0] = t
			if t == len(src) && t == len(dst) {
				break loop
			}
			continue
		}

		lastV := trace[d-1]

		for k := -d; k <= d; k += 2 {
			if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
				//go right
				x = lastV[k+1]
			} else {
				//go down
				x = lastV[k-1] + 1
			}

			y = x - k

			//diagonal
			for x < n && y < m && src[x] == dst[y] {
				x++
				y++
			}

			v[k] = x

			if x == n && y == m {
				break loop
			}

		}
	}

	// for debug
	//printTrace(trace)

	var script []operation
	x = n
	y = m
	var k, prevK, prevX, prevY int
	for d := len(trace) - 1; d > 0; d-- {
		k = x - y
		lastV := trace[d-1]

		if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}

		prevX = lastV[prevK]
		prevY = prevX - prevK

		for x > prevX && y > prevY {
			script = append(script, MOVE)
			x--
			y--
		}

		if x == prevX {
			script = append(script, INSERT)
		} else if y == prevY {
			script = append(script, DELETE)
		}
		x, y = prevX, prevY
	}

	if trace[0][0] != 0 {
		for i := 0; i < trace[0][0]; i++ {
			script = append(script, MOVE)
		}
	}
	return reverse(script)
}

func printTrace(trace []map[int]int) {
	for d := 0; d < len(trace); d++ {
		fmt.Printf("d = %d:\n", d)
		v := trace[d]
		for k := -d; k <= d; k += 2 {
			x := v[k]
			y := x - k
			fmt.Printf("  k = %2d: (%d, %d)\n", k, x, y)
		}
	}
}

func reverse(s []operation) []operation {
	size := len(s)
	result := make([]operation, size)
	for i, v := range s {
		result[size-i-1] = v
	}
	return result
}
