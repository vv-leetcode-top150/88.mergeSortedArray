package sort

func Insertion(source []int) []int {
	result := source

	for j := 1; j < len(result); j++ {
		key := result[j]
		i := j - 1

		for (i >= 0) && (result[i] > key) {
			result[i+1] = result[i]
			i = i - 1
		}
		result[i+1] = key
	}

	return result
}
