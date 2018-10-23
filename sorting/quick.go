package sorting

func quickSort(n []int) []int {
	if len(n) > 1 {
		quick(n, 0, len(n)-1)
	}
	return n
}

func quick(n []int, low, high int) {
	if low < high {
		p := partition(n, low, high)
		quick(n, low, p-1)
		quick(n, p+1, high)
	}
}

func partition(n []int, low, high int) int {
	pivot := n[low]
	i := low + 1
	j := high
	for {
		for ; n[i] < pivot; i++ {
			if i == high {
				break
			}
		}
		for ; n[j] > pivot; j-- {
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		n[i], n[j] = swap(n[i], n[j])
	}
	n[low], n[j] = swap(n[low], n[j])
	return j
}
