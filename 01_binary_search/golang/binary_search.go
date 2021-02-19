package main

func binSearch(list []int, in int) bool {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == in {
			return true
		}

		if list[mid] < in {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
