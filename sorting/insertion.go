package sorting

func insertionSort(n []int) []int {
	if len(n) > 1 {
		for i := 1; i < len(n); i++ {
			for j := i; j > 0; j-- {
				if n[j] < n[j-1] {
					n[j], n[j-1] = swap(n[j], n[j-1])
				}
			}
		}
	}
	return n
}
