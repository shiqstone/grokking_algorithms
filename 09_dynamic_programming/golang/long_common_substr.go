package main

func substr(a, b string) string {
	lcs := 0
	lastSubIndex := 0
	cells := CreateMatrix(len(a)+1, len(b)+1)

	for i:=1; i<= len(a); i++ {
		for j:=1; j<= len(b); j++ {
			if a[i-1] == b[j-1] {
				cells[i][j] = cells[i-1][j-1] + 1

				if cells[i][j] > lcs {
					lcs = cells[i][j]
					lastSubIndex = i
				}
			} else {
				cells[i][j] = 0
			}
		}
	}
	return a[lastSubIndex-lcs: lastSubIndex]
}
