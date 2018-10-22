package sorting

func bubbleSort(n []int) []int {
	if len(n) > 1 {
		for i := len(n) - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				if n[j] > n[j+1] {
					n[j], n[j+1] = swap(n[j], n[j+1])
				}
			}
		}
	}
	return n
}
