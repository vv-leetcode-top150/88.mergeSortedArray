package sort

func swap(source []int, n, m int) bool {
	temp := source[m]
	source[m] = source[n]
	source[n] = temp
	return true
}