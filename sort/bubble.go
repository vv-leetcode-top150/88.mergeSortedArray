package sort

func Bubble(source []int) []int {
	result := source
	var swapped bool
	for {
		swapped = false
		for i := 1; i < len(result); i++ {
			if(result[i] < result[i - 1]){
				swapped = swap(result, i, i - 1)
			}
		}
		if(!swapped){
			break
		}
	}
	return result
}

func swap(source []int, n, m int) bool {
	temp := source[m]
	source[m] = source[n]
	source[n] = temp
	return true
}
