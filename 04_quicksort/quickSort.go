package main

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less = []int{}
	var greater = []int{}
	for _, num := range arr[1:] {
		if pivot > num {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	less = quickSort(less)
	greater = quickSort(greater)
	return append(append(less, pivot), greater...)
}
