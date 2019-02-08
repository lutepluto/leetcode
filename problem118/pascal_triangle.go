package problem118

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		if i > 1 {
			previousRow := triangle[i-1]
			for j := 1; j < i; j++ {
				row[j] = previousRow[j-1] + previousRow[j]
			}
		}
		triangle[i] = row
	}
	return triangle
}
