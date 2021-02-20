package main

func subSequeue(a, b string) int {
	cells := CreateMatrix(len(a)+1, len(b)+1)

	for i:=1; i<= len(a); i++ {
		for j:=1; j<= len(b); j++ {
			if a[i-1] == b[j-1] {
				cells[i][j] = cells[i-1][j-1] + 1
			} else {
				cells[i][j] = cells[i-1][j]

				if cells[i][j] < cells[i][j-1] {
					cells[i][j] = cells[i][j-1]
				}
			}
		}
	}
	return cells[len(a)][len(b)]
}
