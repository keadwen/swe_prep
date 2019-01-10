// Package sorts contains various sorting algorithms.
package sorts

// Insertion implements insertion sort algorithm for int slices.
func Insertion(data []int) []int {
	sorted := make([]int, len(data))
	copy(sorted, data)

	for i := 1; i < len(data); i++ {
		j := i - 1
		item := sorted[i]
		for ; j >= 0 && item < sorted[j]; j -= 1 {
			sorted[j+1] = sorted[j]
		}
		sorted[j+1] = item
	}
	return sorted
}
