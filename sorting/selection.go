package sorting

func selectionSort(n []int) []int {
	if len(n) > 1 {
		for i := 0; i < len(n); i++ {
			smallest := i
			for j := i + 1; j < len(n); j++ {
				if n[smallest] > n[j] {
					smallest = j
				}
			}
			n[i], n[smallest] = swap(n[i], n[smallest])
		}
	}
	return n
}
