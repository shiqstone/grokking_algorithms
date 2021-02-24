package main

func selectionSort(arr []int) []int {
	size := len(arr)

	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	idx := 0
	for i, ele := range arr {
		if ele < smallest {
			smallest = ele
			idx = i
		}
	}
	return idx
}
