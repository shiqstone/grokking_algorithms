package main

import (
	"fmt"
	"math"
	"strconv"
)

func factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n float64 = 100
	res := factorial(n)
	fmt.Println(res)

	str := ""
	for res > 1 {
		mod := int(math.Mod(res, 10))
		res = res / 10
		str = strconv.Itoa(mod) + str
	}

	fmt.Println(str[:100])
}
